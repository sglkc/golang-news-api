package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Heartbeat("/ping"))
	r.Mount("/debug", middleware.Profiler())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HELLO WORLD!"))
	})

	http.ListenAndServe(":3000", r)
}
