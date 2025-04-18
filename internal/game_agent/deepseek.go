package gameagent

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"uooobarry/soup/internal/client"
	"uooobarry/soup/internal/model"
	"uooobarry/soup/internal/service"

	"github.com/google/uuid"
)

type DeepSeekGameAgent struct {
	UUID        string
	client      *client.DeepSeekClient
	Soup        *model.Soup
	PerviousMsg []*client.DeepSeekMessage
}

type GameResponse struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Hint     string `json:"hint"`
	GameEnd  bool   `json:"gameend"`
}

func InitDS(soupID uint, service *service.SoupService) (*DeepSeekGameAgent, error) {
	baseUri := os.Getenv("DEEPSEEK_BASE_URI")
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	s := client.InitDS(baseUri, apiKey)

	soup, err := service.GetByID(soupID)
	if err != nil {
		return nil, err
	}
	return &DeepSeekGameAgent{client: s, UUID: uuid.New().String(), Soup: soup}, nil
}

func (agent *DeepSeekGameAgent) Start() error {
	if agent.Soup == nil {
		return errors.New("No soup is set to this agent.")
	}
	systemPrompt := `你是一位畅销推理小说作家兼经验丰富的海龟汤主持，你将根据提供的谜题和答案来主持这一局海龟汤游戏。\n" +
                     **规则：接下来用户会不知道谜底的情况下向你提问，你只允许回答：是，不是，是或者不是，不相关。当用户的提问比较模糊时，你允许纠正用户的提问。同时，当用户的提问比较接近真相时，你可以引导玩家往正确地方向猜，但是不宜太明显。
                     你接下来将会根据用户提供的海龟汤谜题的答案来回答后面用户的提问，你只允许回答：是，不是，是或不是，不相关,
                     同时你会给出一些引导信息,
                     当故事大致已经被猜对时，gamnend字段为true, answer字段为完整汤底
        EXAMPLE JSON OUTPUI WHEN GAME START:
        json
        {
            question: "",
            answer: "开始游戏"
            gamened: false
        }
        EXAMPLE JSON OUTPUT WHEN USER ASK:
        json
        {
            question: "男子是被胁迫喝的海龟汤吗?"
            answer: "不是",
            hint: "可以考虑男子是否曾经喝过'海龟汤'才另男子情绪崩溃。",
            gamnend: false
        }
        EXAMPLE JSON OUT WHEN USER HAVE THE answer
        json
        {
            question: "男人曾经被困在海上的时候，喝过亲近的人用血肉给他做的海龟汤，亲人告诉他，这是海龟做的汤。直到他获救到了餐厅，喝到了海龟汤，发现真正的海龟汤并不是当初喝到的味道。"
            answer: "这个男人曾经遭遇海难，和同伴（可能是亲人或挚友）在海上漂流，濒临饿死。同伴为了让他活下去，谎称煮了“海龟汤”给他喝，但实际上是用自己的血肉熬制的。男人活了下来，但同伴牺牲了。多年后，他在餐厅点了真正的海龟汤，尝出味道完全不同，瞬间明白当年的真相，因无法承受巨大的愧疚与悲痛，选择自杀。",
            hint: "*你猜对了！"
            gameend: true
        }
    `
	userPrompt := "<海龟汤谜题>" + agent.Soup.SoupQuestion + "</海龟汤谜题>\n" +
		"<海龟汤谜底>" + agent.Soup.SoupAnswer + "</海龟汤谜底>"

	systemMsg := client.DeepSeekMessage{
		Role: "system", Content: systemPrompt,
	}
	userMsg := client.DeepSeekMessage{Role: "user", Content: userPrompt}
	agent.AppendMsg(&systemMsg)
	rsp, err := agent.client.Chat(&userMsg, agent.PerviousMsg, client.SetModel("deepseek-chat"), client.SetResponseFmt("json_object"))
	if err != nil {
		return err
	}
	log.Print(rsp)
	agent.AppendMsg(&userMsg)
	agent.AppendMsg(&rsp.Choices[0].Message)

	return nil
}

func (agent *DeepSeekGameAgent) AppendMsg(msg *client.DeepSeekMessage) {
	agent.PerviousMsg = append(agent.PerviousMsg, msg)
}

func (agent *DeepSeekGameAgent) Ask(question string) (*GameResponse, error) {
	userMsg := client.DeepSeekMessage{Role: "user", Content: question}

	rsp, err := agent.client.Chat(&userMsg, agent.PerviousMsg, client.SetModel("deepseek-chat"), client.SetResponseFmt("json_object"))
	if err != nil {
		return nil, err
	}
	agent.AppendMsg(&userMsg)
	agent.AppendMsg(&rsp.Choices[0].Message)

	var gameResponse GameResponse
	if err := json.Unmarshal([]byte(rsp.Choices[0].Message.Content), &gameResponse); err != nil {
		return nil, err
	}
	return &gameResponse, nil
}
