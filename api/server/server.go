package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	HTTPServer *http.Server
}

func New() (Server, func() error, error) {
	h, hclnup, err := NewHTTP()
	if err != nil {
		return Server{}, nil, err
	}

	clnup := func() error {
		hclnup()
		return nil
	}

	return Server{
		HTTPServer: h,
	}, clnup, nil
}

func (s Server) Run(ctx context.Context) {
	g, ctx := errgroup.WithContext(ctx)
	log.Info().Msg("Starting server")

	g.Go(func() error {
		err := s.HTTPServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cs := make(chan os.Signal, 1)
	signal.Notify(cs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	select {
	case <-ctx.Done():
		break
	case <-cs:
		break
	}

	s.HTTPServer.Shutdown(ctx)

	err := g.Wait()
	if err != nil {
		os.Exit(2)
	}
}
