package config

import (
	"net/http"
	"uooobarry/soup/internal/handler"
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
	r.Use(middleware.ErrorHandleMiddleware())

	soupRepo := repository.NewSoupRepository(db)
	soupService := service.NewSoupService(soupRepo)
	soupHandler := handler.NewSoupHandler(soupService)

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	gameHandler := handler.NewGameHandler(soupService, authService)

	public := r.Group("/")
	{
		public.GET("/soups", soupHandler.ListSoups)
		public.POST("/auth/login", authHandler.Login)
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	protected := r.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/auth/profile", authHandler.Profile)

		game := protected.Group("/game")
		{
			game.POST("/create", gameHandler.CreateGame)
			game.POST("/:uuid/start", gameHandler.StartGame)
			game.POST("/:uuid/ask", gameHandler.GameAskQuestion)
			game.DELETE("/:uuid/end", gameHandler.EndGame)
		}
	}
}
