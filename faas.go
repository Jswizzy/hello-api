package faas

import (
	"github.com/jswizzy/hello-api/handlers/rest"
	"net/http"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
