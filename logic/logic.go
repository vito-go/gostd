package logic

import (
	"gitea.com/liushihao/gostd/logic/api/httpserver"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/internal/data/api/teacher"
	"gitea.com/liushihao/gostd/logic/api/myrpc"
	"gitea.com/liushihao/gostd/logic/conf"
)

type App struct {
	Cfg        *conf.Cfg
	HTTPServer *httpserver.Server
	RpcServer  *myrpc.Server
	StudentAPI *student.API
	TeacherAPI *teacher.API
}

// NewApp 需要不断的增加参数.
func NewApp(cfg *conf.Cfg, httpHandler *httpserver.Server, studentApi *student.API, teacherAPI *teacher.API,
	rpcServer *myrpc.Server) *App {
	return &App{Cfg: cfg, StudentAPI: studentApi, HTTPServer: httpHandler, TeacherAPI: teacherAPI, RpcServer: rpcServer}
}

type Apple struct {
}

func (Apple) Add(a int, result *int) (err error) {
	*result += a
	return
}
