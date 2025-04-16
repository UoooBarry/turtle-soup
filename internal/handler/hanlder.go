package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/users", CreateUser)
	r.GET("/users", ListUsers)
	r.GET("/soups", ListSoups)

	r.POST("/game/create", CreateGame)
	r.POST("/game/start", StartGame)
	r.POST("/game/ask", GameAskQuestion)
}
