package info

import (
	"github.com/liushihao/gostd/internal/data/dao/teacherdao"
)

type Cli struct {
	dao *teacherdao.Dao
}

func (c *Cli) GetInfoByID(id int64) {
	panic("implement me")
}

func NewCli(dao *teacherdao.Dao) *Cli {
	return &Cli{dao: dao}
}
