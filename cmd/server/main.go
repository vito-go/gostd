package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitea.com/liushihao/mylog"

	_ "gitea.com/liushihao/daemon"

	"gitea.com/liushihao/gostd/logic/api/httpserver"
	"gitea.com/liushihao/gostd/logic/api/myrpc"
	"gitea.com/liushihao/gostd/logic/conf"
	"gitea.com/liushihao/gostd/logic/wireinject"
)

func main() {
	envPath := flag.String("env", "cmd/server/app.yaml", "指定配置文件")
	out := flag.Bool("out", false, "是否为标准输出")
	flag.Parse()
	// //////////////////
	cfg, err := conf.NewCfg(conf.Env(*envPath))
	if err != nil {
		panic(err)
	}
	cfgBytes, err := json.MarshalIndent(&cfg, "", "  ")
	if err != nil {
		panic(err)
	}
	if !*out {
		f, err := os.Create("gostd.log")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		mylog.Init(f)
	}
	mylog.Info("out:", *out)
	mylog.Info("envPath:", *envPath)
	mylog.Info("config:", string(cfgBytes))
	app, err := wireinject.InitApp(cfg)
	if err != nil {
		panic(err)
	}
	chanExit := make(chan struct{}, 1)
	go func() {
		mylog.Infof("正在启动http服务,addr: [:%d]", app.Cfg.HttpServer.Port)
		if err := app.HTTPServer.Start(); err != nil {
			mylog.Errorf("http服务启动失败！%d err: %s", app.Cfg.HttpServer.Port, err.Error())
			chanExit <- struct{}{}
		}
	}()
	go func() {
		mylog.Infof("正在启动rpc服务,addr: [:%d]", app.Cfg.RpcServer.Port)
		rpcRegisters := []interface{}{&myrpc.Hello{}, &myrpc.HelloService{}} //
		if err := app.RpcServer.Start(rpcRegisters...); err != nil {
			mylog.Errorf("rpc服务启动失败！%d err: %s", app.Cfg.RpcServer.Port, err.Error())
			chanExit <- struct{}{}
		}
	}()

	safeExit(chanExit, app.HTTPServer, app.RpcServer)
}

// safeExit 优雅的实现程序退出. 退出当前程序不影响下次程序启动，但是会在设定时间内优先处理完当前未完成的链接.
func safeExit(chanExit chan struct{}, httpSrv *httpserver.Server, rpcSrv *myrpc.Server) {
	c := make(chan os.Signal)
	// If no signals are provided, all incoming signals will be relayed to c.
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT) // 监听键盘终止，以及 kill-15 的信号。注意无法获取kill -9的信号
	select {
	case <-chanExit:
		os.Exit(1)
	case sig := <-c:
		mylog.Warnf("收到进程退出信号: %s", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		gracefulStopChan := make(chan struct{})
		go func() {
			defer func() {
				gracefulStopChan <- struct{}{}
			}()
			if err := httpSrv.Stop(ctx); err != nil {
				mylog.Errorf("httpServer stop failed. err: %s", err)
				return
			}
			mylog.Info("httpServer has gracefully stopped")
		}()
		go func() {
			defer func() {
				gracefulStopChan <- struct{}{}
			}()
			if err := rpcSrv.Stop(ctx); err != nil {
				mylog.Errorf("rpcServer stop failed. err: %s", err)
				return
			}
			mylog.Info("rpcServer has gracefully stopped")
		}()

		select {
		case <-ctx.Done():
			mylog.Warn("gracefulExit  with timeout")
			os.Exit(1)
		case <-gracefulStopChan:
			os.Exit(1)
		}
	}
}
