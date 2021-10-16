package main

import (
	"flag"

	"local/gostd/logic"
	"local/gostd/logic/conf"
)

func main() {
	env := flag.String("env", "app.yaml", "指定配置文件")
	flag.Parse()

	// //////////////////
	app := logic.Init(conf.Env(*env))

	err := app.Start()
	if err != nil {
		panic(err)
	}
}
