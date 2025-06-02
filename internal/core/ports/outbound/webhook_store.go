package outbound

import (
	"github.com/webhookrouter/webhookrouter/internal/core/domain/webhook"
)

type WebhookStore interface {
	// FindByID retrieves a webhook by its ID.
	FindByID(id string) (*webhook.Webhook, error)
	// Save stores a webhook in the store.
	Save(webhook *webhook.Webhook) error
}
