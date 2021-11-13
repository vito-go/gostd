package logic

import (
	"github.com/liushihao/gostd/conf"
	"github.com/liushihao/gostd/internal/data/api/student"
	"github.com/liushihao/gostd/internal/data/api/teacher"
	"github.com/liushihao/gostd/logic/httpserver"
	"github.com/liushihao/gostd/logic/myrpc"
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
