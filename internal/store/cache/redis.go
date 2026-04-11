package cache

import (
	ctx "bluearchive-go/config"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(addr, pw string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     ctx.Load().Cache.RedisAddr,
		Password: ctx.Load().Cache.RedisPassword,
		DB:       ctx.Load().Cache.RedisDB,
	})
}