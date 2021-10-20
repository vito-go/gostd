package main

import (
	"flag"

	"gitea.com/liushihao/gostd/logic"
	"gitea.com/liushihao/gostd/logic/conf"
)

func main() {
	env := flag.String("env", "app.yaml", "指定配置文件")
	flag.Parse()

	// //////////////////
	app := logic.Init(conf.Env(*env))

	if err := app.Start(); err != nil {
		panic(err)
	}
}
