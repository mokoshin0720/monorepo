package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/mokoshin0720/monorepo/api/cmd/server"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting lambda server...")

	router := server.HandleRequest()

	lambda.Start(chiadapter.New(router).ProxyWithContext)

	// lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
