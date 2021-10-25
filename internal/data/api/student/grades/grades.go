package grades

import (
	"gitea.com/liushihao/gostd/internal/data/database/studentdb"
)

type Cli struct {
	db *studentdb.Dao
}

func (a *Cli) GetGradesByNameAndId(id int64, name string) {
	panic("implement me")
}

func (a *Cli) GetTotalGradesByID(id int64) int64 {
	return id * id
}

func NewCli(dao *studentdb.Dao) *Cli {
	return &Cli{db: dao}
}
