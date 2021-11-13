package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/liushihao/gostd/logic/httpserver/file"
	"github.com/liushihao/gostd/pkg/mylog"

	"github.com/liushihao/gostd/conf"
	"github.com/liushihao/gostd/internal/data/api/student"
	"github.com/liushihao/gostd/internal/data/api/teacher"
)

type Server struct {
	cfg        *conf.Cfg
	srv        *http.Server
	redisCli   *redis.Client
	serverMux  *http.ServeMux
	stu        *student.API
	teacherAPI *teacher.API
}

func NewServer(cfg *conf.Cfg, stu *student.API, teacherAPI *teacher.API, redisCli *redis.Client) *Server {
	serverMux := http.NewServeMux()
	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.HTTPServer.Port),
		Handler:           serverMux,
		ReadTimeout:       time.Duration(*cfg.HTTPServer.ReadTimeout),
		ReadHeaderTimeout: time.Duration(*cfg.HTTPServer.ReadHeaderTimeout),
		WriteTimeout:      time.Duration(*cfg.HTTPServer.WriteTimeout),
		IdleTimeout:       time.Duration(*cfg.HTTPServer.IdleTimeout),
		MaxHeaderBytes:    *cfg.HTTPServer.MaxHeaderBytes,
	}
	s := &Server{cfg: cfg, srv: &srv, serverMux: serverMux, stu: stu, teacherAPI: teacherAPI, redisCli: redisCli}
	return s
}

func (s *Server) Start() error {
	f := file.NewFile(s.stu, s.redisCli)
	// start启动的时候 显示的指明加载的路由地址及handler函数
	s.serverMux.HandleFunc("/", preHandle(Index, logging))
	s.serverMux.HandleFunc("/hello", preHandle(f.Hello, logging))
	s.serverMux.HandleFunc("/user", preHandle(f.UserData, logging))
	s.serverMux.HandleFunc("/name", preHandle(f.Name, logging))
	return s.srv.ListenAndServe()
}
func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

// logging 可能不需要 自身有捕获.
func logging(r *http.Request) {
	path := r.URL.Path
	if !strings.HasSuffix(path, ".log") && !strings.HasSuffix(path, "favicon.ico") {
		mylog.Infof("%s %s %s", r.RemoteAddr, path, r.PostForm.Encode())
	}
}
func preHandle(handlerFunc http.HandlerFunc, preHandlers ...func(r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer handlerFunc(w, r)
		for _, ph := range preHandlers {
			ph(r)
		}
	}
}
