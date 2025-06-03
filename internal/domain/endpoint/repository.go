package endpoint

type EndpointRepository interface {
	// FindByID retrieves an endpoint by its ID.
	FindByID(id string) (*Endpoint, error)
	// Save stores an endpoint in the store.
	Save(endpoint *Endpoint) error
	// Delete removes an endpoint from the store.
	Delete(id string) error
}
