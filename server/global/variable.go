package global

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	REDISTTL          = time.Minute * 10
	REDISSAVETOKENTTL = time.Minute * 60
)

var (
	TokenExpirationTime = jwt.NewNumericDate(time.Now().Add(60 * time.Minute))
)
