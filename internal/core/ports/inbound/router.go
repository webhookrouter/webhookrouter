package inbound

import "github.com/webhookrouter/webhookrouter/internal/core/domain"

type WebhookRouter interface {
	Route(webhook *domain.Webhook) error
}
