package restgo

import (
	"log"
	"net/http"
	"strings"

	"github.com/techidea8/restgo/middleware"
)

type MiddlewareRule struct {
	fun      middleware.Middleware
	excludes []string
}

//跨域开关
var cors bool = false

func EnableCors() {
	cors = true
}
func DisableCors() {
	cors = false
}

type GroupRouter struct {
	Module         string
	MiddlewareRule []MiddlewareRule
}

//哪些需要排除在外
//exclude 那些需要配出在外
func (g *GroupRouter) Using(fun func(w http.ResponseWriter, req *http.Request) (bool, error), excludes ...string) (p *GroupRouter) {
	g.MiddlewareRule = append(g.MiddlewareRule, MiddlewareRule{
		fun:      fun,
		excludes: excludes,
	})
	return g
}

var ModuleMap map[string]*GroupRouter = make(map[string]*GroupRouter, 0)

func Module(module string) (p *GroupRouter) {
	pg, ok := ModuleMap[module]
	if !ok {
		pg = &GroupRouter{
			Module:         module,
			MiddlewareRule: make([]MiddlewareRule, 0),
		}
		ModuleMap[module] = pg
	}

	return pg
}

var rpcModules []interface{} = make([]interface{}, 0)

func RegisterRpcModule(module ...interface{}) {
	rpcModules = append(rpcModules, module)
}

//路由处理函数
func (p *GroupRouter) handlerouter(method, act string, fun func(w http.ResponseWriter, req *http.Request)) (r *GroupRouter) {

	if strings.HasPrefix(act, "/") {
		act = strings.TrimPrefix(act, "/")
	}
	prefix := ""
	if len(p.Module) > 0 {
		prefix = "/" + p.Module
	}

	http.HandleFunc(prefix+"/"+act, func(w http.ResponseWriter, req *http.Request) {
		//如果跨域支持

		if cors {
			headers := w.Header()
			headers.Add("Access-Control-Allow-Origin", "*")  //允许访问所有域
			headers.Add("Access-Control-Allow-Headers", "*") //header的类型
			headers.Add("Access-Control-Allow-Methods", "*")
		}
		if req.Method == "OPTIONS" {
			if cors {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusForbidden)
			}
			return
		}
		for _, v := range p.MiddlewareRule {
			shouldrun := true
			for _, a := range v.excludes {
				if a == act {
					shouldrun = false
					break
				}
			}
			if shouldrun {
				next, err := v.fun(w, req)
				//如果有错误,那么需要处理错误
				if err != nil {
					RespFail(w, err.Error(), http.StatusInternalServerError)
					return
				}
				//如果不往下走,
				if !next {
					RespFail(w, "你没有权限进行该操作", http.StatusForbidden)
					return
				}
			}
		}
		//options方法直接返回啦,跨域
		if method == "*" || method == req.Method {
			fun(w, req)
		} else {
			RespFail(w, "不支持该方法", http.StatusNotFound)
		}
	})
	log.Default().Printf("register  [%s] [%s]", method, prefix+"/"+act)
	return p
}

//路由处理函数
func (p *GroupRouter) Router(act string, fun func(w http.ResponseWriter, req *http.Request)) (r *GroupRouter) {
	return p.handlerouter("*", act, fun)
}

//路由处理函数
func (p *GroupRouter) Post(act string, fun func(w http.ResponseWriter, req *http.Request)) (r *GroupRouter) {
	return p.handlerouter("POST", act, fun)
}

//路由处理函数
func (p *GroupRouter) Get(act string, fun func(w http.ResponseWriter, req *http.Request)) (r *GroupRouter) {
	return p.handlerouter("GET", act, fun)
}

//默认的路由规则
func Router(pantern string, fun func(w http.ResponseWriter, req *http.Request)) (p *GroupRouter) {
	return Module("").Router(pantern, fun)
}
