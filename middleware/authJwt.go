package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"web-gin/common/pkg"
	gin_jwt "web-gin/common/pkg/jwt"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := pkg.SUCCESS
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			code = pkg.ERROR_AUTH_CHECK_TOKEN_PARSE
		} else {
			token := tokenString[7:]
			claims, err := gin_jwt.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = pkg.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = pkg.ERROR_AUTH_CHECK_TOKEN_PARSE
				}
			}
			c.Set("user_service", claims)
		}
		if code != pkg.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "returnCode": code, "msg": pkg.GetMsg(code), "data": nil,})
			c.Abort()
			return
		}
		c.Next()
	}
}
