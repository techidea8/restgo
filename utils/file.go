package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"os/exec"

	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
)

const Seprate = "/"

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func FileExist(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(info)
			return false
		}
	}
	return true
}
func FileCompare(filename string, size int64) bool {
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(info)
			return false
		}
	}
	return size == info.Size()
}

/**
 * @Description: 获取文件创建时间
 * @param path
 * @return int64
 */
func GetFileCreateTime(path string) int64 {

	return time.Now().Unix()
}
func Exec(cmdrun string, params ...string) (string, error) {
	var outInfo, outErr bytes.Buffer
	cmd := exec.Command(cmdrun, params...)
	cmd.Stdout = &outInfo
	cmd.Stderr = &outErr
	if err := cmd.Run(); err != nil {
		garbledStr := ConvertByte2String(outInfo.Bytes(), GB18030)
		return garbledStr, errors.New(ConvertByte2String(outErr.Bytes(), GB18030))
	}
	return ConvertByte2String(outInfo.Bytes(), GB18030), nil
}

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func SavePic(b64 string, dir, filename string) (string, error) {

	i := strings.Index(b64, ",")
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64[i+1:]))
	picName := filename
	f, err := os.Create(dir + picName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, dec)
	if err != nil {
		return "", err
	}
	return dir + picName, nil
}
