package utils

import (

	"math/rand"
	"regexp"
	"strconv"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func SubStr(str string,num int)string{
	nameRune := []rune(str)
	if(num>len(nameRune)){
		num = len(nameRune)
	}
	return string(nameRune[:num])
}


func ParseStr(str string) float32 {
	digitsRegexp := regexp.MustCompile(`(\d+)\D+(\d+)`)
	ret := digitsRegexp.FindStringSubmatch(str)
	if len(ret)<2{
		return 0
	}else{
		a,e:=strconv.ParseFloat(ret[1],10)
		if(e!=nil){
			return 0.0
		}else{
			return float32(a)
		}
	}

}
