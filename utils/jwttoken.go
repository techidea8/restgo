package utils

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenToken 生成token
func GenerateToken(values map[string]interface{}, secretKey string) (string, error) {
	claim := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
	}
	for key, value := range values {
		claim[key] = value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// VerfyToken 验证token
func ParseToken(token, secretKey string) (map[string]interface{}, error) {
	tokenObj, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokenObj.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot convert claim to mapclaim")
	}

	if !tokenObj.Valid {
		return nil, errors.New("token is invalid")
	}
	return claims, nil
}
