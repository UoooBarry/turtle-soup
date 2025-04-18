package client

import (
	"encoding/json"
	"fmt"
	"log"

	"resty.dev/v3"
)

type DeepSeekClient struct {
	BaseUri string
	ApiKey  string
	Client  *resty.Client
}

type DeepSeekResponse struct {
	ID      string           `json:"id"`
	Object  string           `json:"object"`
	Created int64            `json:"created"`
	Model   string           `json:"model"`
	Choices []DeepSeekChoice `json:"choices"`
	Usage   DeepSeekUsage    `json:"usage"`
}

type DeepSeekChoice struct {
	Index        int16           `json:"index"`
	Message      DeepSeekMessage `json:"message"`
	Logprobs     interface{}     `json:"logprobs"` // Can be null or a struct
	FinishReason string          `json:"finish_reason"`
}

// DeepSeekMessage is the message type used for API requests.
type DeepSeekMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type DeepSeekUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type DeepSeekRequestBody struct {
	Model          string                 `json:"model"`
	Messages       []DeepSeekMessage      `json:"messages"`
	ResponseFormat DeepSeekResponseFormat `json:"response_format"`
}

type DeepSeekResponseFormat struct {
	Type string `json:"type"`
}

func InitDS(baseUri string, apiKey string) *DeepSeekClient {
	return &DeepSeekClient{BaseUri: baseUri, ApiKey: apiKey, Client: resty.New()}
}

func (ds *DeepSeekClient) R() *resty.Request {
	return ds.Client.R().SetHeader("Accept", "application/json").SetHeader("Authorization", fmt.Sprintf("Bearer %s", ds.ApiKey))
}

type ChatOption func(*ChatConfig)

type ChatConfig struct {
	Model          string
	ResponseFormat string
}

func SetModel(model string) ChatOption {
	return func(cfg *ChatConfig) {
		cfg.Model = model
	}
}

func SetResponseFmt(fmt string) ChatOption {
	return func(cfg *ChatConfig) {
		cfg.ResponseFormat = fmt
	}
}

func (ds *DeepSeekClient) Chat(question *DeepSeekMessage, prevs []*DeepSeekMessage, opts ...ChatOption) (*DeepSeekResponse, error) {
	// Config optional arugments
	cfg := ChatConfig{
		Model:          "deepseek-chat",
		ResponseFormat: "text",
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	// Build request body
	messages := make([]DeepSeekMessage, 0, len(prevs)+1)
	for _, prev := range prevs {
		messages = append(messages, *prev)
	}
	messages = append(messages, *question)
	for _, msg := range messages {
		log.Println(msg)
	}

	requestBody := DeepSeekRequestBody{
		Messages: messages,
		Model:    cfg.Model,
		ResponseFormat: DeepSeekResponseFormat{
			Type: cfg.ResponseFormat,
		},
	}

	// Send the request with the body
	resp, err := ds.R().
		SetBody(requestBody).
		Post(fmt.Sprintf("%s/chat/completions", ds.BaseUri))
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return nil, fmt.Errorf("request status error: %s, body: %s", resp.Status(), resp.String())
	}
	log.Print(resp.String())

	var dsr DeepSeekResponse
	if err := json.NewDecoder(resp.Body).Decode(&dsr); err != nil {
		return nil, fmt.Errorf("parsed JSON failed: %v", err)
	}

	return &dsr, nil
}
