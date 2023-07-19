package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/mokoshin0720/monorepo/service/exercise"
)

type gqlController struct {
	exercise.Controller
}

const PORT = 9000

func newHandler() (*chi.Mux, error) {
	s, err := SchemaString()
	if err != nil {
		return nil, err
	}

	exercise, err := exercise.New()
	if err != nil {
		return nil, err
	}

	cont := &gqlController{*exercise}
	schema := graphql.MustParseSchema(s, cont)

	h := &relay.Handler{Schema: schema}
	r := chi.NewRouter()

	r.Mount("/", h)
	r.Group(func(r chi.Router) {
		r.Mount("/graphql", h)
	})

	r.Get("/schema.graphql", func(w http.ResponseWriter, r *http.Request) {
		Text(w, s)
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		Response(w, map[string]string{"messsage": "ok"})
	})

	return r, nil
}

// newServer http.Serverを生成する。内部で各サービスのコントローラを呼び出し、schemaとの整合性チェックする
func newServer(r *chi.Mux) (*http.Server, error) {
	port := fmt.Sprintf(":%d", PORT)

	srv := &http.Server{
		Addr:              port,
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
	}
	return srv, nil
}
