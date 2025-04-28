package handler

import (
	"net/http"
	"uooobarry/soup/internal/model"
	"uooobarry/soup/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthenticatedHandler struct {
	authService *service.AuthService
}

func (h *AuthenticatedHandler) GetCurrentUser(c *gin.Context) *model.User {
	user_id, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
	user, err := h.authService.GetUserByID(user_id.(uint))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

	return user
}
