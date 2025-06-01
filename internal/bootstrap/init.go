package bootstrap

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/webhookrouter/webhookrouter/internal/adapters/inbound/httpserver"
	"github.com/webhookrouter/webhookrouter/internal/adapters/outbound/dummy"
	"github.com/webhookrouter/webhookrouter/internal/adapters/outbound/httpclient"
	"github.com/webhookrouter/webhookrouter/internal/adapters/outbound/inmemory"
	"github.com/webhookrouter/webhookrouter/internal/adapters/outbound/postgres"
	"github.com/webhookrouter/webhookrouter/internal/app"
	"github.com/webhookrouter/webhookrouter/internal/common"
	"github.com/webhookrouter/webhookrouter/internal/config"
	"github.com/webhookrouter/webhookrouter/internal/core/domain"
	"github.com/webhookrouter/webhookrouter/internal/core/services"
)

func InitApplication(ctx context.Context, cfg config.Config, logger zerolog.Logger) *app.Application {
	logger = logger.With().Str("component", "application").Logger()

	postgres, err := postgres.NewPostgresStore(cfg.Postgres, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize PostgreSQL store")
	}

	dispatcher := httpclient.NewDispatcher(logger)
	router := services.NewRouter(
		dispatcher, postgres,
	)
	httpserver := httpserver.NewHttpServer(cfg.HttpServer, logger, router)
	// Initialize the application
	application := app.NewApplication(logger,
		httpserver,
		[]common.Shutdowner{
			httpserver,
		},
	)
	return application
}

func InitTestApplication(ctx context.Context, cfg config.Config, logger zerolog.Logger) *app.Application {
	cfg.HttpServer.Port = 8080

	logger = logger.With().Str("component", "application").Logger()

	inmemory, err := inmemory.NewInMemoryStore(inmemory.Config{}, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize InMemory store")
	}

	inmemory.Save(&domain.Endpoint{
		ID:       "test-endpoint",
		TenantID: "test-tenant"})

	dispatcher := dummy.NewDispatcher(logger)
	router := services.NewRouter(
		dispatcher, inmemory,
	)
	httpserver := httpserver.NewHttpServer(cfg.HttpServer, logger, router)
	// Initialize the application for testing
	application := app.NewApplication(logger,
		httpserver,
		[]common.Shutdowner{
			httpserver,
		},
	)
	return application
}
