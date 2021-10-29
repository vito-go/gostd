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
	"gitea.com/liushihao/gostd/internal/data/database/studentdb"
	"gitea.com/liushihao/gostd/internal/data/database/teacherdb"
	"gitea.com/liushihao/gostd/logic"
	"gitea.com/liushihao/gostd/logic/api/httpserver"
	"gitea.com/liushihao/gostd/logic/api/myrpc"
	"gitea.com/liushihao/gostd/logic/conf"
)

func InitApp(cfg *conf.Cfg) (*logic.App, error) {
	wire.Build(httpserver.NewServer,
		logic.NewApp,
		myrpc.NewServer,
		teacherProviders,
		studentProviders,
	)
	return nil, nil
}

var teacherProviders = wire.NewSet(
	teacherdb.NewDao, teacherdb.NewTeacherDB, teacherdb.NewInfoRepo,
	teacher.NewApi, info.NewCli,
)
var studentProviders = wire.NewSet(studentdb.NewDao, studentdb.NewStudentDB, studentdb.NewUserInfoRepo,
	studentdb.NewClassRepo,
	grades.NewCli, class.NewCli, userinfo.NewCli, student.NewAPI,
)
