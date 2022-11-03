package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 使用 MySQL 数据库初始化一个 Xorm 适配器
	a, err := xormadapter.NewAdapter("mysql", "root:ivGP5eI6cmANoMUp@tcp(192.168.104.54:3306)/casbin?charset=utf8", true)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}

	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}

	sub := "alice" // 想要访问资源的用户
	obj := "data1" // 将要被访问的资源
	act := "read"  // 用户对资源实施的操作

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		panic(err)
	}

	if ok == true {
		// 允许 alice 读取 data1
		fmt.Println("允许 alice 读取 data1")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("允许 alice 读取 data1")
	}
}
