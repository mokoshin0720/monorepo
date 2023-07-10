package main

import (
	"net/http"

	"github.com/mokoshin0720/monorepo/api/cmd/server"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting local server...")

	router := server.HandleRequest()

	http.ListenAndServe(":9000", router)
}
