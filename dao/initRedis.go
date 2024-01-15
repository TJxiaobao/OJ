package dao

import "github.com/go-redis/redis/v8"

var RDB = InitRedis()

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "这里写自己虚拟机得ip地址、或者是本地得ip地址，视个人情况而定:6379",
		Password: "",
		DB:       0,
	})
}
