package student

import (
	"context"

	helloblogdao "github.com/vito-go/gostd/internal/data/dao/helloblog-dao"
)

type UserInfo struct {
	Dao *helloblogdao.Dao
}

func NewUserInfo(dao *helloblogdao.Dao) *UserInfo {
	return &UserInfo{Dao: dao}
}

func (c *UserInfo) GetUserInfoByProfile(ctx context.Context, profile string) (*helloblogdao.UserInfoModel, error) {
	return c.Dao.UserInfoDao.ItemByProfile(ctx, profile)
}
