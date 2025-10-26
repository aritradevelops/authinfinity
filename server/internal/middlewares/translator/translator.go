package translator

import (
	"fmt"
	"os"
	"path"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

func New() fiber.Handler {
	cwd, _ := os.Getwd()
	localesPath := path.Join(cwd, "locales")
	fmt.Println(localesPath)
	return fiberi18n.New(&fiberi18n.Config{
		RootPath:         localesPath,
		AcceptLanguages:  []language.Tag{language.English, language.Bengali},
		DefaultLanguage:  language.English,
		UnmarshalFunc:    yaml.Unmarshal,
		FormatBundleFile: "yaml",
	})
}

func Localize(c *fiber.Ctx, id string, data ...any) string {
	msg, err := fiberi18n.Localize(c, &i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data[0],
	})
	if err != nil {
		// TODO: log notice level
		return id
	}
	return msg
}
