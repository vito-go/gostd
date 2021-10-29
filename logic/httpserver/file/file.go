package file

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"gitea.com/liushihao/gostd/logic/httpserver/resp"
	"gitea.com/liushihao/gostd/logic/mylog"

	"github.com/go-redis/redis/v8"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/internal/data/dberr"
)

type File struct {
	studentAPI *student.API
	redisCli   *redis.Client
}

func NewFile(s *student.API, redisCli *redis.Client) *File {
	return &File{studentAPI: s, redisCli: redisCli}
}
func (f *File) Hello(w http.ResponseWriter, r *http.Request) {
	s := f.studentAPI.UserInfoCliAPI.Hello()
	_, _ = w.Write([]byte(fmt.Sprintf("<h1>%s</h1>", s)))
}
func (f *File) UserData(w http.ResponseWriter, r *http.Request) {
	// 调用redis
	fake := f.redisCli.Get(context.Background(), "fake")
	_ = fake
	id := r.FormValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		_, _ = w.Write([]byte(`<h1>error: id参数不正确，必须为一个数字</h1>`))
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
	_, _ = w.Write([]byte(fmt.Sprintln(userInfo, grades)))
}
func (f *File) Name(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		_, _ = w.Write(resp.DataErr("id参数不能为空"))
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		_, _ = w.Write(resp.DataErrF("id参数必须数字: %s", err.Error()))
		return
	}
	s, err := f.studentAPI.UserInfoCliAPI.GetNameByID(idInt)
	if err != nil {
		if err == dberr.ErrNotFound {
			mylog.Warnf("名字获取失败 id: %d  error: %s", idInt, err.Error())
		} else {
			mylog.Errorf("名字获取失败 id: %d  error: %s", idInt, err.Error())
		}
		_, _ = w.Write(resp.DataErr(err.Error()))
		return
	}
	_, _ = w.Write(resp.DataOK(s))
}
