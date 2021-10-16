package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Cfg struct {
	HttpAddr string `yaml:"http_addr"`
	Redis    string
}
type Env string

func GetCfg(env Env) (*Cfg, error) {
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
