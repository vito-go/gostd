//go:build !gorpc

package api

import (
	"context"
	"errors"

	helloblogdao "github.com/vito-go/gostd/internal/data/dao/helloblog-dao"
	"github.com/vito-go/gostd/pkg/rpc"
)

type UserCli struct {
	rpcCli      *rpc.Client
	serviceName string
}
type UserServer interface {
	GetUserInfoByProfile(context.Context, string) (*helloblogdao.UserInfoModel, error)
}

func RegisterUserServer(server *rpc.Server, s UserServer) error {
	receiver := &User{
		getUserInfoByProfile: s.GetUserInfoByProfile,

		serviceName: "github.com/vito-go/gostd/api.User",
	}
	return server.RegisterName(receiver.serviceName, receiver)
}
func NewUserCli(rpcCli *rpc.Client) *UserCli {
	return &UserCli{rpcCli: rpcCli, serviceName: "github.com/vito-go/gostd/api.User"}
}
func (h *User) GetUserInfoByProfile(ctx context.Context, in string, resp *helloblogdao.UserInfoModel) error {
	if h.getUserInfoByProfile == nil {
		return errors.New("nil func")
	}
	out, err := h.getUserInfoByProfile(ctx, in)
	if err != nil {
		return err
	}
	*resp = *out
	return nil
}

func (r *UserCli) GetUserInfoByProfile(arg string) (*helloblogdao.UserInfoModel, error) {
	var resp = new(helloblogdao.UserInfoModel)
	err := r.rpcCli.Call(r.serviceName+".GetUserInfoByProfile", arg, resp)
	return resp, err
}
