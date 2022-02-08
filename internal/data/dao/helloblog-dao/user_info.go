package helloblogdao

import (
	"context"
	"fmt"

	"github.com/vito-go/gostd/internal/data/dao"
	"github.com/vito-go/gostd/pkg/sqls"
)

const UserInfoTableName = "user_info"

type userInfoDao struct {
	db *dao.HelloBlogDB
}

func (u *userInfoDao) Db() *dao.HelloBlogDB {
	return u.db
}
func (u *userInfoDao) TableName() string {
	return "user_info"
}

func NewUserInfoDao(db *dao.HelloBlogDB) *userInfoDao {
	return &userInfoDao{db: db}
}

// ItemByProfile 根据profile进行查询获取user信息。
func (u *userInfoDao) ItemByProfile(ctx context.Context, profile string) (*UserInfoModel, error) {
	q := fmt.Sprintf("SELECT * FROM %s where profile='%s';", u.TableName(), profile)
	var m UserInfoModel
	err := sqls.QueryRowContext(ctx, u, q, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
