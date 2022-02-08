package data

import (
	"github.com/google/wire"

	"github.com/vito-go/gostd/internal/data/dao"
	helloblogdao "github.com/vito-go/gostd/internal/data/dao/helloblog-dao"

	"github.com/vito-go/gostd/internal/data/repo/student"
)

var Providers = wire.NewSet(
	helloblogdao.NewDao,
	helloblogdao.NewUserInfoDao,

	student.NewClient,
	student.NewUserInfo,

	dao.NewHelloBlogDB,
)
