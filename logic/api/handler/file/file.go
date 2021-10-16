package file

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"local/gostd/internal/data/api/student"
	userinfo "local/gostd/internal/data/api/student/user-info"
)

type File struct {
	studentAip *student.Api
}

func NewFile(s *student.Api) *File {
	return &File{studentAip: s}
}
func (f *File) Hello(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(f.studentAip.U.Hello()))
}
func (f *File) UserData(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	var wg sync.WaitGroup
	var userInfo *userinfo.UserInfo
	var grades int64
	wg.Add(1)
	go func() {
		defer wg.Done()
		userInfo = f.studentAip.U.GetUserInfoById(idInt)
	}()
	wg.Add(1) // 千万不要不写否则无法捕获goroutine的panic
	go func() {
		defer wg.Done()
		grades = f.studentAip.G.GetTotalGradesById(idInt)
	}()
	wg.Wait()
	panic("能不能捕获异常")

	w.Write([]byte(fmt.Sprintln(userInfo, grades)))
}
