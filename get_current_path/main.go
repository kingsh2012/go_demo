package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func plan1() {
	fmt.Println("os.Args[0]:", os.Args[0])

	fmt.Println("plan1 ================================ end")
}

func plan2() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("os.Executable:", exePath)
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	fmt.Println("filepath.EvalSymlinks:", res)

	fmt.Println("plan2 ================================ end")
}

func plan3() {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = filepath.Dir(filename)
	}
	fmt.Println("runtime.Caller:", filename)
	println("path.Dir(filename):", abPath)

	fmt.Println("plan3 ================================ end")
}

func plan4() {
	p, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("./")
	}
	fmt.Println("filepath.Abs(filepath.Dir(os.Args[0])):", p)

	fmt.Println("plan3 ================================ end")
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

func work() {
	plan1()
	plan2()
	plan3()
	plan4()
}

func main() {
	//for i := 0; i < 60; i++ {
	//	work()
	//	time.Sleep(time.Second * 1)
	//}
	work()
}

/*
// 这个是直接在这个是将/root/awesomeProject/get_current_path 目录下将./get_current_path软连接到dd上。
os.Args[0]: ./dd
plan1 ================================ end
os.Executable: /root/awesomeProject/get_current_path/get_current_path
filepath.EvalSymlinks: /root/awesomeProject/get_current_path
plan2 ================================ end
runtime.Caller: /root/awesomeProject/get_current_path/main.go
path.Dir(filename): /root/awesomeProject/get_current_path
plan3 ================================ end
filepath.Abs(filepath.Dir(os.Args[0])): /root/awesomeProject/get_current_path
plan4 ================================ end

// 这个是直接在这个是将/root/awesomeProject/get_current_path/ 目录下执行的./get_current_path
os.Args[0]: ./get_current_path
plan1 ================================ end
os.Executable: /root/awesomeProject/get_current_path/get_current_path
filepath.EvalSymlinks: /root/awesomeProject/get_current_path
plan2 ================================ end
runtime.Caller: /root/awesomeProject/get_current_path/main.go
path.Dir(filename): /root/awesomeProject/get_current_path
plan3 ================================ end
filepath.Abs(filepath.Dir(os.Args[0])): /root/awesomeProject/get_current_path
plan4 ================================ end

// 这个是将/root/awesomeProject/get_current_path/get_current_path软连接到/opt/下
// 可以发现plan1返回的就是./dd，锁定的是当前软连接的位置
// plan4使用的是os.Args所以定位也是/opt/目录下
os.Args[0]: ./dd
plan1 ================================ end
os.Executable: /root/awesomeProject/get_current_path/get_current_path
filepath.EvalSymlinks: /root/awesomeProject/get_current_path
plan2 ================================ end
runtime.Caller: /root/awesomeProject/get_current_path/main.go
path.Dir(filename): /root/awesomeProject/get_current_path
plan3 ================================ end
filepath.Abs(filepath.Dir(os.Args[0])): /opt
plan4 ================================ end

最终结果是使用plan2或者plan3
*/
