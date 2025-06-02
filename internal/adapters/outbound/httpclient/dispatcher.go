package httpclient

import (
	"github.com/rs/zerolog"
	"github.com/webhookrouter/webhookrouter/internal/core/domain/endpoint"
	"github.com/webhookrouter/webhookrouter/internal/core/domain/webhook"
)

type Dispatcher struct {
	logger zerolog.Logger
}

func NewDispatcher(logger zerolog.Logger) *Dispatcher {
	return &Dispatcher{
		logger: logger.With().Str("component", "httpclient").Logger(),
	}
}
func (d *Dispatcher) Dispatch(webhook *webhook.Webhook, destination endpoint.Destination) error {
	d.logger.Info().Str("webhook", webhook.ID).Str("destination", destination.URL).Msg("Dispatching webhook")
	return nil
}
func (d *Dispatcher) Shutdown() error {
	d.logger.Info().Msg("Shutdown")
	return nil
}
