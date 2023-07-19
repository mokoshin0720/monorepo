package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

// API APIサーバーの構造体
type API struct {
	server *http.Server
}

func NewRouter() (*chi.Mux, error) {
	return newHandler()
}

// NewAPI APIサーバーのコンストラクタ
func NewAPI(r *chi.Mux) (API, error) {
	s, err := newServer(r)
	if err != nil {
		return API{}, err
	}

	return API{
		server: s,
	}, nil
}

// Run APIサーバーを起動する関数
func (a API) Run(ctx context.Context) {
	if err := a.server.ListenAndServe(); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(2)
	}
}
