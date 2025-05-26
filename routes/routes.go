package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sglkc/golang-news-api/models"
)

func RootRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PONG!"))
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		var news models.News

		body := render.DecodeJSON(r.Body, news)

		render.JSON(w, r, body)
	})

	return r
}
