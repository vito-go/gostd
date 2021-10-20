package handler

import (
	"fmt"
	"net/http"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/logic/api/handler/file"
)

type Server struct {
	serverMux *http.ServeMux
	stu       *student.API
}

func (s *Server) ServerMux() *http.ServeMux {
	return s.serverMux
}
func NewServer(stu *student.API) *Server {
	s := &Server{serverMux: http.NewServeMux(), stu: stu}
	s.initHandler()
	return s
}
func (s *Server) initHandler() {
	f := file.NewFile(s.stu)
	s.serverMux.HandleFunc("/", Index)
	s.serverMux.HandleFunc("/hello", f.Hello)
	s.serverMux.HandleFunc("/user", f.UserData)
}

// SafeHandle 可能不需要 自身有捕获
func SafeHandle(f func(writer http.ResponseWriter, request *http.Request)) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err)
			}
		}()
		f(writer, request)
	}
}
