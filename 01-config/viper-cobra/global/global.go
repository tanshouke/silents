package global

import (
	"encoding/json"
	"github.com/spf13/viper"
)

/*
全局配置
*/
var (
	Cfg Config
	VP  viper.Viper
)

/*
全局配置结构体
*/
type Config struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Test   Test   `json:"test" yaml:"test"`
}

func (e *Config) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type System struct {
	Env    string `mapstructure:"env" json:"env" yaml:"env"`
	Addr   int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
}

func (e *System) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type Test struct {
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

func (e *Test) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}
