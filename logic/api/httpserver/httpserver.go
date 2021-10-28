package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"gitea.com/liushihao/mylog"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/internal/data/api/teacher"
	"gitea.com/liushihao/gostd/logic/api/httpserver/file"
	"gitea.com/liushihao/gostd/logic/conf"
)

type Server struct {
	cfg        *conf.Cfg
	srv        *http.Server
	serverMux  *http.ServeMux
	stu        *student.API
	teacherAPi *teacher.API
}

func NewServer(cfg *conf.Cfg, stu *student.API, teacherAPi *teacher.API) *Server {
	serverMux := http.NewServeMux()
	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.HttpServer.Port),
		Handler:           serverMux,
		ReadTimeout:       time.Duration(*cfg.HttpServer.ReadTimeout),
		ReadHeaderTimeout: time.Duration(*cfg.HttpServer.ReadHeaderTimeout),
		WriteTimeout:      time.Duration(*cfg.HttpServer.WriteTimeout),
		IdleTimeout:       time.Duration(*cfg.HttpServer.IdleTimeout),
		MaxHeaderBytes:    *cfg.HttpServer.MaxHeaderBytes,
	}
	s := &Server{srv: &srv, serverMux: serverMux, stu: stu, teacherAPi: teacherAPi}
	return s
}

func (s *Server) Start() error {
	f := file.NewFile(s.stu)
	// 初始化handler 都卸载这里面。 handler过多的时候可以适当封装一下.
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
