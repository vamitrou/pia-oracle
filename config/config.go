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
	Listen   string
	Port     int32
	Hostname string
}

type DatabaseConf struct {
	Host        string
	Port        int32
	Username    string
	Password    string
	ServiceName string `toml:"service_name"`
	Schema      string
	SchemeOut   string `"toml:"schema_out"`
	Table       string
	TableOut    string `toml:"table_out"`
	Query       string
	QueryOut    string `toml:"query_out"`
	QueryOutImp string `toml:"query_out_imp"`
}

type RestConf struct {
	PredictionEndpoint string `toml:"prediction_endpoint"`
	AppHeader          string `toml:"application_header"`
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
