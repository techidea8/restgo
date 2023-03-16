package utils

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Rows  interface{} `json:"rows"`
	Msg   string      `json:"msg"`
	Total int64       `json:"total"`
}

func (r Result) ResponseJson(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(r)
}
func ResultRows(data interface{}, total int64) Result {
	return Result{
		Code:  http.StatusOK,
		Rows:  data,
		Total: total,
		Data:  map[string]interface{}{},
		Msg:   "",
	}
}
func ResultOkMsg(msg string) Result {

	return Result{
		Code:  http.StatusOK,
		Msg:   msg,
		Data:  map[string]interface{}{},
		Rows:  make([]string, 0),
		Total: 0,
	}
}
func ResultOk(data ...interface{}) Result {
	if len(data) == 0 {
		return Result{
			Code:  http.StatusOK,
			Data:  map[string]interface{}{},
			Rows:  make([]string, 0),
			Total: 0,
		}
	} else if len(data) == 1 {
		return Result{
			Code:  http.StatusOK,
			Data:  data[0],
			Rows:  make([]string, 0),
			Total: 0,
		}
	} else {
		return Result{
			Code:  http.StatusOK,
			Data:  data[0],
			Msg:   data[1].(string),
			Rows:  make([]string, 0),
			Total: 0,
		}
	}
}
func ResultError(msg any) Result {
	str := ""
	if str1, ok := msg.(string); ok {
		str = str1
	} else if str1, ok := msg.(error); ok {
		str = str1.Error()
	}
	return Result{
		Code:  http.StatusNotFound,
		Msg:   str,
		Data:  map[string]interface{}{},
		Rows:  make([]string, 0),
		Total: 0,
	}
}
