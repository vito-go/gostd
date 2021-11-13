package dao

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/liushihao/gostd/conf"
)

type Dao struct {
	pgConf conf.PgConf // 可以做导出方法
	DB     *sql.DB
}

func Open(pgConf conf.PgConf) (*Dao, error) {
	db, err := sql.Open(pgConf.DriverName, pgConf.Dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Dao{
		pgConf: pgConf,
		DB:     db,
	}, nil
}
