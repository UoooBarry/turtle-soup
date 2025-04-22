package middleware

import (
	"uooobarry/soup/internal/i18n"

	"github.com/gin-gonic/gin"
)

func I18nMildware(i18n *i18n.I18nHelper) gin.HandlerFunc {
	return func(c *gin.Context) {
		if lng := c.GetHeader("Accept-Language"); lng != "" {
			c.Set("lang", lng)
		}
		if lng := c.Query("lang"); lng != "" {
			c.Set("lang", lng)
		}

		c.Set("i18n", i18n)
		c.Next()
	}
}
