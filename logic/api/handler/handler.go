package handler

import (
	"net/http"
	"strings"

	"gitea.com/liushihao/mylog"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/internal/data/api/teacher"
	"gitea.com/liushihao/gostd/logic/api/handler/file"
)

type Server struct {
	serverMux  *http.ServeMux
	stu        *student.API
	teacherAPi *teacher.API
}

func (s *Server) ServerMux() *http.ServeMux {
	return s.serverMux
}
func NewServer(stu *student.API, teacherAPi *teacher.API) *Server {
	s := &Server{serverMux: http.NewServeMux(), stu: stu, teacherAPi: teacherAPi}
	s.initHandler()
	return s
}

// initHandler 初始化handler 都卸载这里面。 handler过多的时候可以适当封装一下.
func (s *Server) initHandler() {
	f := file.NewFile(s.stu)
	s.serverMux.HandleFunc("/", preHandle(Index, logging))
	s.serverMux.HandleFunc("/hello", preHandle(f.Hello, logging))
	s.serverMux.HandleFunc("/user", preHandle(f.UserData, logging))
}

// logging 可能不需要 自身有捕获
func logging(r *http.Request) {
	path := r.URL.Path
	if !strings.HasSuffix(path, ".log") && !strings.HasSuffix(path, "favicon.ico") {
		mylog.InfoF("%s %s %s", r.RemoteAddr, path, r.PostForm.Encode())
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
