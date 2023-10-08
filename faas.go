// Package faas is used for function definitions
package faas

import (
	"github.com/jswizzy/hello-api/handlers/rest"
	"github.com/jswizzy/hello-api/translation"
	"net/http"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)

	translateHandler.TranslateHandler(w, r)
}
