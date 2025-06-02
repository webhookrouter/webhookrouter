package endpoint

type Destination struct {
	ID      string            // Unique identifier for the destination
	Name    string            // Name of the destination
	URL     string            // URL of the destination where webhooks will be sent
	Method  string            // HTTP method to use for the webhook request (e.g., GET, POST)
	Headers map[string]string // Headers to be included in the webhook request
	Timeout int               // Timeout in seconds for the webhook request
	Enabled bool              // Indicates if the destination is enabled or not
}
