package utils

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

//
func Sha1(in string) string {
	h := sha1.New()
	h.Write([]byte(in))
	l := fmt.Sprintf("%x", h.Sum(nil))
	return l
}

//
func Md5(in string) string {
	h := sha1.New()
	h.Write([]byte(in))
	l := fmt.Sprintf("%x", h.Sum(nil))
	return l
}
func MD5(data string) string {
	return strings.ToUpper(Md5(data))
}
