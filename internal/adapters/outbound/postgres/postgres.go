package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/webhookrouter/webhookrouter/internal/domain/endpoint"
)

type PostgresStore struct {
	logger     zerolog.Logger
	connection pgx.Conn
}

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  string
}

func NewPostgresStore(cfg Config, logger zerolog.Logger) (*PostgresStore, error) {
	// Here you would typically initialize the PostgreSQL connection
	// using the provided configuration and return a Postgres instance.
	// For now, we will just return a new Postgres instance with the logger.

	return &PostgresStore{
		logger: logger.With().Str("component", "postgres").Logger(),
	}, nil
}

func (p *PostgresStore) FindByID(id string) (*endpoint.Endpoint, error) {
	p.logger.Info().Str("id", id).Msg("Finding endpoint by ID")
	return nil, nil
}

func (p *PostgresStore) Delete(id string) error {
	p.logger.Info().Str("id", id).Msg("Delete endpoint by ID")
	// TODO: implement the actual deletion logic
	return nil
}

func (p *PostgresStore) Save(endpoint *endpoint.Endpoint) error {
	p.logger.Info().Str("id", endpoint.ID).Msg("Saving endpoint")
	return nil
}

func (p *PostgresStore) Close() error {
	p.logger.Debug().Msg("Closing PostgreSQL connection")
	p.connection.Close(context.Background())
	p.logger.Debug().Msg("Closed PostgreSQL connection")
	return nil
}
