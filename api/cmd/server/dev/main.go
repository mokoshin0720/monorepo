package main

import (
	"github.com/mokoshin0720/monorepo/api/cmd/server"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting local server...")
	
	server.HandleRequest()
}
