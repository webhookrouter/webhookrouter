package app

import (
	"log"

	"github.com/rs/zerolog"
	"github.com/webhookrouter/webhookrouter/internal/adapters/inbound/httpserver"
	"github.com/webhookrouter/webhookrouter/internal/common"
)

type Application struct {
	logger     zerolog.Logger
	httpServer *httpserver.HttpServer
	shutdowner []common.Shutdowner
}

func NewApplication(logger zerolog.Logger, httpServer *httpserver.HttpServer, shutdowner []common.Shutdowner) *Application {
	return &Application{
		httpServer: httpServer,
		logger:     logger,
		shutdowner: shutdowner,
	}
}

func (a *Application) Shutdown() {
	a.logger.Debug().Msg("Cleaning up resources...")
	for _, sd := range a.shutdowner {
		if err := sd.Shutdown(); err != nil {
			log.Printf("shutdown error: %v", err)
		}
	}
}

func (a *Application) StartHTTP() error {

	// Create a new HTTP server
	return a.httpServer.Start()
}
