package services

import (
	"github.com/webhookrouter/webhookrouter/internal/core/domain"
	"github.com/webhookrouter/webhookrouter/internal/core/ports/outbound"
)

type Router struct {
	dispatcher    outbound.WebhookDispatcher
	endpointStore outbound.EndpointStore
}

func NewRouter(dispatcher outbound.WebhookDispatcher, endpointStore outbound.EndpointStore) *Router {
	if dispatcher == nil {
		panic("WebhookDispatcher cannot be nil")
	}
	if endpointStore == nil {
		panic("EndpointStore cannot be nil")
	}
	return &Router{
		dispatcher:    dispatcher,
		endpointStore: endpointStore,
	}
}

func (r *Router) Route(w *domain.Webhook) error {

	ep, err := r.endpointStore.FindByID(w.EndpointID)
	if err != nil {
		return err
	}
	if ep == nil {
		return domain.ErrEndpointNotFound
	}

	// Call the handler to process the webhook
	return r.dispatcher.Dispatch(w)
}
