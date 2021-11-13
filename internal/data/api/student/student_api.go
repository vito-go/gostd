package student

import (
	"github.com/liushihao/gostd/internal/data/api/student/class"
	"github.com/liushihao/gostd/internal/data/api/student/grades"
	userinfo "github.com/liushihao/gostd/internal/data/api/student/user-info"
)

type API struct {
	GradesCliAPI   grades.Interface
	UserInfoCliAPI userinfo.Interface
	ClassCliAPI    class.Interface
}

func NewAPI(g *grades.Cli, u *userinfo.Cli, c *class.Cli) *API {
	return &API{GradesCliAPI: g, UserInfoCliAPI: u, ClassCliAPI: c}
}
