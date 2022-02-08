package helloblogdao

import (
	"github.com/vito-go/gostd/conf"
)

type Dao struct {
	cfg *conf.Cfg

	UserInfoDao *userInfoDao

	// 各种表
}

func NewDao(cfg *conf.Cfg, userInfoDao *userInfoDao) *Dao {
	return &Dao{cfg: cfg, UserInfoDao: userInfoDao}
}
