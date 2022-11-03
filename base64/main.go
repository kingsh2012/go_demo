package main

import "encoding/base64"

// Base64Encode Base64编码
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode Base64解码
func Base64Decode(s string) string {
	str, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(str)
}
