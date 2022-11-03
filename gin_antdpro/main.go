package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type DataRequest struct {
	Text string `json:"text,omitempty" form:"text"`
}

type DataResponse struct {
	Text string `json:"text,omitempty" form:"text"`
}

func (d *DataRequest) String() string {
	bs, _ := json.Marshal(&d)
	return string(bs)
}

func Failed(c *gin.Context, errString string) {
	c.JSON(http.StatusOK, &Response{
		Success: false,
		Data:    errString,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    data,
	})
}

func GetData(c *gin.Context) {
	var req DataRequest
	err := c.BindQuery(&req)
	if err != nil {
		fmt.Println("Get Data Error:", err)
		Failed(c, err.Error())
	} else {
		fmt.Println("Get Data Request:", req)
		Success(c, &DataResponse{Text: fmt.Sprintf("Get Data Request Is: %s, Reback Is: %s", req.Text, time.Now().String())})
	}
}

func PostData(c *gin.Context) {
	var req DataRequest
	err := c.BindJSON(&req)
	if nil != err {
		fmt.Println("Post Data Error: ", err)
		Failed(c, err.Error())
	} else {
		fmt.Println("Post Data Request: ", req)
		Success(c, &DataResponse{
			Text: fmt.Sprintf("Post Data Request Is: %s, Reback Is: %s", req.Text, time.Now().String()),
		})
	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		Failed(c, "unknown error")
	})

	r.GET("/api/getData", GetData)    // 获取库区列表
	r.POST("/api/postData", PostData) // 获取库区列表

	r.Run(":8080")
}
