package utils

import (
	"fmt"
	"io/ioutil"
	"net"

	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	log "github.com/techidea8/restgo/pkg/log"
)

// HTTPGet get 请求
func HttpGet(url string, headers ...map[string]string) ([]byte, error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	for _, headermap := range headers {
		for k, v := range headermap {
			reqest.Header.Add(k, v)
		}
	}

	resp, err := client.Do(reqest)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

// PostJSON post json 数据请求
func PostJSON(uri string, structObj interface{}) ([]byte, error) {

	jsonData, err := json.Marshal(structObj)
	if err != nil {
		return nil, err
	}

	jsonData = bytes.Replace(jsonData, []byte("\\u003c"), []byte("<"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u003e"), []byte(">"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u0026"), []byte("&"), -1)
	return PostJsonBytes(uri, jsonData)
}

// PostJSON post json 数据请求
func PostJsonString(uri string, jsonData string) ([]byte, error) {
	return PostJsonBytes(uri, []byte(jsonData))
}

// 这是配置json
func PostJsonWithHeader(url string, structObj interface{}, headers map[string]string) (r []byte, err error) {

	bytes, _ := json.Marshal(structObj)

	req, _ := http.NewRequest("POST", url, strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	fmt.Printf("%s=>%s", url, string(bytes))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func PostJsonBytes(uri string, jsonData []byte) ([]byte, error) {

	jsonData = bytes.Replace(jsonData, []byte("\\u003c"), []byte("<"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u003e"), []byte(">"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u0026"), []byte("&"), -1)

	body := bytes.NewBuffer(jsonData)
	log.Debugf("%s=>%s", uri, jsonData)
	response, err := http.Post(uri, "application/json;charset=utf-8", body)

	if err != nil {
		log.Debug(err.Error())
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// PostFile 上传文件
func PostFile(fieldname, filename, uri string) ([]byte, error) {
	fields := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: fieldname,
			Filename:  filename,
		},
	}
	return PostMultipartForm(fields, uri)
}

// MultipartFormField 保存文件或其他字段信息
type MultipartFormField struct {
	IsFile    bool
	Fieldname string
	Value     []byte
	Filename  string
}

// PostMultipartForm 上传文件或其他多个字段
func PostMultipartForm(fields []MultipartFormField, uri string) (respBody []byte, err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for _, field := range fields {
		if field.IsFile {
			fileWriter, e := bodyWriter.CreateFormFile(field.Fieldname, field.Filename)
			if e != nil {
				err = fmt.Errorf("error writing to buffer , err=%v", e)
				return
			}

			fh, e := os.Open(field.Filename)
			if e != nil {
				err = fmt.Errorf("error opening file , err=%v", e)
				return
			}
			defer fh.Close()

			if _, err = io.Copy(fileWriter, fh); err != nil {
				return
			}
		} else {
			partWriter, e := bodyWriter.CreateFormField(field.Fieldname)
			if e != nil {
				err = e
				return
			}
			valueReader := bytes.NewReader(field.Value)
			if _, err = io.Copy(partWriter, valueReader); err != nil {
				return
			}
		}
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, e := http.Post(uri, contentType, bodyBuf)
	if e != nil {
		err = e
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	respBody, err = ioutil.ReadAll(resp.Body)
	return
}

// PostXML perform a HTTP/POST request with XML body
func PostXML(uri string, obj interface{}) ([]byte, error) {
	xmlData, err := xml.Marshal(obj)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(xmlData)
	response, err := http.Post(uri, "application/xml;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// PostXML perform a HTTP/POST request with XML body
func PostXMLWithhttps(uri string, xmldata string, certFile, keyFile string) (r []byte, err error) {

	body := bytes.NewBuffer([]byte(xmldata))
	cliCrt, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Post(uri, "application/xml;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	r, err = ioutil.ReadAll(response.Body)
	return r, err
}

func ClientIp(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
