package userinfo

import (
	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/internal/data/database"
)

type API struct {
	db *database.DB
}

func (u API) Hello() string {
	return "hello world"
}

func NewAPI(api *student.API) *API {
	return &API{db: api.DB}
}

func (u API) GetUserInfoByID(id int64) *UserInfo {
	return &UserInfo{Name: "xiaoming"}
}

type UserInfo struct {
	Name string
}
