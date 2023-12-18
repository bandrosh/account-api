package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/account-api/config"
	"github.com/account-api/infrastructure/logger"
	"github.com/account-api/infrastructure/postgres"
	"github.com/account-api/internal/adapter/http/handlers"
	"github.com/account-api/internal/adapter/http/middlewares"
	"github.com/account-api/internal/adapter/repository"
	"github.com/account-api/internal/core/ports"
	"github.com/account-api/internal/core/services"
)

type api struct {
	config   config.Configuration
	services svs
}

type svs struct {
	account ports.AccountService
}

func New(ctx context.Context, cfg config.Configuration) (a api) {
	a.config = cfg

	pgql, err := postgres.ConnectPostgresDB(ctx, a.config.Postgres.Url)
	if err != nil {
		logger.Fatal(logger.ConfigError, fmt.Sprintf("Cannot connect postgresql error: %v", err))
	}

	accountRepository := repository.NewAccountRepository(pgql)
	a.services.account = services.NewAccountService(a.config, accountRepository)
	return a
}

func (a *api) Run(ctx context.Context, cancel context.CancelFunc) func() error {
	return func() error {
		defer cancel()

		router := gin.Default()

		router.Use(middlewares.Recover())

		handlers.SetAccountRoutes(ctx, a.config, router, a.services.account)

		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", a.config.Server.Port),
			Handler: router,
		}

		go shutdown(ctx, server)
		return server.ListenAndServe()
	}
}

func shutdown(ctx context.Context, server *http.Server) {
	<-ctx.Done()
	logger.Info(logger.ServerInfo, "New we can do an shutdown gracefully")
	server.Shutdown(ctx)
}
