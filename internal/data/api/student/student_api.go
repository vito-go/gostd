package student

import (
	"gitea.com/liushihao/gostd/internal/data/api/student/grades"
	userinfo "gitea.com/liushihao/gostd/internal/data/api/student/user-info"
)

type API struct {
	G grades.Interface
	U userinfo.Interface
}

func NewApi(g *grades.API, u *userinfo.API) *API {
	return &API{G: g, U: u}
}
