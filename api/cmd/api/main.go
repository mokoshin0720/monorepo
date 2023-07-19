package main

import (
	"context"

	api "github.com/mokoshin0720/monorepo"
	"github.com/mokoshin0720/monorepo/app/config"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	api, err := api.NewAPI()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	api.Run(ctx)
}
