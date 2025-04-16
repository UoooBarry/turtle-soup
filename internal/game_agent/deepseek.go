package gameagent

import (
	"encoding/json"
	"log"
	"os"
	"uooobarry/soup/internal/model"
	"uooobarry/soup/internal/service"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type DeepSeekGameAgent struct {
	UUID        string
	Service     *service.DeepSeekService
	Soup        *model.Soup
	PerviousMsg []*service.DeepSeekMessage
}

type GameResponse struct {
	Question string `json:"question"`
	Answear  string `json:"answear"`
	Hint     string `json:"hint"`
	GameEnd  bool   `json:"gameend"`
}

func InitDS(soupID uint) (*DeepSeekGameAgent, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseUri := os.Getenv("DEEPSEEK_BASE_URI")
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	s := service.InitDS(baseUri, apiKey)

	var soup model.Soup
	if err := model.DB.First(&soup, soupID).Error; err != nil {
		return nil, err
	}
	return &DeepSeekGameAgent{Service: s, UUID: uuid.New().String()}, nil
}

func (agent *DeepSeekGameAgent) Start() error {
	systemPrompt := `你接下来将会根据用户提供的海龟汤谜题的答案来回答后面用户的提问，你只允许回答：是，不是，是或不是，不相关,
                     同时你会给出一些引导信息,
                     当故事大致已经被猜对时，gamnend字段为true, answear字段为完整汤底
        EXAMPLE JSON OUTPUI WHEN GAME START:
        json
        {
            question: "",
            answear: "开始游戏"
            gamened: false
        }
        EXAMPLE JSON OUTPUT WHEN USER ASK:
        json
        {
            question: "男子是被胁迫喝的海龟汤吗?"
            answear: "不是",
            hint: "可以考虑男子是否曾经喝过'海龟汤'才另男子情绪崩溃。",
            gamnend: false
        }
        EXAMPLE JSON OUT WHEN USER HAVE THE ANSWEAR
        json
        {
            question: "男人曾经被困在海上的时候，喝过亲近的人用血肉给他做的海龟汤，亲人告诉他，这是海龟做的汤。直到他获救到了餐厅，喝到了海龟汤，发现真正的海龟汤并不是当初喝到的味道。"
            answear: "这个男人曾经遭遇海难，和同伴（可能是亲人或挚友）在海上漂流，濒临饿死。同伴为了让他活下去，谎称煮了“海龟汤”给他喝，但实际上是用自己的血肉熬制的。男人活了下来，但同伴牺牲了。多年后，他在餐厅点了真正的海龟汤，尝出味道完全不同，瞬间明白当年的真相，因无法承受巨大的愧疚与悲痛，选择自杀。",
            hint: "*你猜对了！"
            gameend: true
        }
    `
	userPrompt := "**你是一位畅销推理小说作家兼经验丰富的海龟汤主持，你将根据提供的谜题和答案来主持这一局海龟汤游戏。\n" +
		"**规则：接下来用户会不知道谜底的情况下向你提问，你只允许回答：是，不是，是或者不是，不相关。当用户的提问比较模糊时，你允许纠正用户的提问。同时，当用户的提问比较接近真相时，你可以引导玩家往正确地方向猜，但是不宜太明显。\n" +
		"**海龟汤谜题:" + agent.Soup.SoupQuestion + "\n" +
		"**海龟汤谜底：" + agent.Soup.SoupAnswear + "\n" +
		"现在我们可以开始游戏了。"

	systemMsg := service.DeepSeekMessage{
		Role: "system", Content: systemPrompt,
	}
	userMsg := service.DeepSeekMessage{Role: "user", Content: userPrompt}
	agent.AppendMsg(&systemMsg)
	rsp, err := agent.Service.Chat(&userMsg, agent.PerviousMsg, service.SetModel("deepseek-chat"), service.SetResponseFmt("json_object"))
	if err != nil {
		return err
	}
	log.Print(rsp)
	agent.AppendMsg(&userMsg)
	agent.AppendMsg(&rsp.Choices[0].Message)

	return nil
}

func (agent *DeepSeekGameAgent) AppendMsg(msg *service.DeepSeekMessage) {
	agent.PerviousMsg = append(agent.PerviousMsg, msg)
}

func (agent *DeepSeekGameAgent) Ask(question string) (*GameResponse, error) {
	userMsg := service.DeepSeekMessage{Role: "user", Content: question}

	rsp, err := agent.Service.Chat(&userMsg, agent.PerviousMsg, service.SetModel("deepseek-chat"), service.SetResponseFmt("json_object"))
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
