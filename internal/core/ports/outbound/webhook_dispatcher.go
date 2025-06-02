package outbound

import (
	"github.com/webhookrouter/webhookrouter/internal/core/domain/endpoint"
	"github.com/webhookrouter/webhookrouter/internal/core/domain/webhook"
)

type WebhookDispatcher interface {
	Dispatch(event *webhook.Webhook, destination endpoint.Destination) error
}
