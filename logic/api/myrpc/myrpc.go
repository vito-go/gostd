package myrpc

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"

	"gitea.com/liushihao/mylog"

	"gitea.com/liushihao/gostd/logic/conf"
)

type Server struct {
	cfg    *conf.Cfg
	server *rpc.Server
}

func NewServer(cfg *conf.Cfg) *Server {
	return &Server{
		cfg:    cfg,
		server: rpc.NewServer(),
	}
}

// Start 添加注册方法. 这里需要显示的要求传入注册方法，显式，不是隐式,
// 因为有些接口可能有特殊赋值需求.
func (s *Server) Start(rcvs ...interface{}) error {
	if len(rcvs) == 0 {
		return errors.New("rpc注册方法为空")
	}
	if s.cfg.RpcAddr == "" {
		return errors.New("rpc服务监听地址为空")

	}
	for _, r := range rcvs {
		err := s.server.Register(r)
		if err != nil {
			return fmt.Errorf("%+v rpc方法注册失败！%w", r, err)
		}
	}
	lis, err := net.Listen("tcp", s.cfg.RpcAddr)
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				mylog.Errorf("rpc服务结束！rpc.Serve: accept: %s", err.Error())
				return
			}
			mylog.Infof("rpc.Serve: accept: %s", conn.RemoteAddr())
			go s.server.ServeConn(conn)
		}
	}()
	return nil
}
