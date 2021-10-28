package conf

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

type Env string

// Cfg 配置文件. *Cfg后字段的值可以不用指针了. 基础类型需要加指针才能取出空判断.
type Cfg struct {
	HttpServer httpServer `yaml:"http_server" json:"http_server"`
	RpcServer  rpcServer  `yaml:"rpc_server" json:"rpc_server"`
	RedisConf  redisConf  `yaml:"redis" json:"redis"`
	Database   database   `yaml:"database" json:"database"`
}

func NewCfg(env Env) (*Cfg, error) {
	b, err := os.ReadFile(string(env))
	if err != nil {
		return nil, err
	}
	var cfg Cfg
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}
	if err = checkZeroValue(cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

type rpcServer struct {
	Port int `yaml:"port" json:"port"`
}
type httpServer struct {
	Port        int  `yaml:"port" json:"port"`
	ReadTimeout *int `yaml:"read_timeout" json:"read_timeout"`
	// ReadHeaderTimeout is the amount of time allowed to read
	// request headers. The connection's read deadline is reset
	// after reading the headers and the Handler can decide what
	// is considered too slow for the body. If ReadHeaderTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	ReadHeaderTimeout *int `yaml:"read_header_timeout" json:"read_header_timeout"`
	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	// A zero or negative value means there will be no timeout.
	WriteTimeout *int `yaml:"write_timeout" json:"write_timeout"`
	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	IdleTimeout *int `yaml:"idle_timeout" json:"idle_timeout"`
	// MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	// If zero, DefaultMaxHeaderBytes is used.
	MaxHeaderBytes *int `yaml:"max_header_bytes" json:"max_header_bytes"`
}

type PgConf struct {
	Dsn        string `yaml:"dsn" json:"-"` // 日志不输出Dsn
	DriverName string `yaml:"driver_name" json:"driver_name"`
}
type database struct {
	Student PgConf `yaml:"student" json:"student"`
	Teacher PgConf `yaml:"teacher" json:"teacher"`
	Class   PgConf `yaml:"class" json:"class"`
}

type redisConf struct {
	Port     int     `yaml:"port" json:"port"`
	UserName *string `yaml:"user_name" json:"user_name"`
	Password *string `yaml:"password" json:"password"`
}

// checkRequired 检验配置文件各个字段不能为空. str必须为一个结构体类型.
func checkZeroValue(str interface{}) error {
	t := reflect.TypeOf(str)
	if t.Kind() != reflect.Struct {
		return errors.New("non-struct type:" + t.String())
	}
	v := reflect.ValueOf(str)
	for k := 0; k < t.NumField(); k++ {
		fieldType := v.Field(k).Kind()
		if fieldType == reflect.Struct {
			if err := checkZeroValue(v.Field(k).Interface()); err != nil {
				return err
			}
		}
		if v.Field(k).IsZero() {
			return fmt.Errorf("配置文件错误. err: %+v %+v can not be zero", t, t.Field(k).Name)
		}
	}
	return nil
}
