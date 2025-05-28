package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sglkc/golang-news-api/controllers"
)

func ArticlesRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", controllers.ListArticles)
	r.Get("/{id}", controllers.GetArticle)
	r.Post("/", controllers.CreateArticle)
	r.Put("/{id}", controllers.UpdateArticle)
	r.Delete("/{id}", controllers.DeleteArticle)

	return r
}
