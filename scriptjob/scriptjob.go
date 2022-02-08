package scriptjob

import (
	"context"
	"errors"

	"github.com/vito-go/mylog"

	helloblogdao "github.com/vito-go/gostd/internal/data/dao/helloblog-dao"
)

type ScriptJob struct {
	dao *helloblogdao.Dao
}

func NewScriptJob(dao *helloblogdao.Dao) *ScriptJob {
	return &ScriptJob{dao: dao}
}

func (s *ScriptJob) Start(jobName string) error {
	mylog.Ctx(context.TODO()).WithField("jobName", jobName).Info("执行脚本任务...")
	switch jobName {
	case "jobName":
		return nil
	default:
		return errors.New("undefined job")
	}
}
