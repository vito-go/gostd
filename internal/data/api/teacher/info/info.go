package info

import (
	"gitea.com/liushihao/gostd/internal/data/database/teacherdb"
)

type Cli struct {
	dao *teacherdb.Dao
}

func (c *Cli) GetInfoByID(id int64) {
	panic("implement me")
}

func NewCli(dao *teacherdb.Dao) *Cli {
	return &Cli{dao: dao}
}
