package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/mokoshin0720/monorepo/api/cmd/server"
)

var chiLambda *chiadapter.ChiLambda

func init() {
	log.Printf("chi cold start")
	r := server.HandleRequest()

	chiLambda = chiadapter.New(r)
}

func Handler(ctx context.Context, req events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	apiReq := events.APIGatewayProxyRequest{
		Path:                  req.RawPath,
		HTTPMethod:            req.RequestContext.HTTP.Method,
		Headers:               req.Headers,
		QueryStringParameters: req.QueryStringParameters,
		Body:                  req.Body,
	}

	return chiLambda.ProxyWithContext(ctx, apiReq)
}

func main() {
	lambda.Start(Handler)
}
