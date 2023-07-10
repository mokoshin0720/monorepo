package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/log"
)

func HandleRequest() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	log.Info().Msg("Starting server...")

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Info().Msg("routing server...")

	http.ListenAndServe(":9000", r)
}
