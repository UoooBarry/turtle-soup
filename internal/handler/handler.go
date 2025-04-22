package handler

import (
	"net/http"
	"uooobarry/soup/internal/i18n"
	"uooobarry/soup/internal/middleware"
	"uooobarry/soup/internal/repository"
	"uooobarry/soup/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	i18nHelper := i18n.NewI18nHelper()
	r.Use(middleware.I18nMildware(i18nHelper))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	soupRepo := repository.NewSoupRepository(db)
	soupService := service.NewSoupService(soupRepo)
	soupHandler := NewSoupHandler(soupService)
	r.GET("/soups", soupHandler.ListSoups)

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := NewAuthHandler(authService)
	r.POST("/auth/login", authHandler.Login)

	authGroup := r.Group("/")
	authGroup.Use(middleware.JWTAuthMiddleware())
	{
		gameHandler := NewGameHandler(soupService)
		authGroup.POST("/game/create", gameHandler.CreateGame)
		authGroup.POST("/game/:uuid/start", gameHandler.StartGame)
		authGroup.POST("/game/:uuid/ask", gameHandler.GameAskQuestion)
		authGroup.DELETE("/game/:uuid/end", gameHandler.EndGame)
		authGroup.GET("/auth/profile", authHandler.Profile)
	}
}
