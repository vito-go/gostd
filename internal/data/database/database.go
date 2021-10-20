package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"gitea.com/liushihao/gostd/logic/conf"
)

type DB struct {
	cfg    *conf.Cfg
	dbName string
	db     *sql.DB
}
type DBName string

type DBNameIface interface {
	DBName() string
}

func NewDB(cfg *conf.Cfg, d DBNameIface) (*DB, error) {
	db, err := sql.Open("postgresql", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgresql.Host, cfg.Postgresql.Port, cfg.Postgresql.UserName, cfg.Postgresql.Password, d.DBName()))

	if err != nil {
		return nil, err
	}
	return &DB{
		cfg:    cfg,
		dbName: "",
		db:     db,
	}, nil
}
