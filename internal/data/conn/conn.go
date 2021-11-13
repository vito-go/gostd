// Package conn
// 各种客户端实例化链接方法， redis elasticsearch rpc客户端 等等， 具体调用方法参照redisCli的引用之处
package conn

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"gitea.com/liushihao/gostd/conf"
)

// NewRedisClient generate a Redis client representing a pool of zero or more
// underlying connections. It's safe for concurrent use by multiple goroutines.
func NewRedisClient(cfg *conf.Cfg) (*redis.Client, error) {
	redisCli := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf(":%d", cfg.RedisConf.Port),
		Username: cfg.RedisConf.UserName,
		Password: cfg.RedisConf.Password,
		DB:       cfg.RedisConf.DB,
		// 可以在配置中添加更多需要的配置
	})
	if err := redisCli.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return redisCli, nil
}
