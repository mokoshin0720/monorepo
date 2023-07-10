package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/multierr"
)

const PORT = 9000

func NewHTTP() (*http.Server, func() error, error) {
	routers := []struct {
		new  func() (http.Handler, func() error, error)
		path string
	}{
		// {new: collabo.NewHandler, path: "/collaboration"},
		// {new: meeting.New, path: "/meeting"},
		// {new: admin.New, path: "/admin"},
	}

	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	clnups := make([]func() error, len(routers))
	for i, router := range routers {
		h, clnup, err := router.new()
		if err != nil {
			return nil, nil, err
		}
		r.Mount(router.path, h)
		clnups[i] = clnup
	}

	clnup := func() error {
		var errs error
		for _, c := range clnups {
			if c != nil {
				errs = multierr.Append(errs, c())
			}
		}
		return errs
	}

	port := fmt.Sprintf(":%d", PORT)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Debug().Msg("initializing server...")

	srv := &http.Server{Addr: port, Handler: r}
	return srv, clnup, nil
}
