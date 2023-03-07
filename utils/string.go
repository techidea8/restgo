package utils

import (
	"math/rand"
	"time"
	"unicode"
)

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("0123456789")

func RandSeq(n ...int) string {
	return RandStr(letters, n...)
}

func RandNumber(n ...int) string {
	return RandStr(numbers, n...)
}

func RandStr(chars []rune, n ...int) string {
	num := 12
	lens := len(chars)
	if len(n) > 0 {
		num = n[0]
	}
	b := make([]rune, num)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = chars[r.Intn(lens)]
	}
	return string(b)
}
