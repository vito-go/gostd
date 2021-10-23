package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Env string
type PgConf struct {
	Dsn        string `yaml:"dsn" json:"-"` // 日志不输出Dsn
	DriverName string `yaml:"driver_name"`
}
type database struct {
	Student *PgConf `yaml:"student"`
	Teacher *PgConf `yaml:"teacher"`
	Class   *PgConf `yaml:"class"`
}

type RedisConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
}
type Cfg struct {
	HttpAddr string     `yaml:"http_addr"`
	RpcAddr  string     `yaml:"rpc_addr"`
	Redis    *RedisConf `yaml:"redis"`
	Database database   `yaml:"database"`
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

	return &cfg, nil
}
