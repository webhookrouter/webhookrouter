package httpserver

import (
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/webhookrouter/webhookrouter/internal/domain/webhook"
)

func (s *HttpServer) handleWebhook(w http.ResponseWriter, r *http.Request) {

	var payload []byte
	var headers map[string]string
	// Extract the payload and headers from the request
	if r.Body != nil {
		defer r.Body.Close()
		payload, _ = io.ReadAll(r.Body) // Read the request body
	}
	if r.Header != nil {
		headers = make(map[string]string)
		for key, values := range r.Header {
			// Use the first value for each header key
			if len(values) > 0 {
				headers[key] = values[0]
			}
		}
	}

	// Call the router to handle the webhook event
	if err := s.webhookService.Route(&webhook.Webhook{
		ID:         webhook.CreateId(),
		EndpointID: chi.URLParam(r, "endpointId"),
		ReceivedAt: time.Now(),
		Payload:    payload,
		Headers:    headers,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// TODO - consider using a more descriptive response body
	// to indicate that the webhook was processed successfully.
	w.Write([]byte("Webhook received"))
}
