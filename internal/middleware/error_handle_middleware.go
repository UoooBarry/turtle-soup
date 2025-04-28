package middleware

import (
	"net/http"
	"uooobarry/soup/internal/handler"

	"github.com/gin-gonic/gin"
)

func ErrorHandleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := http.StatusInternalServerError
		message := "internal server error"

		if len(c.Errors) <= 0 {
			c.Next()
			return
		}
		err := c.Errors.Last()

		if sc, ok := err.Err.(handler.HandlerError); ok {
			status = sc.StatusCode()
		}
		if err.Err != nil {
			message = err.Error()
		}

		c.JSON(status, gin.H{
			"error": message,
		})
		c.Abort()
	}
}
