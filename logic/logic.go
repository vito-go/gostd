package logic

import (
	"net/http"

	"gitea.com/liushihao/mylog"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/internal/data/api/teacher"
	"gitea.com/liushihao/gostd/logic/api/handler"
	"gitea.com/liushihao/gostd/logic/api/myrpc"
	"gitea.com/liushihao/gostd/logic/conf"
)

type App struct {
	Cfg        *conf.Cfg
	HTTPServer *handler.Server
	RpcServer  *myrpc.Server
	StudentAPI *student.API
	TeacherAPI *teacher.API
}

// NewApp 需要不断的增加参数.
func NewApp(cfg *conf.Cfg, httpHandler *handler.Server, studentApi *student.API, teacherAPI *teacher.API,
	rpcServer *myrpc.Server) *App {
	return &App{Cfg: cfg, StudentAPI: studentApi, HTTPServer: httpHandler, TeacherAPI: teacherAPI, RpcServer: rpcServer}
}

type Apple struct {
}

func (Apple) Add(a int, result *int) (err error) {
	*result += a
	return
}

func (a *App) Start() error {
	// 启动各种服务
	mylog.Info("正在启动rpc服务,RpcAddr: ", a.Cfg.RpcAddr)
	err := a.RpcServer.Start(myrpc.Hello{})
	if err != nil {
		return err
	}
	mylog.Info("正在启动http服务,HttpAddr: ", a.Cfg.HttpAddr)
	err = http.ListenAndServe(a.Cfg.HttpAddr, a.HTTPServer.ServerMux())
	if err != nil {
		return err
	}
	return nil
}
func (a *App) Stop() {

}
