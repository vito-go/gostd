package student

import (
	"gitea.com/liushihao/gostd/internal/data/api/student/class"
	"gitea.com/liushihao/gostd/internal/data/api/student/grades"
	userinfo "gitea.com/liushihao/gostd/internal/data/api/student/user-info"
	"gitea.com/liushihao/gostd/internal/data/database"
)

type API struct {
	DB            *database.DB
	GradesIface   grades.Interface
	UserInfoIface userinfo.Interface
	ClassIface    class.Interface
}

func (A *API) DBName() string {
	return "student"
}


func NewApi(g *grades.API, u *userinfo.API, c *class.API) *API {
	return &API{GradesIface: g, UserInfoIface: u, ClassIface: c}
}
