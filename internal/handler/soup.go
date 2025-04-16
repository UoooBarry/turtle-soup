package handler

import (
	"net/http"
	"uooobarry/soup/internal/model"

	"github.com/gin-gonic/gin"
)

func ListSoups(c *gin.Context) {
	var soups []model.Soup
	if err := model.DB.Find(&soups).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, soups)
}
