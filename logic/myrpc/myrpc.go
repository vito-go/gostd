package myrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"sync"

	"gitea.com/liushihao/gostd/logic/conf"
	"gitea.com/liushihao/gostd/logic/mylog"
)

type registers []interface{}

type Server struct {
	cfg    *conf.Cfg
	wg     *sync.WaitGroup
	lis    net.Listener
	server *rpc.Server
	// registers registers // todo 是直接将rpc接口放结构体中还是 放在 registers registers 需要注入。。
}

func NewServer(cfg *conf.Cfg) *Server {
	return &Server{
		cfg:    cfg,
		wg:     new(sync.WaitGroup),
		server: rpc.NewServer(),
		// registers: registers,
	}
}

// Start 添加注册方法. 这里需要显示的要求传入注册方法，显式，不是隐式,
// 因为有些接口可能有特殊赋值需求.
func (s *Server) Start(rcvs ...interface{}) error {
	if len(rcvs) == 0 {
		return errors.New("rpc注册方法为空")
	}
	for _, r := range rcvs {
		err := s.server.Register(r)
		if err != nil {
			return fmt.Errorf("rpc方法注册失败！receiver %+v err: %w", r, err)
		}
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.RPCServer.Port))
	if err != nil {
		return err
	}
	s.lis = lis // todo 是否有点隐式？
	for {
		conn, err := lis.Accept()
		if err != nil {
			mylog.Errorf("rpc服务结束！rpc.Serve: accept: %s", err.Error())
			return err
		}
		mylog.Infof("rpc.Serve: accept: %s", conn.RemoteAddr())
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.server.ServeConn(conn)
		}()
	}
}
func (s *Server) Stop(ctx context.Context) error {
	if s.lis != nil {
		if err := s.lis.Close(); err != nil {
			return err
		}
	}
	c := make(chan struct{})
	go func() {
		s.wg.Wait() // 等待所有的grpc链接处理完毕
		c <- struct{}{}
	}()
	select {
	case <-c:
	case <-ctx.Done():
	}
	return ctx.Err()
}
