package student

import (
	"gitea.com/liushihao/gostd/internal/data/api/student/class"
	"gitea.com/liushihao/gostd/internal/data/api/student/grades"
	userinfo "gitea.com/liushihao/gostd/internal/data/api/student/user-info"
)

type API struct {
	GradesIface   grades.Interface
	UserInfoIface userinfo.Interface
	ClassIface    class.Interface
}

func NewApi(g *grades.Table, u *userinfo.Table, c *class.Table) *API {
	return &API{GradesIface: g, UserInfoIface: u, ClassIface: c}
}
