package studentdb

import (
	"fmt"

	"gitea.com/liushihao/gostd/internal/data/database"
	"gitea.com/liushihao/gostd/logic/conf"
)

type Dao struct {
	cfg          *conf.Cfg
	db           *studentDB
	UserInfoRepo *userInfoRepo
	ClassRepo    *classRepo
	// 各种表
}

func NewDao(cfg *conf.Cfg, db *studentDB, userInfoRepo *userInfoRepo, classRepo *classRepo) *Dao {
	return &Dao{cfg: cfg, db: db, UserInfoRepo: userInfoRepo, ClassRepo: classRepo}
}

type studentDB database.DB

func NewStudentDB(cfg *conf.Cfg) (*studentDB, error) {
	db, err := database.Open(cfg.Database.Student)
	if err != nil {
		return nil, fmt.Errorf("student库链接失败！ err:%w", err)
	}
	return (*studentDB)(db), nil
}
