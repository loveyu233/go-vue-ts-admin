package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"server/global"
	"time"
)

type Claims struct {
	UserId int
	jwt.RegisteredClaims
}

func CreateToken(id int) (string, error) {
	accessJwtKey := []byte(viper.GetString("token.secret"))
	claims := Claims{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: global.TokenExpirationTime,
			Issuer:    "233",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString(accessJwtKey)
}

func GetPayload(token string) (int, error) {
	parser := jwt.NewParser()
	var claims Claims
	_, _, err := parser.ParseUnverified(token, &claims)
	return claims.UserId, err
}

func VerifyToken(token string) error {
	accessJwtKey := []byte(viper.GetString("token.secret"))
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return accessJwtKey, nil
	})
	return err
}
