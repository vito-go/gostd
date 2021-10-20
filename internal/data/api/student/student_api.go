package student

import (
	"gitea.com/liushihao/gostd/internal/data/api/student/class"
	"gitea.com/liushihao/gostd/internal/data/api/student/grades"
	userinfo "gitea.com/liushihao/gostd/internal/data/api/student/user-info"
)

type API struct {
	GradesClientAPI   grades.Interface
	UserInfoClientAPI userinfo.Interface
	ClassClientAPI    class.Interface
}

func NewApi(g *grades.Table, u *userinfo.Table, c *class.Table) *API {
	return &API{GradesClientAPI: g, UserInfoClientAPI: u, ClassClientAPI: c}
}
