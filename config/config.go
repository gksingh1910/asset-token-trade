package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConfig  DatabaseConfig
	SrvConfig ServerConfig
}

type DatabaseConfig struct {
	Host string
	Port string
}

type ServerConfig struct {
	Host string
	Port string
}

var Cfg *Config

func Init() {
	godotenv.Load(".env")

	DatabaseCfg := DatabaseConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
	}
	ServerCfg := ServerConfig{
		Host: os.Getenv("SRV_HOST"),
		Port: os.Getenv("SRV_PORT"),
	}

	Cfg = &Config{
		DbConfig:  DatabaseCfg,
		SrvConfig: ServerCfg,
	}

	fmt.Printf("Hemmo.. %v \n", Cfg.SrvConfig.Host)

}

func Get() *Config {
	return Cfg
}

func main() {
	Init()
	fmt.Printf("hello data is %v ", Get().DbConfig.Host)

}
