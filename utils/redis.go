package utils

import (
	"github.com/go-redis/redis/v9"
	"github.com/husnulnawafil/dot-id-task/configs"
)

func InitRedis(config *configs.AppConfig) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password, // no password set
		DB:       0,                     // use default DB
	})

	return rdb
}
