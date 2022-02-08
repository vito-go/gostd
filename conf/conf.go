package conf

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"

	"github.com/vito-go/mylog"
)

type Env string

// Cfg 配置文件. *Cfg后字段的值可以不用指针了. 可以添加required tag字段，设为false则跳过空值检查.
type Cfg struct {
	HTTPServer httpServer `yaml:"http_server" json:"http_server"`
	RPCServer  rpcServer  `yaml:"rpc_server" json:"rpc_server"`
	RedisConf  RedisConf  `yaml:"redis" json:"redis"`
	RpcClient  rpcClient  `yaml:"rpc_client" json:"rpc_client"`
	Database   database   `yaml:"database" json:"database"`
	LogPath    logPath    `yaml:"log_path" json:"log_path"`
}
type logPath struct {
	Info string `yaml:"info" json:"info"`
	Err  string `yaml:"err" json:"err"`
}
type rpcClient struct {
	Addr  string `yaml:"addr" json:"addr"`
	Codec Codec  `yaml:"codec" json:"codec"`
}
type Codec string

const (
	CodecGob     = "gob"
	CodecJSON    = "json"
	CodecProto   = "proto"
	CodecMsgPack = "msgpack"
)

func NewCfg(env Env) (*Cfg, error) {
	b, err := os.ReadFile(string(env))
	if err != nil {
		return nil, fmt.Errorf("配置文件错误. err: %w", err)
	}
	var cfg Cfg
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("配置文件错误. err: %w", err)
	}
	if err = CheckZeroValue(cfg); err != nil {
		return nil, fmt.Errorf("配置文件错误. env: %s err: %s", env, err.Error())
	}
	return &cfg, nil
}

type rpcServer struct {
	Port  int   `yaml:"port" json:"port"`
	Codec Codec `yaml:"codec" json:"codec"` // 自定义编码方式 支持gob json protobuf 以及其他自定义
}

// httpServer 有关时间的整数设置，均为毫秒.
type httpServer struct {
	Mode        string // 设置gin的mode
	Port        int    `yaml:"port" json:"port"`
	ReadTimeout int    `yaml:"read_timeout" json:"read_timeout" required:"false"`
	// ReadHeaderTimeout is the amount of time allowed to read
	// request headers. The connection's read deadline is reset
	// after reading the headers and the Handler can decide what
	// is considered too slow for the body. If ReadHeaderTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	ReadHeaderTimeout int `yaml:"read_header_timeout" json:"read_header_timeout" required:"false"`
	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	// A zero or negative value means there will be no timeout.
	WriteTimeout int `yaml:"write_timeout" json:"write_timeout" required:"false"`
	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	IdleTimeout int `yaml:"idle_timeout" json:"idle_timeout" required:"false"`
	// MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	// If zero, DefaultMaxHeaderBytes is used.
	MaxHeaderBytes *int `yaml:"max_header_bytes" json:"max_header_bytes" required:"false"`
}

type DBConf struct {
	Dsn        string `yaml:"dsn" json:"dsn"` // todo 日志不输出Dsn
	DriverName string `yaml:"driver_name" json:"driver_name"`
}

// database 代表不同数据库
type database struct {
	HelloBlog DBConf `yaml:"hello_blog" json:"hello_blog"`
	// Teacher DBConf `yaml:"teacher" json:"teacher"`
	// Class   DBConf `yaml:"class" json:"class"`
}

type RedisConf struct {
	Addr     string `yaml:"addr" json:"addr"`
	UserName string `yaml:"user_name" json:"user_name" required:"false"`
	Password string `yaml:"password" json:"password" required:"false"`
	DB       *int   `yaml:"db" json:"db"`
}

// CheckZeroValue 检验配置文件各个字段不能为空. str必须为一个结构体类型.
func CheckZeroValue(str interface{}) error {
	t := reflect.TypeOf(str)
	if t.Kind() != reflect.Struct {
		return errors.New("non-struct type:" + t.String())
	}
	v := reflect.ValueOf(str)
	for k := 0; k < t.NumField(); k++ {
		fieldType := v.Field(k).Kind()
		if fieldType == reflect.Struct {
			if err := CheckZeroValue(v.Field(k).Interface()); err != nil {
				return err
			}
		}
		if v.Field(k).IsZero() {
			required := t.Field(k).Tag.Get("required")
			if required == "false" {
				mylog.Ctx(context.Background()).Warnf("%+v %+v is zero, 但根据校验规则，required为false.跳过检查.", t, t.Field(k).Name)
				continue
			}
			return fmt.Errorf("%+v %+v can not be zero", t, t.Field(k).Name)
		}
	}
	return nil
}
