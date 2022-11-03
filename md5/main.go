package main

import (
	"crypto/md5"
	"fmt"
)

func Md5Encode(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
