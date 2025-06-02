package webhook

import (
	"time"

	"github.com/google/uuid"
)

type Webhook struct {
	ID          string            // Unique identifier for the webhook
	EndpointID  string            // ID of the endpoint this webhook is associated with
	ReceivedAt  time.Time         // Timestamp when the webhook was received
	ProcessedAt time.Time         // Timestamp when the webhook was processed
	Payload     []byte            // The actual payload of the webhook, typically in JSON format
	Headers     map[string]string // Headers associated with the webhook, e.g., content type, user agent
}

func CreateId() string {
	return uuid.NewString()
}
