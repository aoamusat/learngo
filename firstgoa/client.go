package clients

import (
	"context"
	"log"

	client "learngo.io/firstgoa/gen/client"
)

// client service example implementation.
// The example methods log the requests and return zero values.
type clientsrvc struct {
	logger *log.Logger
}

// NewClient returns the client service implementation.
func NewClient(logger *log.Logger) client.Service {
	return &clientsrvc{logger}
}

// Add implements add.
func (s *clientsrvc) Add(ctx context.Context, p *client.AddPayload) (err error) {
	s.logger.Print("client.add")
	return
}

// Get implements get.
func (s *clientsrvc) Get(ctx context.Context, p *client.GetPayload) (res *client.ClientManagement, err error) {
	res = &client.ClientManagement{}
	s.logger.Print("client.get")
	return
}

// Show implements show.
func (s *clientsrvc) Show(ctx context.Context) (res client.ClientManagementCollection, err error) {
	s.logger.Print("client.show")
	return
}
