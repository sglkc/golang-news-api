package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sglkc/golang-news-api/controllers"
)

func NewsRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", controllers.ListNews)
	r.Get("/:id", controllers.GetNews)
	r.Post("/", controllers.CreateNews)
	r.Put("/:id", controllers.UpdateNews)
	r.Delete("/", controllers.DeleteNews)

	return r
}
