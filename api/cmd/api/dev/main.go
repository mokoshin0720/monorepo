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

	r, err := api.NewRouter()
	if err != nil {
		panic(err)
	}

	api, err := api.NewAPI(r)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	api.Run(ctx)
}
