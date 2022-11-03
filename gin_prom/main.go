package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	r := gin.New()

	r.Use(RequestCostHandler(),
		gin.Recovery(),
		Metrics(),
		Logger(),
	)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		c.JSON(http.StatusOK, gin.H{"code": 200})
	})

	r.Run(":9999")
}
