package file

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	userinfo "gitea.com/liushihao/gostd/internal/data/api/student/user-info"
)

type File struct {
	studentAip *student.API
}

func NewFile(s *student.API) *File {
	return &File{studentAip: s}
}
func (f *File) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.studentAip.UserInfoClientAPI.Hello()))
}
func (f *File) UserData(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.Write([]byte(`<h1>error: id参数不正确，必须为一个数字</h1>`))
		return
	}
	var wg sync.WaitGroup
	var userInfo *userinfo.UserInfo
	var grades int64
	wg.Add(1)
	go func() {
		defer wg.Done()
		userInfo = f.studentAip.UserInfoClientAPI.GetUserInfoByID(idInt)
	}()
	wg.Add(1) // 千万不要不写否则无法捕获goroutine的panic
	go func() {
		defer wg.Done()
		grades = f.studentAip.GradesClientAPI.GetTotalGradesByID(idInt)
	}()
	wg.Wait()
	w.Write([]byte(fmt.Sprintln(userInfo, grades)))
}
