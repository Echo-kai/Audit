package main

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)


func main() {
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
