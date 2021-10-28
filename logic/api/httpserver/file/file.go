package file

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"gitea.com/liushihao/mylog"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/internal/data/dberr"
	"gitea.com/liushihao/gostd/logic/api/httpserver/resp"
)

type File struct {
	studentAPI *student.API
}

func NewFile(s *student.API) *File {
	return &File{studentAPI: s}
}
func (f *File) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.studentAPI.UserInfoCliAPI.Hello()))
}
func (f *File) UserData(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.Write([]byte(`<h1>error: id参数不正确，必须为一个数字</h1>`))
		return
	}
	var wg sync.WaitGroup
	var userInfo map[string]string
	var grades int64
	wg.Add(1)
	go func() {
		defer wg.Done()
		userInfo, err = f.studentAPI.UserInfoCliAPI.GetUserInfoMapByID(idInt)
	}()
	wg.Add(1) // 千万不要不写否则无法捕获goroutine的panic
	go func() {
		defer wg.Done()
		grades = f.studentAPI.GradesCliAPI.GetTotalGradesByID(idInt)
	}()
	wg.Wait()
	w.Write([]byte(fmt.Sprintln(userInfo, grades)))
}
func (f *File) Name(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		w.Write(resp.DataErr("id参数不能为空"))
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.Write(resp.DataErrF("id参数必须数字: %s", err.Error()))
		return
	}
	s, err := f.studentAPI.UserInfoCliAPI.GetNameById(idInt)
	if err != nil {
		if err == dberr.ErrNotFound {
			mylog.Warnf("名字获取失败 id: %d  error: %s", idInt, err.Error())
		} else {
			mylog.Errorf("名字获取失败 id: %d  error: %s", idInt, err.Error())
		}
		w.Write(resp.DataErr(err.Error()))
		return
	}
	w.Write(resp.DataOK(s))
	return
}
