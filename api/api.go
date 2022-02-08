package api

import (
	"context"

	helloblogdao "github.com/vito-go/gostd/internal/data/dao/helloblog-dao"
)

//go:rpc
// 利用标准库go rpc开启rpc服务 我们在这里定义好服务的接口 比如有以下几种，接下来自动生成客户端服务端代码
// 这个是rpc定义的接口

type User struct {
	getUserInfoByProfile func(ctx context.Context, profile string) (*helloblogdao.UserInfoModel, error)

	serviceName string
}
