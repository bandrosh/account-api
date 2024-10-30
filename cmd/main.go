package main

import (
	"context"

	"github.com/hashicorp/go-multierror"

	"github.com/account-api/config"
	"github.com/account-api/infrastructure/logger"
	"github.com/account-api/internal/adapter/http/api"
)

func main() {
	cfg, err := config.LoadAppConfig("./scripts/config")
	if err != nil {
		logger.Fatal(logger.FatalError, "unable to load configurations")
	}

	var g multierror.Group

	ctx, stop := context.WithCancel(context.Background())

	a := api.New(ctx, cfg)
	g.Go(a.Run(ctx, stop))

	if err := g.Wait().ErrorOrNil(); err != nil {
		logger.Fatal(logger.ServerError, err)
	}
}
