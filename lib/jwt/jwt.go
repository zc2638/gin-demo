package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func Create(info map[string]interface{}, secret string, exp int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"info": info,
		"exp":  exp, // 过期时间
	})
	return token.SignedString([]byte(secret))
}

func Parse(tokenStr string, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func ParseInfo(tokenStr string, secret string) (map[string]interface{}, error) {
	token, err := Parse(tokenStr, secret)
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

func CheckValid(tokenStr string, secret string) bool {
	token, err := Parse(tokenStr, secret)
	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}
	return true
}
