package restgo

import (

	"github.com/techidea8/restgo/middleware"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type MiddlewareRule struct {
	fun      middleware.Middleware
	excludes []string
}
//跨域开关
var cors bool = false
func EnableCors(){
	cors = true
}
var debug bool = false
func Debug(flag bool){
	debug = flag
}


func DisableCors(){
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

		if(cors){
			headers := w.Header()
			headers.Add("Access-Control-Allow-Origin", "*") //允许访问所有域
			headers.Add("Access-Control-Allow-Headers", "*") //header的类型
			headers.Add("Access-Control-Allow-Methods", "*")
		}
		if req.Method == "OPTIONS" {
			if(cors){
				w.WriteHeader(http.StatusOK)
			}else{
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

//文件服务器支持
//Aliase("/assets/","./resource/assets/")
//外部请求  http://[ip]/assets/img/test.png,系统将
func Aliase(patern string,dirpath string) {
	http.Handle(patern, http.StripPrefix(patern, http.FileServer(http.Dir(dirpath))))
}

//文件服务器支持
//Root("/assets/","./resource/assets/")
//外部请求  http://[ip]/assets/img/test.png,系统将
func Root(patern string,dirpath string) {
	http.Handle(patern, http.StripPrefix("", http.FileServer(http.Dir(dirpath))))
}

//自定义函数支持
var restFuncMap template.FuncMap = make(template.FuncMap)
type Fun  func(arg ...interface{})interface{}
func RegisterFuncMap(fname string,fun Fun) {
	restFuncMap[fname] = fun
}
var globalTemplete *template.Template
//全局函数
var globalData map[string]interface{} = make(map[string]interface{})
func AddDataToTemplete(key string,data interface{}){
	globalData[key] = data
}

//模板文件批量注册
//registertempletes("/","view/**/*.html","index/index.html")\n
// @title    模板解析函数\n
// @description   批量解析模板文件并自动映射成相应url\n
// @auth      winlion             时间（2019/6/18   10:57 ）\n
// @param     prefix        string         "页面访问前奏,类似spring boot中的context"\n
// @param     patern        string         "模板文件路径格式,如view/**/*.html"\n
// @param     indextplname        string         "首页模板文件路径"\n
// @return    无\n
func RegisterTempletes(prefix string,patern string,indextplname string) {

		_globalTemplete, err := template.ParseGlob(patern)
		if err!=nil{
			//fmt.Println(err.Error())
			log.Default().Printf("register tpl  [%s] error:%s",patern, err.Error())
			return
		}
		globalTemplete= _globalTemplete
		for _,v := range globalTemplete.Templates(){

			tplname := v.Name()
				//http.Handle()
				if debug{
					log.Default().Printf("map TPL [%s]=>%s\n",tplname,prefix+tplname)
				}
                contextpath := prefix+tplname
                if tplname==indextplname{
                	contextpath = prefix
				}
				http.HandleFunc(contextpath, func(w http.ResponseWriter, req *http.Request) {
					//如果再次
					if debug{
						globalTemplete, _ = template.ParseGlob(patern)
					}
					log.Default().Printf("prefix+v.Name()=>%s",prefix+v.Name())
					globalTemplete.Funcs(restFuncMap).ExecuteTemplate(w,tplname,globalData)
				})
			}

}