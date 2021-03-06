package restgo

import (
	"net/http"

	"encoding/json"

	"fmt"
)

type RespData struct {
	Code  interface{} `json:"code"`
	Rows  interface{} `json:"rows,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Msg   interface{} `json:"msg,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

//返回JSON数据
func RespJson(w http.ResponseWriter, data interface{}, statuscode int) {

	header := w.Header()
	header.Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(statuscode)
	ret, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Write(ret)
}

//w,data.msg,code
func RespOk(w http.ResponseWriter, data interface{}, datas ...interface{}) {
	code := http.StatusOK
	d := RespData{
		Code: code,
		Data: data,
	}
	if len(datas) == 1 {
		d.Msg = datas[0]
	}
	if len(datas) == 2 {
		d.Msg = datas[0]
		d.Code = datas[1]
	}
	RespJson(w, d, http.StatusOK)
}

//w,data.msg,code
func RespOkMap(w http.ResponseWriter, data map[string]interface{}) {
	code := http.StatusOK
	RespJson(w, map[string]interface{}{
		"data": data,
		"code": code,
	}, http.StatusOK)
}

//w,data.msg,code
func Forbidden(w http.ResponseWriter) {
	RespFail(w, "你没有权限进行该操作", http.StatusForbidden)
}

//w,data.msg,code
func RespFailMap(w http.ResponseWriter, data map[string]interface{}) {
	code := http.StatusNotFound
	RespJson(w, map[string]interface{}{
		"data": data,
		"code": code,
	}, http.StatusNotFound)
}
func RespFail(w http.ResponseWriter, msg string, datas ...int) {
	code := http.StatusNotFound
	if len(datas) > 0 {
		code = datas[0]
	}
	RespJson(w, RespData{
		Code: code,
		Msg:  msg,
	}, code)
}

func RespList(w http.ResponseWriter, data interface{}, total interface{}) {
	RespJson(w, RespData{
		Code:  http.StatusOK,
		Rows:  data,
		Total: total,
	}, http.StatusOK)
}
