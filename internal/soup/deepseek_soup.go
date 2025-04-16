package soup

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"uooobarry/soup/internal/model"
	"uooobarry/soup/internal/service"

	"github.com/joho/godotenv"
)

func FetchSoupFromDS() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseUri := os.Getenv("DEEPSEEK_BASE_URI")
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	s := service.InitDS(baseUri, apiKey)

	systemPrompt := `用户会让你生成海龟汤，你将会把汤面汤底和标签转换成纯JSON格式
        EXAMPLE JSON OUTPUT:
        json
        {
            soup_question: '一个男人走进一家餐厅，点了一份“海龟汤”。他喝了一口汤，突然脸色大变，接着冲出餐厅，回家后自杀了。为什么？',
            tag: ['红汤', '本格'],
            soup_answear: '这个男人曾经遭遇海难，和同伴（可能是亲人或挚友）在海上漂流，濒临饿死。同伴为了让他活下去，谎称煮了“海龟汤”给他喝，但实际上是用自己的血肉熬制的。男人活了下来，但同伴牺牲了。多年后，他在餐厅点了真正的海龟汤，尝出味道完全不同，瞬间明白当年的真相，因无法承受巨大的愧疚与悲痛，选择自杀.'
        }
    `
	userPrompt := `**你是一位畅销推理小说作家兼经验丰富的海龟汤主持，你会提供一个丰富有趣同时逻辑严密的海龟汤. 
                     你会用清晰的语言给出汤面和汤底，并且告诉我谜题是否为红汤/白汤，本格/超格。
                   **规则：红汤表示故事中有人死去，白汤则相反。本格表示合乎正常逻辑，超格则包含超自然现象。
                   **提示：海龟汤需要逻辑严密的同时充满创意，不宜出现夫妻出轨之类的俗套狗血剧情。剧情需要完整不要突兀地结束。
                   **补充: 汤面与汤底不能相差过多的信息不能与汤底有过多的信息差。`

	systemMsg := []*service.DeepSeekMessage{
		{Role: "system", Content: systemPrompt},
	}
	userMsg := service.DeepSeekMessage{Role: "user", Content: userPrompt}
	rsp, err := s.Chat(&userMsg, systemMsg, service.SetModel("deepseek-chat"), service.SetResponseFmt("json_object"))
	if err != nil {
		return err
	}
	log.Print(rsp)

	// Unmarshal the response content into the Soup model
	var soup model.Soup
	if err := json.Unmarshal([]byte(rsp.Choices[0].Message.Content), &soup); err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if err := model.DB.Create(&soup).Error; err != nil {
		return fmt.Errorf("failed to store soup in database: %v", err)
	}

	return nil
}
