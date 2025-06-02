package outbound

import (
	"github.com/webhookrouter/webhookrouter/internal/core/domain/endpoint"
)

type EndpointStore interface {
	// FindByID retrieves an endpoint by its ID.
	FindByID(id string) (*endpoint.Endpoint, error)
	// Save stores an endpoint in the store.
	Save(endpoint *endpoint.Endpoint) error
	// Delete removes an endpoint from the store.
	Delete(id string) error
}
