package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"time"
)

type Claims struct {
	UserId int
	jwt.RegisteredClaims
}

func CreateToken(id int) (string, error) {
	now := time.Now()
	expire := now.AddDate(0, 0, 1)
	accessJwtKey := []byte("go-vue-ts-secret")
	claims := jwt.MapClaims{
		"user_id":     id,
		"login_time":  now.UnixMilli(),
		"expire_time": expire.UnixMilli(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessJwtKey)
}

func GetPayload(token string) (int, error) {
	parse, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	claims, ok := parse.Claims.(jwt.MapClaims)
	if !ok {
		return -1, errors.New("token parse err")
	}
	return cast.ToInt(claims["user_id"]), nil
}

func VerifyToken(token string) error {
	accessJwtKey := []byte(viper.GetString("token.secret"))
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return accessJwtKey, nil
	})
	return err
}
