package http

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/st3v/translator/microsoft"
)

// RepoMock ...
type RepoMock struct {
}

// NewRepo ...
func NewRepo() *RepoMock {
	return &RepoMock{}
}

// Close ...
func (r *RepoMock) Close() error {
	return nil
}

// Translate ...
func (r *RepoMock) Translate(srcLang string, desLang string, text string) string {
	translator := microsoft.NewTranslator(viper.GetString("API_KEY"))

	translation, err := translator.Translate(text, srcLang, desLang)
	if err != nil {
		log.Panicf("Error during translation: %s", err.Error())
	}

	fmt.Println(translation)

	return translation
}
