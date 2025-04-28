package middleware

import (
	"strings"
	"uooobarry/soup/internal/auth"
	"uooobarry/soup/internal/handler"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Error(handler.NewUnauthorizedError())
			return
		}

		authString := strings.SplitN(authHeader, " ", 2)
		if !(len(authString) == 2 && authString[0] == "Bearer") {
			err := handler.NewInvalidRequest(handler.WithCustomMessage("invalid token format"))
			c.Error(err)
			return
		}

		claims, err := auth.ParseToken(authString[1])
		if err != nil {
			c.Error(handler.NewUnauthorizedError())
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
