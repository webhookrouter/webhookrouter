package outbound

import "github.com/webhookrouter/webhookrouter/internal/core/domain"

type WebhookStore interface {
	// FindByID retrieves a webhook by its ID.
	FindByID(id string) (*domain.Webhook, error)
	// Save stores a webhook in the store.
	Save(webhook *domain.Webhook) error
}
