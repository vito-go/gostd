package teacherdao

import (
	"fmt"

	"gitea.com/liushihao/gostd/conf"
	"gitea.com/liushihao/gostd/internal/data/dao"
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

type teacherDB dao.Dao

func NewTeacherDB(cfg *conf.Cfg) (*teacherDB, error) {
	db, err := dao.Open(cfg.Database.Student)
	if err != nil {
		return nil, fmt.Errorf("teacher库链接失败！ err:%w", err)
	}
	return (*teacherDB)(db), nil
}
