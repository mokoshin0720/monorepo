package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/mokoshin0720/monorepo/api/cmd/server"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting lambda server...")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "aws labs http adapter response!!")
	// })

	server.HandleRequest()

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
