package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mokoshin0720/monorepo/api/cmd/server"
)

func main() {
	lambda.Start(server.HandleRequest)
}
