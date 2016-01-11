package config

import (
	"errors"
	"github.com/BurntSushi/toml"
)

var pia *PiaConf = nil

var Path string = ""

type PiaConf struct {
	Local    LocalConf    `toml:"local"`
	Avro     AvroConf     `toml:"avro"`
	Database DatabaseConf `toml:"database"`
	Rest     RestConf     `toml:"http"`
}

type LocalConf struct {
	Listen   string
	Port     int32
	Hostname string
}

type AvroConf struct {
	OuterSchema string `toml:"outer_schema"`
	InnerSchema string `toml:"inner_schema"`
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

func (c *PiaConf) Load(path string) error {
	if _, err := toml.DecodeFile(path, c); err != nil {
		return errors.New("Could not open config file")
	}
	return nil
}

func GetConfig(confpath string) (*PiaConf, error) {
	var err error
	if pia == nil {
		pia = new(PiaConf)
		err = pia.Load(confpath)
	}
	return pia, err
}
