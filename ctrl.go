package restgo

import (
	"net/http"
)

type Ctrl struct {
}

func (*Ctrl) RespJson(w http.ResponseWriter, data interface{}, statuscode int) {
	RespJson(w, data, statuscode)
}
//响应成功数据
func (*Ctrl) RespOk(w http.ResponseWriter, data interface{}, datas ...interface{}) {
	RespOk(w, data, datas...)
}

func (*Ctrl) RespOkMap(w http.ResponseWriter, data map[string]interface{}) {
	RespOkMap(w, data)
}

func (*Ctrl) RespList(w http.ResponseWriter, data interface{}, total interface{}) {
	RespList(w, data, total)
}

func (*Ctrl) RespFail(w http.ResponseWriter, msg string, datas ...int) {
	RespFail(w, msg, datas...)
}

func (*Ctrl) Forbidden(w http.ResponseWriter) {
	Forbidden(w)
}
