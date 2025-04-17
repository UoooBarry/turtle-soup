package handler

import (
	"log"
	"net/http"
	gameagent "uooobarry/soup/internal/game_agent"

	"github.com/gin-gonic/gin"
)

type CreateGameRequest struct {
	SoupID uint `json:"soup_id" binding:"required"`
}

type CreateGameResponse struct {
	UUID string `json:"uuid"`
}

type AskQuestionRequest struct {
	Question string `json:"question" binding:"required"`
}

func CreateGame(c *gin.Context) {
	var req CreateGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	agent, err := gameagent.NewSession(req.SoupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create the game"})
		return
	}

	resp := CreateGameResponse{UUID: agent.(*gameagent.DeepSeekGameAgent).UUID}
	c.JSON(http.StatusOK, resp)
}

func StartGame(c *gin.Context) {
	uuid := c.Param("uuid")

	agent, exist := gameagent.GetSession(uuid)
	if !exist {
		log.Println("Session not found for UUID:", uuid)
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	if err := agent.Start(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to start game"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func GameAskQuestion(c *gin.Context) {
	var req AskQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Invalid request:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	uuid := c.Param("uuid")
	agent, exists := gameagent.GetSession(uuid)
	if !exists {
		log.Println("Session not found for UUID:", uuid)
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	rsp, err := agent.Ask(req.Question)
	if err != nil {
		log.Println("Failed to process question:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process question"})
		return
	}

	c.JSON(http.StatusOK, rsp)
}

func EndGame(c *gin.Context) {
	uuid := c.Param("uuid")
	_, exist := gameagent.GetSession(uuid)
	if !exist {
		log.Println("Session not found for UUID:", uuid)
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	gameagent.EndSession(uuid)
	c.JSON(http.StatusOK, nil)
}
