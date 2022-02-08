// go:build wireinject
//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/google/wire"

	"github.com/vito-go/gostd/conf"
	"github.com/vito-go/gostd/http-service"
	"github.com/vito-go/gostd/internal/conn"
	"github.com/vito-go/gostd/internal/data"
	"github.com/vito-go/gostd/rpc-service"
	openblog "openblog"
)

func InitAppBlog(cfg *conf.Cfg) (*openblog.AppBlog, error) {
	wire.Build(
		openblog.NewAppBlog,
		// scriptjob.NewScriptJob,
		httpserver.ProviderSets,
		// wirerepo.Providers,
		// conn.NewRedisBlogCli,
		conn.NewUserRpcCli,
	)
	return nil, nil
}

func InitAppUser(cfg *conf.Cfg) (*openblog.AppUser, error) {
	wire.Build(
		openblog.NewAppUser,
		data.Providers,

		rpcsrv.Providers,
	)
	return nil, nil
}
