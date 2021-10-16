package database

import "local/gostd/logic/conf"

type Db struct {
	cfg *conf.Cfg
}

func NewDb(cfg *conf.Cfg) *Db {
	return &Db{cfg: cfg}
}
