package config

import (
	"github.com/webhookrouter/webhookrouter/internal/adapters/inbound/httpserver"
	"github.com/webhookrouter/webhookrouter/internal/adapters/outbound/postgres"
)

type Config struct {
	HttpServer httpserver.Config // Embedding the HTTP server configuration
	Postgres   postgres.Config   // PostgreSQL configuration
}

func Load() Config {
	// In a real application, this function would load configuration from files, environment variables, etc.
	// For simplicity, we return an empty Config struct here.
	return Config{}
}
