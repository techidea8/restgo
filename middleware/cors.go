package middleware

import (
	"net/http"
)

// 处理跨域请求,支持options访问
func Cors() Middleware {
	return func(w http.ResponseWriter, req *http.Request) (bool, error) {
		headers := w.Header()
		headers.Add("Access-Control-Allow-Origin", "*") //允许访问所有域
		//w.Header().Add("Access-Control-Allow-Headers", "Authorization,Content-Type,Depth,User-Agent,X-File-Size,X-Requested-With,X-Requested-By,If-Modified-Since,X-File-Name, X-File-Type,Cache-Control,Origin,Content-Type,X-Client,X-Platform,X-Token") //header的类型
		headers.Add("Access-Control-Allow-Headers", "*") //header的类型
		headers.Add("Access-Control-Allow-Methods", "*")
		headers.Add("Access-Control-Allow-Headers", "*") //header的类型
		return true, nil
	}
}
