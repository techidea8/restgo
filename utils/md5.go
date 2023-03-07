package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// md5加密
func Md5(in string) string {
	m := md5.New()
	m.Write([]byte(in))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// md5加密
func MD5(in string) string {
	return strings.ToUpper(Md5(in))
}
