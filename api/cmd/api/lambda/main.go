package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	api "github.com/mokoshin0720/monorepo"
	"github.com/mokoshin0720/monorepo/app/config"
)

var chiLambda *chiadapter.ChiLambda

func init() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	r, err := api.NewRouter()
	if err != nil {
		panic(err)
	}

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
