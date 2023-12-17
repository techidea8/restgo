package core

import (
	"encoding/json"
	"net/http"
	"strings"
)

type RespData struct {
	Code       int         `json:"code"`
	Rows       interface{} `json:"rows,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Msg        string      `json:"msg,omitempty"`
	Total      interface{} `json:"total,omitempty"`
	HttpStatus int         `json:"-"`
}

func NewRespData() *RespData {
	return &RespData{
		HttpStatus: http.StatusOK,
	}
}

// 返回msg
func (r *RespData) Ok(msgs ...string) *RespData {
	if len(msgs) > 0 {
		r.Msg = strings.Join(msgs, ",")
	}
	r.Code = http.StatusOK
	r.HttpStatus = http.StatusOK
	return r
}

// 返回msg
func (r *RespData) Fail(msg string) *RespData {
	r.Msg = msg
	r.Code = http.StatusNotFound
	r.HttpStatus = http.StatusOK
	return r
}

// 返回msg
func (r *RespData) WithData(data interface{}) *RespData {
	r.Data = data
	return r
}

// 返回msg
func (r *RespData) WithRows(rows interface{}) *RespData {
	r.Rows = rows
	return r
}

// 返回msg
func (r *RespData) WithCode(code int) *RespData {
	r.Code = code
	return r
}

// 返回msg
func (r *RespData) WithTotal(total interface{}) *RespData {
	r.Total = total
	return r
}

// 返回msg
func (r *RespData) WithHttpStatus(code int) *RespData {
	r.HttpStatus = code
	return r
}

// 返回msg
func (r *RespData) Json(w http.ResponseWriter) {
	header := w.Header()
	header.Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(r.HttpStatus)
	json.NewEncoder(w).Encode(*r)
}
