package util

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var privateKey []byte = []byte("private key")

func Encrypt(mapClaims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println("db Encrypt error:", err)
		return ""
	} else {
		return tokenString
	}
}

func Decrypt(tokenString string) (jwt.MapClaims, error) {
	if len(tokenString) < 2 {
		return nil, errors.New("token to short")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
