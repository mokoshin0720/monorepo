package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func HandleRequest() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	log.Info().Msg("Starting server...")

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Get("/memos", func(w http.ResponseWriter, r *http.Request) {
		memos := []Memo{
			{
				ID:    1,
				Title: "memo1",
			},
			{
				ID:    2,
				Title: "memo2",
			},
		}

		if err := json.NewEncoder(w).Encode(memos); err != nil {
			log.Error().Err(err).Msg("failed to encode memos")
		}
	})

	return r
}

type Memo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
