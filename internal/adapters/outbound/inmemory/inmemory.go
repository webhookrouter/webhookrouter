package inmemory

import (
	"github.com/rs/zerolog"
	"github.com/webhookrouter/webhookrouter/internal/core/domain"
)

type InMemoryStore struct {
	logger    zerolog.Logger
	endpoints map[string]*domain.Endpoint
}
type Config struct{}

func NewInMemoryStore(cfg Config, logger zerolog.Logger) (*InMemoryStore, error) {
	// Initialize the in-memory store with an empty map for endpoints
	return &InMemoryStore{
		logger:    logger.With().Str("component", "inmemory").Logger(),
		endpoints: make(map[string]*domain.Endpoint),
	}, nil
}
func (i *InMemoryStore) FindByID(id string) (*domain.Endpoint, error) {
	i.logger.Debug().Str("id", id).Msg("Finding endpoint by ID")
	endpoint, exists := i.endpoints[id]
	if !exists {
		return nil, nil // or return an error if preferred
	}
	return endpoint, nil
}
func (i *InMemoryStore) Delete(id string) error {
	i.logger.Debug().Str("id", id).Msg("Delete endpoint by ID")
	if _, exists := i.endpoints[id]; !exists {
		return nil // or return an error if preferred
	}
	delete(i.endpoints, id)
	return nil
}
func (i *InMemoryStore) Save(endpoint *domain.Endpoint) error {
	i.logger.Debug().Str("id", endpoint.ID).Msg("Saving endpoint")
	if endpoint == nil {
		return nil // or return an error if preferred
	}
	i.endpoints[endpoint.ID] = endpoint
	return nil
}
func (i *InMemoryStore) Close() error {
	i.logger.Debug().Msg("Closing InMemory store")
	// No resources to close in an in-memory store, but we can log it
	i.endpoints = nil // Clear the map if needed
	i.logger.Debug().Msg("Closed InMemory store")
	return nil
}
