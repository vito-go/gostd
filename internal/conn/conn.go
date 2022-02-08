// Package conn
// 各种客户端实例化链接方法， redis elasticsearch rpc客户端 等等， 具体调用方法参照redisCli的引用之处
package conn

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vito-go/logging/tid"
	"github.com/vito-go/mylog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vito-go/gostd/pkg/rpc"
	"github.com/vito-go/gostd/pkg/rpc/jsonrpc"
	"github.com/vito-go/gostd/pkg/rpc/msgpackrpc"

	"github.com/vito-go/gostd/conf"
)

// RedisBlogCli 博客redis client.
type RedisBlogCli = redis.Client

// UserRpcCli user服务rpcclient
type UserRpcCli = rpc.Client

func NewRedisBlogCli(cfg *conf.Cfg) (*RedisBlogCli, error) {
	return newRedisClient(cfg)
}
func NewUserRpcCli(cfg *conf.Cfg) (*UserRpcCli, error) {
	return newRpcCliRetry(cfg)
}

// newRedisClient generate a Redis client representing a pool of zero or more
// underlying connections. It's safe for concurrent use by multiple goroutines.
func newRedisClient(cfg *conf.Cfg) (*redis.Client, error) {
	redisCli := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     cfg.RedisConf.Addr,
		Username: cfg.RedisConf.UserName,
		Password: cfg.RedisConf.Password,
		DB:       *cfg.RedisConf.DB,
		// 可以在配置中添加更多需要的配置
	})
	if err := redisCli.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return redisCli, nil
}

// newRpcCli generate  RpcClient client It's safe for concurrent use by multiple goroutines.
// 可以采用gob编码 也可以采用json编码（可以支持java写的rpc服务）。纯go项目请采用gob编码
// rpcCli支持异步回调
func newRpcCli(cfg *conf.Cfg) (*rpc.Client, error) {
	codec, err := getRpcCodec(cfg.RpcClient.Codec, cfg.RpcClient.Addr)
	if err != nil {
		return nil, err
	}
	return rpc.NewClientWithCodec(codec), nil
}

func getRpcCodec(codec conf.Codec, addr string) (rpc.ClientCodec, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	switch codec {
	case conf.CodecJSON:
		return jsonrpc.NewClientCodec(conn), nil
	case conf.CodecGob:
		return rpc.NewGobClientCodec(conn), nil
	case conf.CodecMsgPack:
		return msgpackrpc.NewClientCodec(conn), nil
	default:
		return nil, fmt.Errorf("unknown rpc client codec: %s", codec)
	}
}

func newRpcCliRetry(cfg *conf.Cfg) (*rpc.Client, error) {
	rpcCli, err := newRpcCli(cfg)
	if err != nil {
		return nil, err
	}
	go func() {
		ctx := context.WithValue(context.Background(), "tid", tid.Get()) // retry 服务
		for {
			var err error
			var codecCli rpc.ClientCodec
			time.Sleep(time.Second) // 防止死循环
			mylog.Ctx(ctx).Info("waiting to server ping stop")
			err = rpcCli.Call("Ping.Stop", 1, new(int)) // 阻塞等待服务端响应，表示服务端结束的信号
			if err != nil {
				mylog.Ctx(ctx).Error("server ping stop", err.Error())
			} else {
				mylog.Ctx(ctx).Info("server ping stop")
			}
			for {
				codecCli, err = getRpcCodec(cfg.RpcClient.Codec, cfg.RpcClient.Addr)
				if err != nil {
					mylog.Ctx(ctx).Error(err.Error())
					time.Sleep(time.Second)
					continue
				}
				mylog.Ctx(ctx).Info("重新链接", cfg.RpcClient.Addr)

				break
			}
			rpcCli.SetCodec(codecCli)
		}
	}()
	return rpcCli, err
}

// mongoCli 如ugo需要可以配置mongo client。 todo 若需要 待实现
func mongoCli(cfg *conf.Cfg) (*mongo.Client, error) {
	return mongo.NewClient(&options.ClientOptions{
		AppName: nil,
		Auth: &options.Credential{
			AuthSource: "",
			Username:   "",
			Password:   "",
		},
	})

}
