package database

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 地址
	})
}

func GetRedisClient() *redis.Client {
	return rdb
}
