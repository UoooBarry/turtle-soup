package i18n

import (
	"embed"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

//go:embed gameagent.zh.toml gameagent.en.toml
var fs embed.FS

type I18nHelper struct {
	Bundle *i18n.Bundle
}

func NewI18nHelper() *I18nHelper {
	b := i18n.NewBundle(language.Chinese)
	b.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	if _, err := b.LoadMessageFileFS(fs, "gameagent.zh.toml"); err != nil {
		log.Fatal(err)
	}
	if _, err := b.LoadMessageFileFS(fs, "gameagent.en.toml"); err != nil {
		log.Fatal(err)
	}

	return &I18nHelper{Bundle: b}
}

func (helper *I18nHelper) GetLocalizer(c *gin.Context) *i18n.Localizer {
	lang, exist := c.Get("lang")
	if !exist {
		return defaultLocalizer(helper.Bundle)
	}

	langStr, ok := lang.(string)
	if !ok {
		return defaultLocalizer(helper.Bundle)
	}

	switch langStr {
	case "zh", "en":
		return i18n.NewLocalizer(helper.Bundle, langStr)
	default:
		return defaultLocalizer(helper.Bundle)
	}
}

func defaultLocalizer(b *i18n.Bundle) *i18n.Localizer {
	return i18n.NewLocalizer(b, language.Chinese.String())
}
