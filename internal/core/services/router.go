package services

import (
	"github.com/rs/zerolog"
	"github.com/webhookrouter/webhookrouter/internal/core/domain"
	"github.com/webhookrouter/webhookrouter/internal/core/domain/webhook"
	"github.com/webhookrouter/webhookrouter/internal/core/ports/outbound"
)

type Router struct {
	dispatcher    outbound.WebhookDispatcher
	endpointStore outbound.EndpointStore
	logger        zerolog.Logger
}

func NewRouter(dispatcher outbound.WebhookDispatcher, endpointStore outbound.EndpointStore, logger zerolog.Logger) *Router {
	if dispatcher == nil {
		panic("WebhookDispatcher cannot be nil")
	}
	if endpointStore == nil {
		panic("EndpointStore cannot be nil")
	}
	return &Router{
		dispatcher:    dispatcher,
		endpointStore: endpointStore,
		logger:        logger.With().Str("component", "router").Logger(),
	}
}

func (r *Router) Route(w *webhook.Webhook) error {

	ep, err := r.endpointStore.FindByID(w.EndpointID)
	if err != nil {
		return err
	}
	if ep == nil {
		return domain.ErrEndpointNotFound
	}

	for _, dest := range ep.Destinations {
		if dest.Enabled == false {
			// Skip disabled destinations
			r.logger.Debug().Str("destination", dest.URL).Msg("Skipping disabled destination")
			continue
		}
		// Call the handler to process the webhook
		err := r.dispatcher.Dispatch(w, dest)
		if err != nil {
			// Log the error and continue processing other destinations
			r.logger.Error().Err(err)
		}

	}
	return nil

}
