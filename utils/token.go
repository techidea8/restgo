package utils

import (
	"net/http"
	"strings"
)

func GetAuthorizationFromRequest(r *http.Request) string {
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	token = strings.TrimPrefix(token, " ")
	return token
}
