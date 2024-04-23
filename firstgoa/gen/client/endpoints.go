// Code generated by goa v3.16.1, DO NOT EDIT.
//
// client endpoints
//
// Command:
// $ goa gen learngo.io/firstgoa/design

package client

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "client" service endpoints.
type Endpoints struct {
	Add  goa.Endpoint
	Get  goa.Endpoint
	Show goa.Endpoint
}

// NewEndpoints wraps the methods of the "client" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Add:  NewAddEndpoint(s),
		Get:  NewGetEndpoint(s),
		Show: NewShowEndpoint(s),
	}
}

// Use applies the given middleware to all the "client" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Add = m(e.Add)
	e.Get = m(e.Get)
	e.Show = m(e.Show)
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "client".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*AddPayload)
		return nil, s.Add(ctx, p)
	}
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "client".
func NewGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetPayload)
		res, err := s.Get(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedClientManagement(res, "default")
		return vres, nil
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "client".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		res, err := s.Show(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedClientManagementCollection(res, "default")
		return vres, nil
	}
}