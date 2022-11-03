package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func uploadOneFile(c *gin.Context) {
	log.Println(c.Request.Header)
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// header调用Filename方法，就可以得到文件名
	filename := fmt.Sprintf("%s/http_upload_and_download/upload/%s", getCurrentWorkPath(), header.Filename)

	// 创建一个文件，文件名为filename，这里的返回值out也是一个File指针
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	// 将file的内容拷贝到out
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusCreated, "upload successful \n")
}

func uploadMultiFile(c *gin.Context) {
	err := c.Request.ParseMultipartForm(4 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件大于4MB"})
		return
	}

	files := c.Request.MultipartForm.File["file"]
	for _, v := range files {
		file, err := v.Open()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件打开失败"})
			return
		}

		b, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}

		filename := fmt.Sprintf("%s/http_upload_and_download/upload/%s", getCurrentWorkPath(), v.Filename)

		f, err := os.Create(filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "上传文件新建失败"})
			return
		}

		_, err = io.Copy(f, bytes.NewReader(b))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "io.Copy失败"})
			return
		}

		f.Close()
		file.Close()
	}
	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}

func downloadFile(c *gin.Context) {
	txt := "hello.xlsx"
	filename := fmt.Sprintf("%s/http_upload_and_download/upload/%s", getCurrentWorkPath(), txt)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", txt))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Writer.Write(b)
}

func downloadFile2(c *gin.Context) {
	resp, err := http.Get("http://localhost:8000/download")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "hello2.xlsx"))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Writer.Write(b)
}

func main() {
	router := gin.Default()

	// 调用POST方法，传入路由参数和路由函数
	router.POST("/upload", uploadOneFile)
	// 上传多少个文件
	router.POST("/upload2", uploadMultiFile)
	// 从本地读取文件，进行下载。
	router.GET("/download", downloadFile)
	// 先从上一个路由读取文件，进行下载。
	router.GET("/download2", downloadFile2)

	// 监听端口8000，注意冒号。
	router.Run(":8000")
}

func getCurrentExecPath() string {
	exePath, err := os.Executable()
	if err != nil {
		return "./"
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

func getCurrentWorkPath() string {
	path, err := os.Getwd()
	if err != nil {
		return "./"
	}
	return path
}
