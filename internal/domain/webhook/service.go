package webhook

import (
	"github.com/rs/zerolog"
	"github.com/webhookrouter/webhookrouter/internal/domain"
	"github.com/webhookrouter/webhookrouter/internal/domain/endpoint"
)

type WebhookService interface {
	Route(webhook *Webhook) error
}

type webhookService struct {
	dispatcher   WebhookDispatcher
	endpointRepo endpoint.EndpointRepository
	logger       zerolog.Logger
}

func NewService(dispatcher WebhookDispatcher, endpointRepo endpoint.EndpointRepository, logger zerolog.Logger) *webhookService {
	if dispatcher == nil {
		panic("WebhookDispatcher cannot be nil")
	}
	if endpointRepo == nil {
		panic("EndpointStore cannot be nil")
	}
	return &webhookService{
		dispatcher:   dispatcher,
		endpointRepo: endpointRepo,
		logger:       logger.With().Str("component", "router").Logger(),
	}
}

func (r *webhookService) Route(w *Webhook) error {

	ep, err := r.endpointRepo.FindByID(w.EndpointID)
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
