package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

func check(e *casbin.Enforcer, sub string, obj string, act string) {
	b, err := e.Enforce(sub, obj, act)
	if err != nil {
		log.Fatalf("error: Enforce: %s", err)
	}
	if b {
		fmt.Printf("sub:%s,obj:%s,act:%s 通过验证\n", sub, obj, act)
	} else {
		fmt.Printf("sub:%s,obj:%s,act:%s 未通过验证\n", sub, obj, act)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./casbin/basic/model.conf", "./casbin/basic/policy.csv")
	if err != nil {
		log.Fatalf("error: NewEnforcer: %s", err)
	}

	check(e, "dajun", "data1", "read")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")
	check(e, "dajun", "data2", "write")
	fmt.Println()
	check(e, "root", "data1", "read")
	check(e, "root", "data2", "read")
	check(e, "root", "data1", "write")
	check(e, "root", "data2", "write")
	check(e, "root", "data3", "write")
	check(e, "root", "data3", "write")
	fmt.Println()
	check(e, "lizi", "data1", "read")
	check(e, "lizi", "data1", "write")
	check(e, "lizi", "data2", "read")
	check(e, "lizi", "data2", "write")
}
