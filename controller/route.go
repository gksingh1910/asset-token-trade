package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)


func PlayerTokenRoute() http.Handler {
	r2 := chi.NewRouter()
	r2.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Getting Player Token data")) })
	return r2
}
