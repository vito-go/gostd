package student

import (
	"gitea.com/liushihao/gostd/internal/data/api/student/class"
	"gitea.com/liushihao/gostd/internal/data/api/student/grades"
	userinfo "gitea.com/liushihao/gostd/internal/data/api/student/user-info"
)

type API struct {
	GradesCliAPI   grades.Interface
	UserInfoCliAPI userinfo.Interface
	ClassCliAPI    class.Interface
}

func NewApi(g *grades.Table, u *userinfo.Cli, c *class.Table) *API {
	return &API{GradesCliAPI: g, UserInfoCliAPI: u, ClassCliAPI: c}
}
