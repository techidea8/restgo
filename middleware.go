package restgo

import "net/http"

//定义个中间件
type Middleware func(w http.ResponseWriter, req *http.Request) (bool, error)
