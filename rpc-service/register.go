package rpcsrv

import (
	"context"

	"github.com/vito-go/mylog"

	"github.com/vito-go/gostd/api"
	"github.com/vito-go/gostd/rpc-service/gorpc"
)

type Register struct {
	User *gorpc.User
}

type Ping struct {
	stopChan chan struct{}
}

func newPing() *Ping {
	c := make(chan struct{})
	return &Ping{stopChan: c}
}

var PingAll = newPing()

func (p *Ping) Stop(ctx context.Context, arg int, reply *int) error {
	mylog.Ctx(ctx).Info("Ping.Stop waiting --->>>", arg, *reply)
	select {
	case <-ctx.Done():
		mylog.Ctx(ctx).Warn("client close ctx done <<<-----")
	case <-p.stopChan:
		mylog.Ctx(ctx).Warn("gracefully ping stop chan <<<-----")
	}
	return nil
}

// Register 在这里注册rpc方法.
func (s *Server) Register() error {
	// ping服务
	if err := s.server.Register(PingAll); err != nil {
		return err
	}

	// 用户服务
	if err := api.RegisterUserServer(s.server, s.registers.User); err != nil {
		return err
	}
	return nil
}
