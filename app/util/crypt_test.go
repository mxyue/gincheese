package util

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func TestEncrypt(t *testing.T) {
	mapClaims := jwt.MapClaims{"foo": "bar"}
	tokenString := Encrypt(mapClaims)
	if tokenString != "" {
		t.Log("pass")
	} else {
		t.Error("token 生成错误")
	}
}

func TestDecrypt(t *testing.T) {
	mapClaims := jwt.MapClaims{"foo": "bar"}
	tokenString := Encrypt(mapClaims)
	claims, err := Decrypt(tokenString)
	if err != nil {
		t.Error("decrypt失败:", err)
	} else if claims["foo"] == "bar" {
		t.Log("pass")
	}
}
