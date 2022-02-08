package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/vito-go/mylog"

	"github.com/vito-go/gostd/pkg/database/sql"

	"github.com/vito-go/gostd/conf"
)

type (
	HelloBlogDB = sql.DB
	// BlogDB = sql.DB

)

func NewHelloBlogDB(cfg *conf.Cfg) (*HelloBlogDB, error) {
	return open(cfg.Database.HelloBlog)
}

// func NewBlogDB(cfg *conf.Cfg) (*studentRepo, error) {
// 	return open(cfg.Database.HelloBlog)
// }

func open(dbConf conf.DBConf) (*sql.DB, error) {
	db, err := sql.Open(dbConf.DriverName, dbConf.Dsn)
	if err != nil {
		return nil, err
	}
	const connectTimeOut = time.Second * 3
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeOut)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("dbConf: %+v err: %s", dbConf, err.Error())
	}
	mylog.Ctx(context.TODO()).WithField("dbConf", dbConf).Info("数据库已链接")
	return db, nil
}

func init() {
	sql.Register("postgres", pq.Driver{})
}
