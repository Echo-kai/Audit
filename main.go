package main

import (
	"Audit/client"
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)

func main() {
	client.InitRedis()
	client.InitMinIo()
	r := gin.Default()
	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("audit"),
		ginprom.Path("metrics"))
	r.Use(p.Instrument())
	r.POST("/audit/submit/:id", Submit)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
