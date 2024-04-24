package firstgoa

import (
	"context"
	"fmt"
	"log"

	"goa.design/goa/v3/security"
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

// JWTAuth implements the authorization logic for service "client" for the
// "jwt" security scheme.
func (s *clientsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
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
func (s *clientsrvc) Show(ctx context.Context, p *client.ShowPayload) (res client.ClientManagementCollection, err error) {
	s.logger.Print("client.show")
	return
}
