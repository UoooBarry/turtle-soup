package handler

import (
	"net/http"
	"uooobarry/soup/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(NewInvalidRequest())
		return
	}

	user, err := h.service.Register(req.Username, req.Password, req.Email)
	if err != nil {
		c.Error(NewInternalError(WithCustomMessage(err.Error())))
		return
	}

	c.JSON(http.StatusCreated, user)
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(NewInvalidRequest(WithCustomMessage(err.Error())))
		return
	}

	user, token, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		c.Error(NewUnauthorizedError())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"user_id":  user.ID,
		"username": user.Username,
	})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.Error(NewUnauthorizedError())
		return
	}

	user, err := h.service.GetUserByID(userID.(uint))
	if err != nil {
		c.Error(NewInternalError())
		return
	}

	c.JSON(http.StatusOK, user)
}
