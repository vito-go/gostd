package httpserver

import (
	"github.com/google/wire"

	"github.com/vito-go/gostd/http-service/handler/express"
	"github.com/vito-go/gostd/http-service/handler/user"
)

var ProviderSets = wire.NewSet(
	NewServer,
	express.NewExpress,
	express.NewGetPackage,

	user.NewUser,
	user.NewGetUserInfo,
)
