package dao

import (
	"github.com/go-redis/redis/v8"
	"strconv"
)

var RDB = InitRedis()

func InitRedis() *redis.Client {
	addr := config.Redis.Host + ":" + strconv.Itoa(config.Redis.Port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Redis.Auth,
		DB:       0,
	})
}
