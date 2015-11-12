package config

import (
	"github.com/BurntSushi/toml"
)

var pia *PiaConf = nil

type PiaConf struct {
	Local    LocalConf    `toml:"local"`
	Database DatabaseConf `toml:"database"`
	Rest     RestConf     `toml:"http"`
}

type LocalConf struct {
	Listen string
	Port   int32
}

type DatabaseConf struct {
	Host        string
	Port        int32
	Username    string
	Password    string
	ServiceName string `toml:"service_name"`
	Schema      string
	Table       string
	Query       string
}

type RestConf struct {
	PredictionEndpoint string `toml:"prediction_endpoint"`
}

func (c *PiaConf) Load(path string) {
	if _, err := toml.DecodeFile(path, c); err != nil {
		panic(err)
	}
}

func GetConfig() *PiaConf {
	if pia == nil {
		pia = new(PiaConf)
		pia.Load("conf/pia-oracle.toml")
	}
	return pia
}
