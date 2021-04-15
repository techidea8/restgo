package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	Secret                 = "turingapp" // 加盐
	ExpireTime             = 3600 * 72   // token有效期
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
	TokenStr               = "X-Token"
)

//中间件
func JwtAuth() Middleware {
	return func(w http.ResponseWriter, req *http.Request) (bool, error) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := req.Header.Get("x-token")

		if token == "" {
			return false, errors.New("你没有权限进行该操作")
		}
		_, err := parseToken(token)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

//创建token
func CreateToken(userId uint) (string, error) {
	claims := &jWTClaims{
		UserId: userId,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	return getToken(claims)
}

//校验token
func RefreshToken(strToken string) (string, error) {

	claims, err := parseToken(strToken)
	if err != nil {
		return "", err
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := getToken(claims)
	return signedToken, err
}
//解析token
func ParseToken(token string) (userId uint, err error) {
	return parseUserId(token)
}
//解析token
func CurrentUserId(req *http.Request) (userId uint, err error) {
	strToken := req.Header.Get("X-token")
	return parseUserId(strToken)
}

//解析token
func parseUserId(strToken string) (userId uint, err error) {
	claims, err := parseToken(strToken)
	if err != nil {
		return 0, err
	}
	return claims.UserId, nil
}

type jWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

func parseToken(strToken string) (*jWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &jWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy)
	}
	claims, ok := token.Claims.(*jWTClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	return claims, nil
}

func getToken(claims *jWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorReason_ServerBusy)
	}
	return signedToken, nil
}
