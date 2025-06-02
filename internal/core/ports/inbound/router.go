package inbound

import (
	"github.com/webhookrouter/webhookrouter/internal/core/domain/webhook"
)

type WebhookRouter interface {
	Route(webhook *webhook.Webhook) error
}
