package logic

import (
	"net/http"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/logic/api/handler"
	"gitea.com/liushihao/gostd/logic/conf"
)

type App struct {
	Cfg         *conf.Cfg
	HTTPHandler *handler.Server
	StudentAPI  *student.API
}

// NewApp 需要不断的增加参数.
func NewApp(cfg *conf.Cfg, studentApi *student.API, httpHandler *handler.Server) *App {
	return &App{Cfg: cfg, StudentAPI: studentApi, HTTPHandler: httpHandler}
}

func (a *App) Start() error {
	// 启动各种欧冠你服务
	return http.ListenAndServe(a.Cfg.HttpAddr, a.HTTPHandler.ServerMux())
}
func (a *App) Stop() {

}
