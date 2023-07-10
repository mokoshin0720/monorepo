package main

import (
	"context"

	"github.com/mokoshin0720/monorepo/api/server"
)

func main() {
	server, sclnup, err := server.New()
	if err != nil {
		panic(err)
	}
	defer sclnup()

	ctx := context.Background()
	server.Run(ctx)
}
