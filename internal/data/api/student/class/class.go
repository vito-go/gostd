package class

import (
	"gitea.com/liushihao/gostd/internal/data/dao/studentdao"
)

type Cli struct {
	dao *studentdao.Dao
}

func NewCli(db *studentdao.Dao) *Cli {
	return &Cli{dao: db}
}

func (A Cli) GetNameByID(id int64) string {
	panic("implement me")
}
