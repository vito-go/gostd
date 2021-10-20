package conf

import (
	"encoding/json"
	"os"

	"gitea.com/liushihao/mylog"
	"gopkg.in/yaml.v3"
)

type Cfg struct {
	HttpAddr string `yaml:"http_addr"`
	Redis    string
	Database database
}

type database struct {
	Postgresql pgsql
}
type Env string
type pgsql struct {
	DriverName string
	Host       string
	Port       int
	UserName   string
	Password   string
}

func (p pgsql) Info() string {
	return "链接信息"
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
	cfgBytes, err := json.MarshalIndent(&cfg, "", "  ")
	if err != nil {
		return nil, err
	}
	mylog.Info("config:", string(cfgBytes))
	return &cfg, nil
}
