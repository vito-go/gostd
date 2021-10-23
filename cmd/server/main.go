package main

import (
	"encoding/json"
	"flag"
	"os"

	"gitea.com/liushihao/mylog"

	_ "gitea.com/liushihao/daemon"

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
	if err := app.Start(); err != nil {
		panic(err)
	}
}
