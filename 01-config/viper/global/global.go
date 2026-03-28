package global

import (
	"encoding/json"
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/*
全局配置
*/
var (
	GormDB   *gorm.DB
	Rdb      *redis.Client
	Cfg      Config
	VP       *viper.Viper
	Log      *zap.Logger
	Enforcer *casbin.Enforcer
)

/*
全局配置结构体
*/
type Config struct {
	System     System     `mapstructure:"system" json:"system" yaml:"system"`
	MySQL      MySQL      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	JWT        JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap        Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`
	CORS       CORS       `mapstructure:"cors" json:"cors" yaml:"cors"`
	Monitoring Monitoring `mapstructure:"monitoring" json:"monitoring" yaml:"monitoring"`
	GRPC       GRPC       `mapstructure:"grpc" json:"grpc" yaml:"grpc"`
	Casbin     Casbin     `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
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

type MySQL struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`
}

func (e *MySQL) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	PoolSize int    `mapstructure:"pool-size" json:"pool-size" yaml:"pool-size"`
}

func (e *Redis) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`
	ExpiresTime int64  `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"`
	BufferTime  int64  `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}

func (e *JWT) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director      string `mapstructure:"director" json:"director" yaml:"director"`
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"`
}

func (e *Zap) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type CORS struct {
	Mode      string   `mapstructure:"mode" json:"mode" yaml:"mode"`
	Whitelist []string `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`
}

func (e *CORS) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type Monitoring struct {
	Enabled bool          `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Metrics MetricsConfig `mapstructure:"metrics" json:"metrics" yaml:"metrics"`
}

func (e *Monitoring) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type MetricsConfig struct {
	Enabled bool       `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Path    string     `mapstructure:"path" json:"path" yaml:"path"`
	Auth    AuthConfig `mapstructure:"auth" json:"auth" yaml:"auth"`
}

func (e *MetricsConfig) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type AuthConfig struct {
	Enabled  bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

func (e *AuthConfig) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type GRPC struct {
	Enabled bool      `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Port    int       `mapstructure:"port" json:"port" yaml:"port"`
	TLS     TLSConfig `mapstructure:"tls" json:"tls" yaml:"tls"`
}

func (e *GRPC) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type TLSConfig struct {
	Enabled bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Cert    string `mapstructure:"cert" json:"cert" yaml:"cert"`
	Key     string `mapstructure:"key" json:"key" yaml:"key"`
}

func (e *TLSConfig) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"model-path" yaml:"model-path"`
}

func (e *Casbin) String() string {
	dj, _ := json.Marshal(&e)
	return string(dj)
}
