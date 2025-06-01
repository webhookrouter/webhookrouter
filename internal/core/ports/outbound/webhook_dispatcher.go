package outbound

import "github.com/webhookrouter/webhookrouter/internal/core/domain"

type WebhookDispatcher interface {
	Dispatch(event *domain.Webhook) error
}
