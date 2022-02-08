package main

import (
	"context"
	"encoding/json"
	"flag"
	"io"
	// _ "net/http/pprof" //性能、goroutine监控服务
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/vito-go/logging/tid"

	_ "github.com/vito-go/daemon"
	"github.com/vito-go/mylog"

	"github.com/vito-go/gostd/conf"
	"github.com/vito-go/gostd/http-service"
	"github.com/vito-go/gostd/wireinject"
)

// 更改配置一定要自测保证配置校验通过，无论线上还是测试！！！ 可以本地进行运行 添加 check_cfg参数，提交代码前校验.
func main() {
	// 在main函数显示指出设定命令行参数
	ctx := context.WithValue(context.Background(), "tid", tid.Get())
	envPath := flag.String("env", "cmd/openblog/app-dev.yaml", "specify the configuration")
	out := flag.Bool("out", true, "only print in os.StdOut, usually for the local running")
	checkCfg := flag.Bool("check_cfg", false, "check the config. exit with code 0 if ok")
	// sjob := flag.String("sjob", "", "start script job and exit")
	flag.Parse()
	// //////////////////
	cfg, err := conf.NewCfg(conf.Env(*envPath))
	if err != nil {
		// 提交代码前需使用make acp 经过配置文件检查
		panic(err)
	}
	cfgB, _ := json.MarshalIndent(cfg, "", "  ")
	if !*out {
		fInfo, err := os.OpenFile(cfg.LogPath.Info, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		fErr, err := os.OpenFile(cfg.LogPath.Err, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		mylog.Init(fInfo, io.MultiWriter(fInfo, fErr), io.MultiWriter(fInfo, fErr), "tid")
		cfgB, _ = json.Marshal(cfg)
	}
	if *checkCfg {
		if err = conf.CheckZeroValue(*cfg); err != nil {
			mylog.Ctx(ctx).Errorf("envPath: %s: Failed！配置文件校验失败！err: %s", *envPath, err)
			os.Exit(1)
		}
		mylog.Ctx(ctx).Infof("envPath: %s: Successfully！配置文件校验通过！", *envPath)
		os.Exit(0)
	} else {
		mylog.Init(os.Stdout, os.Stdout, os.Stdout, "tid")
	}

	app, err := wireinject.InitAppBlog(cfg)
	if err != nil {
		panic(err)
	}
	// if *sjob != "" {
	// 	mylog.Ctx(ctx).Info("正在启动script job服务")
	// 	if err := app.ScriptJob.Start(*sjob); err != nil {
	// 		mylog.Ctx(ctx).Errorf("script job服务启动失败！err: %s", err.Error())
	// 		os.Exit(1)
	// 	}
	// 	os.Exit(0)
	// }
	mylog.Ctx(ctx).Infof("%s: 开源博客系统", filepath.Base(os.Args[0]))
	mylog.Ctx(ctx).Info("envPath:", *envPath)
	mylog.Ctx(ctx).WithField("cfg", cfgB).Info()
	mylog.Ctx(ctx).Info("out:", *out)
	mylog.Ctx(ctx).Info("gin mod:", cfg.HTTPServer.Mode)
	mylog.Ctx(ctx).Info("pid:", os.Getpid())
	chanExit := make(chan struct{}, 1)

	// go func() {
	// 	mylog.Ctx(ctx).Info("正在启动pprof性能监控服务,addr: [:6060]")
	// 	if err := http.ListenAndServe(":6060", nil); err != nil {
	// 		mylog.Ctx(ctx).Errorf("pprof服务启动失败！[:6060] err: %s", err)
	// 		chanExit <- struct{}{}
	// 	}
	// }()

	go func() {
		mylog.Ctx(ctx).Infof("正在启动http服务,addr: [:%d]", app.Cfg.HTTPServer.Port)
		if err := app.HTTPServer.Start(); err != nil {
			mylog.Ctx(ctx).Errorf("http服务启动失败！%d err: %s", app.Cfg.HTTPServer.Port, err.Error())
			chanExit <- struct{}{}
		}
	}()

	// go func() {
	// 	log.Infof("正在启动rpc服务,addr: [:%d]", a.Cfg.RpcServer.Port)
	// 	if err := a.GrpcServer.Start(); err != nil {
	// 		log.Errorf("rpc服务启动失败！%d err: %s", a.Cfg.RpcServer.Port, err)
	// 		chanExit <- struct{}{}
	// 	}
	safeExit(ctx, chanExit, app.HTTPServer)
}

// safeExit 优雅的实现程序退出. 退出当前程序不影响下次程序启动，但是会在设定时间内优先处理完当前未完成的链接.
func safeExit(ctx context.Context, chanExit chan struct{}, httpSrv *httpserver.Server) {
	c := make(chan os.Signal)
	// If no signals are provided, all incoming signals will be relayed to c.
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT) // 监听键盘终止，以及 kill-15 的信号。注意无法获取kill -9的信号
	select {
	case <-chanExit:
		os.Exit(1)
	case sig := <-c:
		mylog.Ctx(ctx).Warnf("收到进程退出信号: %s", sig.String())
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		gracefulStopChan := make(chan struct{})
		go func() {
			if err := httpSrv.Stop(ctx); err != nil {
				mylog.Ctx(ctx).Errorf("httpServer stop failed. err: %s", err)
			}
			gracefulStopChan <- struct{}{}
		}()
		select {
		case <-ctx.Done():
			mylog.Ctx(ctx).Warn("gracefulExit with timeout")
			os.Exit(1)
		case <-gracefulStopChan:
			os.Exit(1)
		}
	}
}
