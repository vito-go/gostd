package userinfo

import "local/gostd/internal/data/database"

type Api struct {
	db *database.Db
}

func (u Api) Hello() string {
	return "hello world"

}

func NewApi(db *database.Db) *Api {
	return &Api{db: db}
}

func (u Api) GetUserInfoById(id int64) *UserInfo {
	return &UserInfo{Name: "xiaoming"}
}

type UserInfo struct {
	Name string
}
