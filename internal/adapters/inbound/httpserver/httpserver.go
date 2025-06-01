package httpserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/webhookrouter/webhookrouter/internal/core/ports/inbound"
)

type HttpServer struct {
	config        Config
	logger        zerolog.Logger
	webhookRouter inbound.WebhookRouter
	chi           *chi.Mux
}

type Config struct {
	Port int
}

func NewHttpServer(config Config, logger zerolog.Logger, router inbound.WebhookRouter) *HttpServer {
	return &HttpServer{
		config:        config,
		logger:        logger.With().Str("component", "httpserver").Logger(),
		webhookRouter: router,
	}
}

func (s *HttpServer) Start() error {
	s.logger.Debug().Msgf("Starting HTTP server at port %s", fmt.Sprintf(":%d", s.config.Port))

	s.chi = chi.NewRouter()
	s.chi.Use(loggerMiddleware(s.logger))

	s.chi.Post("/webhooks/{endpointId}", s.handleWebhook)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.chi)
}

func (s *HttpServer) Shutdown() error {
	s.logger.Debug().Msg("Shutting down HTTP server...")
	return nil
}
