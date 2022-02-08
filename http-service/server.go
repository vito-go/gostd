package httpserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/vito-go/gostd/conf"
	"github.com/vito-go/gostd/http-service/handler/express"
	"github.com/vito-go/gostd/http-service/handler/user"
)

// Server 启动http服务.
type Server struct {
	cfg    *conf.Cfg
	engine *gin.Engine
	srv    *http.Server

	user    *user.User
	express *express.Express
}

func (s *Server) Cfg() *conf.Cfg {
	return s.cfg
}

func (s *Server) Engine() *gin.Engine {
	return s.engine
}

func NewServer(cfg *conf.Cfg, user *user.User, express *express.Express) *Server {
	gin.SetMode(cfg.HTTPServer.Mode)
	engine := gin.New()
	// engine.Use(gin.Recovery())
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.HTTPServer.Port),
		Handler:      engine,
		ReadTimeout:  time.Millisecond * time.Duration(cfg.HTTPServer.ReadTimeout),
		WriteTimeout: time.Millisecond * time.Duration(cfg.HTTPServer.WriteTimeout),
	}
	s := &Server{cfg: cfg, srv: srv, engine: engine, user: user, express: express}
	return s
}

func (s *Server) Start() error {
	// 所有的路由路径都显式写在这里，如ugo超过64行可以封装函数
	s.routerInit()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.HTTPServer.Port))
	if err != nil {
		return err
	}
	return s.srv.Serve(lis)
}
func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

// func add(conn *websocket.Conn) {
// 	defer conn.Close()
// 	mylog.Warn("ws已经链接", conn.RemoteAddr())
// 	err := websocket.Message.Send(conn, "hello")
// 	if err != nil {
// 		mylog.Error(err)
// 	}
// }

// func ws(lis net.Listener) error {
// 	// websocket实时日志系统
// 	mux := http.NewServeMux()
// 	mux.Handle("/universe/api/v1/im/ws/log", websocket.Handler(add))
// 	srvMux := http.Server{Handler: mux}
// 	mylog.Info("启动websocket服务")
// 	err := srvMux.Serve(lis)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
