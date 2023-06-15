package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/cache"
	"server/domain/response"
	"server/domain/valid"
	"server/jwt"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		err := jwt.VerifyToken(token)
		if err != nil {
			log.Println(err)
			c.JSON(200, response.NewErrorResponse(valid.TokenValidationFailed))
			c.Abort()
			return
		}
		uid, _ := jwt.GetPayload(token)
		if !cache.VerifyTokenIsExist(uid, token) {
			c.JSON(200, response.NewErrorResponse(valid.TokenDoesNotExist))
			c.Abort()
			return
		}
		c.Set("uid", uid)
	}
}
