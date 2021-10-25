package database

import (
	"database/sql"

	_ "github.com/lib/pq"

	"gitea.com/liushihao/gostd/logic/conf"
)

type DB struct {
	pgConf *conf.PgConf // 可以做导出方法
	DB     *sql.DB
}

func Open(pgConf *conf.PgConf) (*DB, error) {
	db, err := sql.Open(pgConf.DriverName, pgConf.Dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DB{
		pgConf: pgConf,
		DB:     db,
	}, nil
}
