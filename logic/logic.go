package logic

import (
	"encoding/json"
	"net/http"

	"gitea.com/liushihao/mylog"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/internal/data/api/student/grades"
	userinfo "gitea.com/liushihao/gostd/internal/data/api/student/user-info"
	"gitea.com/liushihao/gostd/internal/data/database"
	"gitea.com/liushihao/gostd/logic/api/handler"
	"gitea.com/liushihao/gostd/logic/conf"
)

type app struct {
	Cfg         *conf.Cfg
	HTTPHandler *handler.Server
	StudentAPI  *student.API
}

// NewApp 需要不断的增加参数.
func NewApp(cfg *conf.Cfg, studentApi *student.API, httpHandler *handler.Server) *app {
	return &app{Cfg: cfg, StudentAPI: studentApi, HTTPHandler: httpHandler}
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
	db := database.NewDB(cfg)
	gradesAPI := grades.NewApi(db)
	userinfoAPI := userinfo.NewAPI(db)
	studentAPI := student.NewApi(gradesAPI, userinfoAPI)
	httpHandler := handler.NewServer(studentAPI)
	return NewApp(cfg, studentAPI, httpHandler)
}
func (a *app) Start() error {
	// 启动各种欧冠你服务
	return http.ListenAndServe(a.Cfg.HttpAddr, a.HTTPHandler.ServerMux())
}
func (a *app) Stop() {

}
