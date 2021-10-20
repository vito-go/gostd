package database

import "gitea.com/liushihao/gostd/logic/conf"

type DB struct {
	cfg *conf.Cfg
}

func NewDB(cfg *conf.Cfg) *DB {
	return &DB{cfg: cfg}
}
