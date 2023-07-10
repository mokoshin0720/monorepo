package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/log"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	log.Info().Msg("Starting server...")

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Info().Msg("routing server...")

	http.ListenAndServe(":9000", r)
}
