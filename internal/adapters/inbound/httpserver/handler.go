package httpserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/webhookrouter/webhookrouter/internal/core/domain/webhook"
)

func (s *HttpServer) handleWebhook(w http.ResponseWriter, r *http.Request) {

	// Call the router to handle the webhook event
	if err := s.webhookRouter.Route(&webhook.Webhook{
		EndpointID: chi.URLParam(r, "endpointId"),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// TODO - consider using a more descriptive response body
	// to indicate that the webhook was processed successfully.
	w.Write([]byte("Webhook received"))
}
