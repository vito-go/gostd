package logic

import (
	"encoding/json"
	"net/http"

	"gitea.com/liushihao/mylog"

	"local/gostd/internal/data/api/student"
	"local/gostd/internal/data/api/student/grades"
	userinfo "local/gostd/internal/data/api/student/user-info"
	"local/gostd/internal/data/database"
	"local/gostd/logic/api/handler"
	"local/gostd/logic/conf"
)

type app struct {
	Cfg         *conf.Cfg
	HttpHandler *handler.Server
	StudentApi  *student.Api
}

// NewApp 需要不断的增加参数.
func NewApp(cfg *conf.Cfg, studentApi *student.Api, httpHandler *handler.Server) *app {
	return &app{Cfg: cfg, StudentApi: studentApi, HttpHandler: httpHandler}
}

func Init(env conf.Env) *app {
	cfg, err := conf.GetCfg(env)
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		panic(err)
	}
	mylog.Info(string(b))
	db := database.NewDb(cfg)
	gradesApi := grades.NewApi(db)
	userinfoApi := userinfo.NewApi(db)
	studentApi := student.NewApi(gradesApi, userinfoApi)
	httpHandler := handler.NewServer(studentApi)
	return NewApp(cfg, studentApi, httpHandler)
}
func (a *app) Start() error {
	// 启动各种欧冠你服务
	return http.ListenAndServe(a.Cfg.HttpAddr, a.HttpHandler.ServerMux())

}
func (a *app) Stop() {

}
