package client

import (
	"github.com/go-redis/redis"
	"time"
)

var RedisClient *redis.Client

func InitRedis(){
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "Echo_kai0214",
		DB:          0,
		ReadTimeout: 1000 * time.Millisecond,
	})
}
