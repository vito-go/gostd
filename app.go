package gostd

import (
	"github.com/vito-go/gostd/conf"
	"github.com/vito-go/gostd/http-service"
	"github.com/vito-go/gostd/rpc-service"
)

type AppUser struct {
	Cfg       *conf.Cfg
	RPCServer *rpcsrv.Server
}

func NewAppUser(cfg *conf.Cfg, server *rpcsrv.Server) *AppUser {
	return &AppUser{Cfg: cfg, RPCServer: server}
}

type AppBlog struct {
	Cfg *conf.Cfg
	// ScriptJob  *scriptjob.ScriptJob
	HTTPServer *httpserver.Server
}

func NewAppBlog(cfg *conf.Cfg, httpServer *httpserver.Server) *AppBlog {
	return &AppBlog{Cfg: cfg, HTTPServer: httpServer}
}
