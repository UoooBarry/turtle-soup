package handler

import (
	"net/http"
	"uooobarry/soup/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type SoupResponse struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	SoupQuestion string         `json:"soup_question"`
	SoupTag      datatypes.JSON `json:"tag" gorm:"type:text"`
	Enabled      bool           `json:"enabled" gorm:"default:true"`
}

func ListSoups(c *gin.Context) {
	var soups []model.Soup
	if err := model.DB.Find(&soups).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Map to exclude soup_answer
	var publicSoups []SoupResponse
	for _, soup := range soups {
		publicSoups = append(publicSoups, SoupResponse{
			ID:           soup.ID,
			SoupQuestion: soup.SoupQuestion,
			SoupTag:      soup.SoupTag,
			Enabled:      soup.Enabled,
		})
	}

	c.JSON(http.StatusOK, publicSoups)
}
