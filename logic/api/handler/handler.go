package handler

import (
	"fmt"
	"net/http"

	"local/gostd/internal/data/api/student"
	"local/gostd/logic/api/handler/file"
)

type Server struct {
	s   *http.ServeMux
	stu *student.Api
}

func (s *Server) ServerMux() *http.ServeMux {
	return s.s
}
func NewServer(stu *student.Api) *Server {
	s := &Server{s: http.NewServeMux()}
	s.initHandler(stu)
	return s
}
func (s *Server) initHandler(stu *student.Api) {
	f := file.NewFile(stu)
	s.s.HandleFunc("/", Index)
	s.s.HandleFunc("/hello", f.Hello)
	s.s.HandleFunc("/user", f.UserData)
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
