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

type StudentDB DB

func NewStudentDB(cfg *conf.Cfg) (*StudentDB, error) {
	db, err := Open(cfg.Database.Student)
	return (*StudentDB)(db), err
}

type TeacherDB DB

func NewTeacherDB(cfg *conf.Cfg) (*TeacherDB, error) {
	db, err := Open(cfg.Database.Teacher)
	return (*TeacherDB)(db), err
}
