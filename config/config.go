package config

import (
	"os"

	"github.com/gksingh1910/asset-token-trade/util"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPServerConfig HTTPServerConf
	CacheConfig      CacheConf
	DatabaseConfig   DBConf
}

type HTTPServerConf struct {
	Host string
	Port string
}

type CacheConf struct {
	Host string
	Port string
}

type DBConf struct {
	Host string
	Port string
}

var Cfg *Config

func Init() error {
	err := godotenv.Load("./.env")
	if err != nil {
		return &util.CustomError{
			ErrorCode: 201,
			ErrorMsg:  "Error in Init",
			ErrorR:    err.Error(),
		}
	}
	SrvConfig := HTTPServerConf{
		Host: os.Getenv("SRVHOST"),
		Port: os.Getenv("SRVPORT"),
	}
	RedisConfig := CacheConf{
		Host: os.Getenv("SRVHOST"),
		Port: os.Getenv("SRVHOST"),
	}
	PGConfig := DBConf{
		Host: os.Getenv("SRVHOST"),
		Port: os.Getenv("SRVHOST"),
	}
	Cfg = &Config{
		HTTPServerConfig: SrvConfig,
		CacheConfig:      RedisConfig,
		DatabaseConfig:   PGConfig,
	}
	return nil
}

func Get() *Config {
	return Cfg
}
