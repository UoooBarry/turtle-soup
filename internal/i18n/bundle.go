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
	lang, exist := c.Get("i18n")
	if !exist {
		lang = language.Chinese.String()
	}

	if tag, ok := lang.(language.Tag); ok {
		return i18n.NewLocalizer(helper.Bundle, tag.String())
	}

	return i18n.NewLocalizer(helper.Bundle, language.Chinese.String())
}
