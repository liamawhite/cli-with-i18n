package i18n

import "github.com/liamawhite/cli-with-i18n/util/ui"

var T ui.TranslateFunc

type LocaleReader interface {
	Locale() string
}

func Init(config LocaleReader) ui.TranslateFunc {
	t, _ := ui.GetTranslationFunc(config)
	return t
}
