package common

type Shutdowner interface {
	// Shutdown gracefully stops the service, cleaning up resources.
	Shutdown() error
	// Start initiates the service, preparing it to handle requests.
}
