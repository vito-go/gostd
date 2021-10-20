//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/google/wire"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/internal/data/api/student/class"
	"gitea.com/liushihao/gostd/internal/data/api/student/grades"
	userinfo "gitea.com/liushihao/gostd/internal/data/api/student/user-info"
	"gitea.com/liushihao/gostd/internal/data/database"
	"gitea.com/liushihao/gostd/logic"
	"gitea.com/liushihao/gostd/logic/api/handler"
	"gitea.com/liushihao/gostd/logic/conf"
)

func InitApp(env conf.Env) (*logic.App, error) {
	wire.Build(conf.NewCfg, handler.NewServer, student.NewApi, logic.NewApp,
		database.NewDB, grades.NewApi, userinfo.NewAPI,class.NewAPI,)
	return nil, nil
}
