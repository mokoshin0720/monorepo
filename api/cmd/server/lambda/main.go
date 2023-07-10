package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mokoshin0720/monorepo/api/cmd/server"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting lambda server...")
	lambda.Start(server.HandleRequest)
}
