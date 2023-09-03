package core

import (
	"encoding/json"
	"net/http"
	"strings"
)

// 这里处理
//
//	type IRestCtrl interface {
//		Router(restApp *RestApp)
//	}

type RestCtrl struct {
	Module string // 模块
	Patern string //请求群组
}

func (ctrl RestCtrl) ModuleName() string {
	return ctrl.Module
}

func (ctrl RestCtrl) PaternString() string {
	return ctrl.Patern
}

func (ctrl *RestCtrl) RespList(w http.ResponseWriter, rows interface{}, total interface{}) error{
	header := w.Header()
	header.Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(RespData{
		Code:  http.StatusOK,
		Msg:   "",
		Rows:  rows,
		Total: total,
	})
}
func (ctrl *RestCtrl) RespJson(w http.ResponseWriter, data interface{}, code int, msg string) error {
	header := w.Header()
	header.Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(RespData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func (ctrl *RestCtrl) Resp(w http.ResponseWriter, data RespData) error{
	header := w.Header()
	header.Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(data)
}

func (ctrl *RestCtrl) RespOk(w http.ResponseWriter, data interface{}, msg ...string) error {
	if len(msg) > 0 {
		return ctrl.RespJson(w, data, http.StatusOK, strings.Join(msg, ";"))
	} else {
		return ctrl.RespJson(w, data, http.StatusOK, "")
	}
}

func (ctrl *RestCtrl) RespOkMsg(w http.ResponseWriter, msg string) error{

	return ctrl.RespJson(w, nil, http.StatusOK, msg)

}

func (ctrl *RestCtrl) RespOkMap(w http.ResponseWriter, data map[string]interface{}) error {
	return ctrl.RespOk(w, data, "")
}

func (ctrl *RestCtrl) RespFail(w http.ResponseWriter, msg string) error{
	return ctrl.RespJson(w, nil, http.StatusNotFound, msg)
}
