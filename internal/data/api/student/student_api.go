package student

import (
	"local/gostd/internal/data/api/student/grades"
	userinfo "local/gostd/internal/data/api/student/user-info"
)

type Api struct {
	G grades.Interface
	U userinfo.Interface
}

func NewApi(g *grades.Api, u *userinfo.Api) *Api {
	return &Api{G: g, U: u}
}
