package utils

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func PostFormUrlencode(api string, data map[string]string) (r []byte, err error) {
	//post请求
	tmp := make([]string, 0)
	for k, v := range data {
		tmp = append(tmp, k+"="+url.QueryEscape(v))
	}
	resp, err := http.Post(api, "application/x-www-form-urlencoded", strings.NewReader(strings.Join(tmp, "&")))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
