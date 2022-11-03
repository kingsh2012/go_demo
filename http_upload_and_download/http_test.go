package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func TestHttpUploadOneFile(t *testing.T) {
	txt := "hello.xlsx"
	// 构建一个buf对象
	buf := new(bytes.Buffer)

	w := multipart.NewWriter(buf)
	// 传去form字段和文件名称
	fw, err := w.CreateFormFile("file", txt)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 打开要传输的文件
	fd, err := os.Open(txt)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 将要传输的文件拷贝到multipart里面
	_, err = io.Copy(fw, fd)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Close()

	// 构造请求
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/upload", buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 设置http请求头，包含fromdata 提交的文件信息，这个头的具体值有w.FormDataContentType给出
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := http.Client{}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	io.Copy(os.Stderr, resp.Body)
}

func TestHttpUploadMultiFile(t *testing.T) {
	// TODO 这个暂时不知道怎么写
}

func TestOpen(t *testing.T) {
	// open的文件都是只读模式
	fd, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	n, err := fd.Write([]byte("okok"))
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
