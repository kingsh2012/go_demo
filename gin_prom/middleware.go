package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
	"time"
)

func RequestCostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("requestStartTime", time.Now())
		c.Next()
	}
}

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()

		cost := time.Since(c.GetTime("requestStartTime"))
		log.Printf("%s %s %s %d %f", c.ClientIP(), c.Request.Method, c.Request.URL.Path, blw.Status(), cost.Seconds())
		fmt.Println(float64(cost.Milliseconds()))
		HttpHistogram.With(prometheus.Labels{"method": c.Request.Method, "code": strconv.Itoa(blw.Status()), "uri": c.Request.URL.Path}).Observe(float64(cost.Milliseconds()))
	}
}

func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
