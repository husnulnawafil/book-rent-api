package configs

import (
	"os"
	"sync"
)

type AppConfig struct {
	Port  string
	MySql struct {
		Driver   string
		Name     string
		Address  string
		Port     string
		Username string
		Password string
	}
	Redis struct {
		Host     string
		Password string
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	defaultConfig.Port = os.Getenv("APP_PORT")
	defaultConfig.MySql.Driver = os.Getenv("MYSQL_DRIVER")
	defaultConfig.MySql.Name = os.Getenv("MYSQL_NAME")
	defaultConfig.MySql.Address = os.Getenv("MYSQL_ADDRESS")
	defaultConfig.MySql.Port = os.Getenv("MYSQL_PORT")
	defaultConfig.MySql.Username = os.Getenv("MYSQL_USERNAME")
	defaultConfig.MySql.Password = os.Getenv("MYSQL_PASSWORD")
	defaultConfig.Redis.Host = os.Getenv("REDIS_HOST")
	defaultConfig.Redis.Password = os.Getenv("REDIS_PASSWORD")

	return &defaultConfig
}
