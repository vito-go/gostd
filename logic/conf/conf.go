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
}
type Env string

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
