package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var stSignKey = []byte("abcdefghijklmnopqrstuvwxyz")

const (
	// TokenType Token 类型
	TokenType = "bearer"
)

type CustomClaims struct {
	jwt.MapClaims
	StandardClaims jwt.StandardClaims
}

func (j CustomClaims) Valid() error {
	return nil
}

func IsTokenValid(tokenStr string) (int64, bool) {

	token, err := parseToken(tokenStr)
	if err != nil {
		return 400, false
	}

	// 校验过期时间
	ok := token.StandardClaims.VerifyExpiresAt(time.Now().Unix(), false)
	if !ok {
		return 401, false
	}

	return 200, true

}

// ParseToken 解析token
func parseToken(tokenStr string) (CustomClaims, error) {

	iJwtCustomClaims := CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})

	if err == nil && !token.Valid {
		err = errors.New("invalid Token")
	}
	return iJwtCustomClaims, err

}
