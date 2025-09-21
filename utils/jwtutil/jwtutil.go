package jwtutil

import (
	"errors"
	"fmt"
	"scaffold-demo/config"
	"scaffold-demo/utils/logs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSignKey = []byte(config.JwtSignKey)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(username string) (string, error) {
	claims := MyCustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.JwtExpTime) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "llkk",
			Subject:   "subRabbit",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(JwtSignKey)
	return ss, err
}

func ParseToken(tokenStr string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (any, error) {
		return JwtSignKey, nil
	})

	if err != nil {
		logs.Error(nil, err.Error())
		return nil, err
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		logs.Info(nil, "获取了token结构体")
		fmt.Printf("获取的username字段是:%s\n", claims.Username)
		return claims, nil
	} else {
		logs.Warn(nil, "无法解析token")
		fmt.Printf("token.Claims: %v\n", claims)
		return nil, errors.New("无法解析token")
	}
}
