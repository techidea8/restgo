package utils

import (
	"crypto/sha1"
	"fmt"
)

func Sha1(in string) string {
	h := sha1.New()
	h.Write([]byte(in))
	l := fmt.Sprintf("%x", h.Sum(nil))
	return l
}
