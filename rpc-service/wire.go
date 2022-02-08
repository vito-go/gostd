package rpcsrv

import (
	"github.com/google/wire"

	"github.com/vito-go/gostd/rpc-service/gorpc"
)

var Providers = wire.NewSet(
	NewServer,
	wire.Struct(new(Register), "*"),
	gorpc.NewUser,
)
