package userinfo

import (
	"gitea.com/liushihao/gostd/internal/data/database"
)

type Table struct {
	db *database.StudentDB
}

func NewTable(db *database.StudentDB) *Table {
	return &Table{db: db}
}

func (u Table) Hello() string {
	return "hello world"
}

func (u Table) GetUserInfoByID(id int64) *UserInfo {
	return &UserInfo{Name: "xiaoming"}
}

type UserInfo struct {
	Name string
}
