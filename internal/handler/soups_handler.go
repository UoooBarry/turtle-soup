package handler

import (
	"net/http"
	"uooobarry/soup/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type SoupResponse struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	SoupQuestion string         `json:"soup_question"`
	SoupTag      datatypes.JSON `json:"tag" gorm:"type:text"`
	Enabled      bool           `json:"enabled" gorm:"default:true"`
}

type SoupHandler struct {
	service *service.SoupService
}

func NewSoupHandler(service *service.SoupService) *SoupHandler {
	return &SoupHandler{service: service}
}

func (h *SoupHandler) ListSoups(c *gin.Context) {
	soups, err := h.service.List()
	if err != nil {
		NewNotfoundError()
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
