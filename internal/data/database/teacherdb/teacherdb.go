package teacherdb

import (
	"gitea.com/liushihao/gostd/internal/data/database"
	"gitea.com/liushihao/gostd/logic/conf"
)

type Dao struct {
	cfg      *conf.Cfg
	db       *teacherDB
	InfoRepo *infoRepo
	// 各种表
}

func NewDao(cfg *conf.Cfg, db *teacherDB, infoRepo *infoRepo) *Dao {
	return &Dao{cfg: cfg, db: db, InfoRepo: infoRepo}
}

type teacherDB database.DB

func NewTeacherDB(cfg *conf.Cfg) (*teacherDB, error) {
	db, err := database.Open(cfg.Database.Student)
	return (*teacherDB)(db), err
}
