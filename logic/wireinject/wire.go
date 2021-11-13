//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/google/wire"

	"github.com/liushihao/gostd/conf"
	"github.com/liushihao/gostd/internal/data/api/student"
	"github.com/liushihao/gostd/internal/data/api/student/class"
	"github.com/liushihao/gostd/internal/data/api/student/grades"
	userinfo "github.com/liushihao/gostd/internal/data/api/student/user-info"
	"github.com/liushihao/gostd/internal/data/api/teacher"
	"github.com/liushihao/gostd/internal/data/api/teacher/info"
	"github.com/liushihao/gostd/internal/data/conn"
	"github.com/liushihao/gostd/internal/data/dao/studentdao"
	"github.com/liushihao/gostd/internal/data/dao/teacherdao"
	"github.com/liushihao/gostd/logic"
	"github.com/liushihao/gostd/logic/httpserver"
	"github.com/liushihao/gostd/logic/myrpc"
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
