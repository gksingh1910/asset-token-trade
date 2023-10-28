package config

import (
	"os"

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

var Cfg Config

func init() {
	godotenv.Load("./.env")
	SrvConfig := HTTPServerConf{
		Host: os.Getenv("SRVHOST"),
		Port: os.Getenv("SRVHOST"),
	}
	RedisConfig := CacheConf{
		Host: os.Getenv("SRVHOST"),
		Port: os.Getenv("SRVHOST"),
	}
	PGConfig := DBConf{
		Host: os.Getenv("SRVHOST"),
		Port: os.Getenv("SRVHOST"),
	}
	Cfg = Config{
		HTTPServerConfig: SrvConfig,
		CacheConfig:      RedisConfig,
		DatabaseConfig:   PGConfig,
	}

}

func get() *Config {
	return &Cfg
}
