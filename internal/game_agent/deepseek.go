package gameagent

import (
	"encoding/json"

	"fmt"
	"os"
	"uooobarry/soup/internal/client"
	"uooobarry/soup/internal/model"
	"uooobarry/soup/internal/service"

	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type (
	ErrAgentInitialization struct {
		Reason string
	}

	ErrGameLogic struct {
		Reason string
	}

	ErrAPIRequest struct {
		Reason string
	}
)

func (e ErrAgentInitialization) Error() string {
	return fmt.Sprintf("agent initialization failed: %s", e.Reason)
}

func (e ErrGameLogic) Error() string {
	return fmt.Sprintf("game logic error: %s", e.Reason)
}

type DeepSeekGameAgent struct {
	UUID        string
	client      *client.DeepSeekClient
	Soup        *model.Soup
	PerviousMsg []*client.DeepSeekMessage
	Localizer   *i18n.Localizer
}

type GameResponse struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Hint     string `json:"hint"`
	GameEnd  bool   `json:"gameend"`
}

func InitDS(soupID uint, service *service.SoupService, localizer *i18n.Localizer) (*DeepSeekGameAgent, error) {
	baseUri := os.Getenv("DEEPSEEK_BASE_URI")
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	s := client.InitDS(baseUri, apiKey)

	soup, err := service.GetByID(soupID)
	if err != nil {
		return nil, ErrAgentInitialization{Reason: fmt.Sprintf("failed to fetch soup: %v", err)}
	}
	return &DeepSeekGameAgent{client: s,
		UUID:      uuid.New().String(),
		Soup:      soup,
		Localizer: localizer}, nil
}

func (agent *DeepSeekGameAgent) Start() error {
	if agent.Soup == nil {
		return ErrGameLogic{Reason: "no soup is set to this agent"}
	}
	systemPrompt, err := agent.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "systemPrompt",
	})
	if err != nil {
		return err
	}

	userPrompt, err := agent.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    "userPrompt",
		TemplateData: map[string]string{"soupQuestion": agent.Soup.SoupQuestion},
	})
	if err != nil {
		return err
	}

	systemMsg := client.DeepSeekMessage{
		Role: "system", Content: systemPrompt,
	}
	userMsg := client.DeepSeekMessage{Role: "user", Content: userPrompt}
	agent.AppendMsg(&systemMsg)
	rsp, err := agent.client.Chat(&userMsg, agent.PerviousMsg, client.SetModel("deepseek-chat"), client.SetResponseFmt("json_object"))
	if err != nil {
		return err
	}
	agent.AppendMsg(&userMsg)
	agent.AppendMsg(&rsp.Choices[0].Message)

	return nil
}

func (agent *DeepSeekGameAgent) AppendMsg(msg *client.DeepSeekMessage) {
	agent.PerviousMsg = append(agent.PerviousMsg, msg)
}

func (agent *DeepSeekGameAgent) Ask(question string, needHint bool) (*GameResponse, error) {
	question += needHintPrompt(needHint)
	userMsg := client.DeepSeekMessage{Role: "user", Content: question}

	var rsp *client.DeepSeekResponse
	var err error
	maxRetries := 3

	for i := 0; i < maxRetries; i++ {
		rsp, err = agent.client.Chat(&userMsg,
			agent.PerviousMsg,
			client.SetModel("deepseek-chat"),
			client.SetResponseFmt("json_object"))
		if err == nil {
			break
		}

		if _, ok := err.(client.ErrAPIResponse); !ok {
			return nil, err // Non-retryable error, return immediately
		}

		if i == maxRetries-1 {
			return nil, err // Exhausted retries, return the error
		}
	}

	agent.AppendMsg(&userMsg)
	agent.AppendMsg(&rsp.Choices[0].Message)

	var gameResponse GameResponse
	if err := json.Unmarshal([]byte(rsp.Choices[0].Message.Content), &gameResponse); err != nil {
		return nil, err
	}
	return &gameResponse, nil
}

func needHintPrompt(needHint bool) string {
	return fmt.Sprintf("<NeedHint>%v</NeedHint>", needHint)
}
