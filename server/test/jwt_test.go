package test

import (
	jwt2 "github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

func TestJWt(t *testing.T) {
	now := time.Now()
	expire := now.AddDate(0, 0, 1)
	accessJwtKey := []byte("go-vue-ts-secret")
	claims := jwt2.MapClaims{
		"user_id":     1,
		"login_time":  now.UnixMilli(),
		"expire_time": expire.UnixMilli(),
	}
	token := jwt2.NewWithClaims(jwt2.SigningMethodHS256, claims)
	signedString, err := token.SignedString(accessJwtKey)
	t.Log(signedString)
	_, err = jwt2.Parse(signedString, func(token *jwt2.Token) (interface{}, error) {
		return accessJwtKey, nil
	})
	t.Log(err)
}
