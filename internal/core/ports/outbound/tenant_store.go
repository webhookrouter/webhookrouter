package outbound

import "github.com/webhookrouter/webhookrouter/internal/core/domain"

type TenantStore interface {
	// FindByID retrieves a tenant by its ID.
	FindByID(id string) (*domain.Tenant, error)
	// Save stores a tenant in the store.
	Save(tenant *domain.Tenant) error
	// Delete removes a tenant from the store.
	Delete(id string) error
}
