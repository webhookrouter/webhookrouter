package endpoint

type Endpoint struct {
	ID           string        // Unique identifier for the endpoint
	TenantID     string        // ID of the tenant this endpoint belongs to
	Destinations []Destination // List of destinations associated with this endpoint
}
