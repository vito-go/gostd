package grades

import (
	"gitea.com/liushihao/gostd/internal/data/dao/studentdao"
)

type Cli struct {
	db *studentdao.Dao
}

func (a *Cli) GetGradesByNameAndID(id int64, name string) {
	panic("implement me")
}

func (a *Cli) GetTotalGradesByID(id int64) int64 {
	return id * id
}

func NewCli(dao *studentdao.Dao) *Cli {
	return &Cli{db: dao}
}
