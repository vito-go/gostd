package rpcsrv

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/vito-go/logging/tid"
	"github.com/vito-go/mylog"

	"github.com/vito-go/gostd/conf"
	"github.com/vito-go/gostd/pkg/rpc"
	"github.com/vito-go/gostd/pkg/rpc/jsonrpc"
	"github.com/vito-go/gostd/pkg/rpc/msgpackrpc"
)

type Server struct {
	cfg       *conf.Cfg
	wg        *sync.WaitGroup
	lis       net.Listener
	server    *rpc.Server
	registers *Register // todo 是直接将rpc接口放结构体中还是 放在 registers registers 需要注入。。
}

func NewServer(cfg *conf.Cfg, registers *Register) *Server {
	return &Server{
		cfg:       cfg,
		wg:        new(sync.WaitGroup),
		server:    rpc.NewServer(),
		registers: registers,
	}
}

func getGobServerCodec(conn net.Conn) rpc.ServerCodec {
	return rpc.NewGobServerCodec(conn)
}

func getJsonServerCodec(conn net.Conn) rpc.ServerCodec {
	return jsonrpc.NewServerCodec(conn)
}

// codecServerCodecMap 支持的各种编码.
var codecServerCodecMap = map[conf.Codec]func(conn net.Conn) rpc.ServerCodec{
	conf.CodecGob: func(conn net.Conn) rpc.ServerCodec {
		return rpc.NewGobServerCodec(conn)
	},
	conf.CodecJSON: func(conn net.Conn) rpc.ServerCodec {
		return jsonrpc.NewServerCodec(conn)
	},
	conf.CodecMsgPack: func(conn net.Conn) rpc.ServerCodec {
		return msgpackrpc.NewServerCodec(conn)
	},
}

// Start 添加注册方法. 这里需要显示的要求传入注册方法，显式，不是隐式,
// 因为有些接口可能有特殊赋值需求.
func (s *Server) Start() error {
	if err := s.Register(); err != nil {
		return err
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.RPCServer.Port))
	if err != nil {
		return err
	}
	s.lis = lis                                                      //
	ctx := context.WithValue(context.Background(), "tid", tid.Get()) // 记录所有的client
	codec := s.cfg.RPCServer.Codec
	getSrvCodecFunc, ok := codecServerCodecMap[codec]
	if !ok {
		return errors.New("unknown rpc server codec. currently supported for json and gob")
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			return err
		}
		mylog.Ctx(ctx).Infof("rpc.Serve accept: %s", conn.RemoteAddr())
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			defer mylog.Ctx(ctx).Infof("client over: %s", conn.RemoteAddr())
			s.server.ServeCodec(getSrvCodecFunc(conn))
		}()
	}
}
func (s *Server) Stop(ctx context.Context) error {
	if s.lis != nil {
		if err := s.lis.Close(); err != nil {
			return err
		}
	}
	close(PingAll.stopChan) // 通知所有的客户端进行关闭
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
