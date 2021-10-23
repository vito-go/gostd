//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/google/wire"

	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/internal/data/api/student/class"
	"gitea.com/liushihao/gostd/internal/data/api/student/grades"
	userinfo "gitea.com/liushihao/gostd/internal/data/api/student/user-info"
	"gitea.com/liushihao/gostd/internal/data/api/teacher"
	"gitea.com/liushihao/gostd/internal/data/api/teacher/info"
	"gitea.com/liushihao/gostd/internal/data/database"
	"gitea.com/liushihao/gostd/internal/data/database/studentdb"
	"gitea.com/liushihao/gostd/logic"
	"gitea.com/liushihao/gostd/logic/api/handler"
	"gitea.com/liushihao/gostd/logic/api/myrpc"
	"gitea.com/liushihao/gostd/logic/conf"
)

func InitApp(cfg *conf.Cfg) (*logic.App, error) {
	wire.Build(handler.NewServer, student.NewApi,
		logic.NewApp, class.NewTable, userinfo.NewCli,
		studentdb.NewStudentDao, studentdb.NewStudentDB, studentdb.NewUserInfoRepo, studentdb.NewClassRepo,
		database.NewStudentDB, database.NewTeacherDB,
		grades.NewTable, myrpc.NewServer,
		teacher.NewApi, info.NewTable,
	)
	return nil, nil
}
