package class

import (
	"gitea.com/liushihao/gostd/internal/data/database/studentdb"
)

type Cli struct {
	dao *studentdb.Dao
}

func (A Cli) GetNameByID(id int64) string {
	panic("implement me")
}

func NewCli(db *studentdb.Dao) *Cli {
	return &Cli{dao: db}
}
