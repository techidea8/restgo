package utils

import (
	"net/http"
	"strings"
)

// 从request 中获取头
func GetAuthorizationFromRequest(r *http.Request) string {
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	token = strings.TrimPrefix(token, " ")
	return token
}
