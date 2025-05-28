package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/sglkc/golang-news-api/database"
	"github.com/sglkc/golang-news-api/routes"
)

func main() {
	err := database.Migrate()
	if err != nil {
		log.Fatalln("Cannot open database")
	}

	log.Println("Database connected and migrated")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/", routes.RootRoutes())
	r.Mount("/news", routes.NewsRoutes())

	http.ListenAndServe(":3000", r)
}
