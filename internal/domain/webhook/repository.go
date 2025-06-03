package webhook

import "github.com/webhookrouter/webhookrouter/internal/domain/endpoint"

type WebhookDispatcher interface {
	Dispatch(event *Webhook, destination endpoint.Destination) error
}

type WebhookRepository interface {
	// FindByID retrieves a webhook by its ID.
	FindByID(id string) (*Webhook, error)
	// Save stores a webhook.
	Save(webhook *Webhook) error
}
