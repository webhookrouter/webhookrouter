package dummy

import (
	"github.com/rs/zerolog"
	"github.com/webhookrouter/webhookrouter/internal/domain/endpoint"
	"github.com/webhookrouter/webhookrouter/internal/domain/webhook"
)

type Dispatcher struct {
	logger zerolog.Logger
}

func NewDispatcher(logger zerolog.Logger) *Dispatcher {
	return &Dispatcher{
		logger: logger.With().Str("component", "dummy").Logger(),
	}
}

func (d *Dispatcher) Dispatch(webhook *webhook.Webhook, destination endpoint.Destination) error {
	d.logger.Info().
		Str("webhook.ID", webhook.ID).
		Any("webhook.Headers", webhook.Headers).
		Str("webhook.Payload", string(webhook.Payload)).
		Str("destination.URL", destination.URL).
		Msg("Dispatching webhook")
	return nil
}

func (d *Dispatcher) Shutdown() error {
	d.logger.Info().Msg("Shutdown")
	return nil
}
