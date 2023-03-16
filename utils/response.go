package utils

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg,omitempty"`
	Total int64       `json:"total,omitempty"`
}

func (r Result) ResponseJson(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(r)
}
func ResultRows(data interface{}, total int64) Result {
	return Result{
		Code:  http.StatusOK,
		Data:  data,
		Total: total,
	}
}
func ResultOkMsg(msg string) Result {

	return Result{
		Code: http.StatusOK,
		Msg:  msg,
	}
}
func ResultOk(data ...interface{}) Result {
	if len(data) == 0 {
		return Result{
			Code: http.StatusOK,
		}
	} else if len(data) == 1 {
		return Result{
			Code: http.StatusOK,
			Data: data[0],
		}
	} else {
		return Result{
			Code: http.StatusOK,
			Data: data[0],
			Msg:  data[0].(string),
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
		Code: http.StatusNotFound,
		Msg:  str,
	}
}
