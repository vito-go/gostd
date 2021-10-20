package main

import (
	"flag"

	"gitea.com/liushihao/mylog"

	"gitea.com/liushihao/gostd/logic/conf"
	"gitea.com/liushihao/gostd/logic/wireinject"
)

func main() {
	envPath := flag.String("env", "cmd/server/app.yaml", "指定配置文件")
	flag.Parse()
	mylog.Info("envPath:", *envPath)
	// //////////////////
	app, err := wireinject.InitApp(conf.Env(*envPath))
	if err != nil {
		panic(err)
	}
	if err := app.Start(); err != nil {
		panic(err)
	}
}
