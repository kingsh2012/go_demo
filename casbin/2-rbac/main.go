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
	e, err := casbin.NewEnforcer("./casbin/2-rbac/model.conf", "./casbin/2-rbac/policy.csv")
	if err != nil {
		log.Fatalf("error: NewEnforcer: %s", err)
	}

	check(e, "dajun", "data", "read")
	check(e, "dajun", "data", "write")
	fmt.Println()
	check(e, "lizi", "data", "read")
	check(e, "lizi", "data", "write")
}
