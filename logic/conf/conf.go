package conf

import (
	"encoding/json"
	"fmt"
	"os"

	"gitea.com/liushihao/mylog"
	"gopkg.in/yaml.v3"
)

type Env string
type PgConf struct {
	DriverName string
	Host       string
	Port       int
	UserName   string
	Password   string
}
type database struct {
	Student *PgConf
	Teacher *PgConf
	Class   *PgConf
}
type Cfg struct {
	HttpAddr string `yaml:"http_addr"`
	Redis    string
	Database database
}

func (p PgConf) Info() string {
	return fmt.Sprintf("host=%s port=%d username=%s password=%s disablesll=true",
		p.Host, p.Port, p.UserName, p.Password)
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
