package jwtutil

import (
	"scaffold-demo/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSignKey = []byte(config.JwtSignKey)

type MyCustomClaims struct {
	Username string `json:"Username"`
	jwt.RegisteredClaims
}

func GenToken(username string) (string, error) {
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
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
