package gorpc

import (
	"context"

	"github.com/vito-go/mylog"

	helloblogdao "github.com/vito-go/gostd/internal/data/dao/helloblog-dao"
	"github.com/vito-go/gostd/internal/data/repo/student"
)

type User struct {
	studentClient *student.Client
}

func NewUser(studentCli *student.Client) *User {
	return &User{studentClient: studentCli}
}

func (u *User) GetUserInfoByProfile(ctx context.Context, profile string) (*helloblogdao.UserInfoModel, error) {
	result, err := u.studentClient.UserInfoRepo.GetUserInfoByProfile(ctx, profile)
	if err != nil {
		mylog.Ctx(ctx).Error(err.Error())
		return nil, err
	}
	return result, nil
}
