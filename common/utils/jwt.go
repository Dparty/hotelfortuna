package utils

import (
	"errors"
	"fmt"
	"god-of-wealth/common/config"

	"github.com/golang-jwt/jwt"
)

var secret []byte

func init() {
	secret = []byte(config.GetString("jwt.secret"))
}

type Claims struct {
	jwt.StandardClaims
	ID string `json:"id"`
}

func NewClaims(id string, expiredAt int64) Claims {
	var c Claims
	c.ID = id
	c.ExpiresAt = expiredAt
	return c
}

func SignJwt(id string, expiredAt int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaims(id, expiredAt))
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}

func VerifyJwt(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if token == nil {
		return nil, errors.New("unvalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("unvalid token")
	}
}
