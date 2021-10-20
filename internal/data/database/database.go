package database

import (
	"database/sql"

	_ "github.com/lib/pq"

	"gitea.com/liushihao/gostd/logic/conf"
)

type DB struct {
	pgConf *conf.PgConf
	db     *sql.DB
}

func open(pgConf *conf.PgConf) (*DB, error) {
	db, err := sql.Open(pgConf.DriverName, pgConf.Info())

	if err != nil {
		return nil, err
	}
	return &DB{
		pgConf: pgConf,
		db:     db,
	}, nil
}

type StudentDB DB

func NewStudentDB(cfg *conf.Cfg) (*StudentDB, error) {
	db, err := open(cfg.Database.Student)
	return (*StudentDB)(db), err
}

type TeacherDB DB

func NewTeacherDB(cfg *conf.Cfg) (*TeacherDB, error) {
	db, err := open(cfg.Database.Student)
	return (*TeacherDB)(db), err
}
