package handler

import (
	"log"
	"net/http"
	gameagent "uooobarry/soup/internal/game_agent"
	"uooobarry/soup/internal/i18n"
	"uooobarry/soup/internal/service"

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
	NeedHint bool   `json:"need_hint"`
}

type GameHandler struct {
	*AuthenticatedHandler
	soupService *service.SoupService
}

func NewGameHandler(s *service.SoupService, a *service.AuthService) *GameHandler {
	return &GameHandler{
		soupService:          s,
		AuthenticatedHandler: &AuthenticatedHandler{authService: a},
	}
}

func (handler *GameHandler) CreateGame(c *gin.Context) {
	var req CreateGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(NewInternalError())
		return
	}

	helper := c.MustGet("i18n").(*i18n.I18nHelper)
	localizer := helper.GetLocalizer(c)
	user := handler.AuthenticatedHandler.GetCurrentUser(c)
	session, err := gameagent.NewSession(req.SoupID, user, handler.soupService, localizer)
	if err != nil {
		c.Error(NewInternalError())
		return
	}

	resp := CreateGameResponse{UUID: session.Agent.(*gameagent.DeepSeekGameAgent).UUID}
	c.JSON(http.StatusOK, resp)
}

func (handler *GameHandler) StartGame(c *gin.Context) {
	_, session := handler.getCurrentSession(c)
	if err := session.Agent.Start(); err != nil {
		log.Println(err)
		c.Error(NewInternalError())
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (handler *GameHandler) GameAskQuestion(c *gin.Context) {
	var req AskQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Invalid request:", err.Error())
		c.Error(NewInvalidRequest())
		return
	}

	uuid, session := handler.getCurrentSession(c)

	rsp, err := session.Agent.Ask(req.Question, req.NeedHint)
	if err != nil {
		log.Println("Failed to process question:", err.Error())
		c.Error(NewInternalError())
		return
	}
	if rsp.GameEnd == true {
		gameagent.EndSession(uuid)
	}

	c.JSON(http.StatusOK, rsp)
}

func (handler *GameHandler) EndGame(c *gin.Context) {
	uuid, _ := handler.getCurrentSession(c)
	gameagent.EndSession(uuid)
	c.JSON(http.StatusOK, nil)
}

func (handler *GameHandler) getCurrentSession(c *gin.Context) (string, *gameagent.SessionInfo) {
	uuid := c.Param("uuid")
	session, exist := gameagent.GetSession(uuid)
	if !exist {
		log.Println("Session not found for UUID:", uuid)
		c.Error(NewNotfoundError(WithCustomMessage("session not found")))
		return "", nil
	}

	return uuid, session
}
