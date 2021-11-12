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
	"gitea.com/liushihao/gostd/internal/data/conn"
	"gitea.com/liushihao/gostd/internal/data/dao/studentdao"
	"gitea.com/liushihao/gostd/internal/data/dao/teacherdao"
	"gitea.com/liushihao/gostd/logic"
	"gitea.com/liushihao/gostd/logic/conf"
	"gitea.com/liushihao/gostd/logic/httpserver"
	"gitea.com/liushihao/gostd/logic/myrpc"
)

func InitApp(cfg *conf.Cfg) (*logic.App, error) {
	wire.Build(httpserver.NewServer,
		logic.NewApp,
		myrpc.NewServer,
		conn.NewRedisClient,
		teacherProviders,
		studentProviders,
	)
	return nil, nil
}

var studentProviders = wire.NewSet(
	student.NewAPI,                                                                                  // student库API
	studentdao.NewDao, studentdao.NewStudentDB, studentdao.NewUserInfoRepo, studentdao.NewClassRepo, // studentdb层面的表
	grades.NewCli, class.NewCli, userinfo.NewCli, // 对于logic层面的调取data层student数据的client
)

var teacherProviders = wire.NewSet(
	teacher.NewApi,                                                     // teacher 库API
	teacherdao.NewDao, teacherdao.NewTeacherDB, teacherdao.NewInfoRepo, // teacherdao 层面的表
	info.NewCli, // 对于logic层面的调取data层teacher数据的client
)
