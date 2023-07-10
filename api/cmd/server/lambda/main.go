package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/mokoshin0720/monorepo/api/cmd/server"
	"github.com/rs/zerolog/log"
)

var chiLambda *chiadapter.ChiLambda

func init() {
	log.Info().Msg("Initializing lambda server...")

	router := server.HandleRequest()

	chiLambda = chiadapter.New(router)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return chiLambda.ProxyWithContext(ctx, req)
}

func main() {
	log.Info().Msg("Starting lambda server...")
	lambda.Start(handler)
}
