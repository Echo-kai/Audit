package main

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"time"
)

var redisClient *redis.Client

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "Echo_kai0214",
		DB:          0,
		ReadTimeout: 1000 * time.Millisecond,
	})
	r := gin.Default()
	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("audio"),
		ginprom.Path("metrics"))
	r.Use(p.Instrument())
	r.POST("/audio/submit", Submit)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
