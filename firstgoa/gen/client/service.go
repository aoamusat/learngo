// Code generated by goa v3.16.1, DO NOT EDIT.
//
// client service
//
// Command:
// $ goa gen learngo.io/firstgoa/design

package client

import (
	"context"

	clientviews "learngo.io/firstgoa/gen/client/views"
)

// The Client service allows access to client members
type Service interface {
	// Add implements add.
	Add(context.Context, *AddPayload) (err error)
	// Get implements get.
	Get(context.Context, *GetPayload) (res *ClientManagement, err error)
	// Show implements show.
	Show(context.Context) (res ClientManagementCollection, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "clients"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "client"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"add", "get", "show"}

// AddPayload is the payload type of the client service add method.
type AddPayload struct {
	// Client ID
	ClientID string
	// Client ID
	ClientName string
}

// ClientManagement is the result type of the client service get method.
type ClientManagement struct {
	// ID is the unique id of the Client.
	ClientID string
	// Name of the Client
	ClientName string
}

// ClientManagementCollection is the result type of the client service show
// method.
type ClientManagementCollection []*ClientManagement

// GetPayload is the payload type of the client service get method.
type GetPayload struct {
	// Client ID
	ClientID string
}

// NotFound is the type returned when the requested data that does not exist.
type NotFound struct {
	// Message of error
	Message string
	// ID of missing data
	ID string
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound is the type returned when the requested data that does not exist."
}

// ErrorName returns "NotFound".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *NotFound) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "NotFound".
func (e *NotFound) GoaErrorName() string {
	return "not_found"
}

// NewClientManagement initializes result type ClientManagement from viewed
// result type ClientManagement.
func NewClientManagement(vres *clientviews.ClientManagement) *ClientManagement {
	return newClientManagement(vres.Projected)
}

// NewViewedClientManagement initializes viewed result type ClientManagement
// from result type ClientManagement using the given view.
func NewViewedClientManagement(res *ClientManagement, view string) *clientviews.ClientManagement {
	p := newClientManagementView(res)
	return &clientviews.ClientManagement{Projected: p, View: "default"}
}

// NewClientManagementCollection initializes result type
// ClientManagementCollection from viewed result type
// ClientManagementCollection.
func NewClientManagementCollection(vres clientviews.ClientManagementCollection) ClientManagementCollection {
	return newClientManagementCollection(vres.Projected)
}

// NewViewedClientManagementCollection initializes viewed result type
// ClientManagementCollection from result type ClientManagementCollection using
// the given view.
func NewViewedClientManagementCollection(res ClientManagementCollection, view string) clientviews.ClientManagementCollection {
	p := newClientManagementCollectionView(res)
	return clientviews.ClientManagementCollection{Projected: p, View: "default"}
}

// newClientManagement converts projected type ClientManagement to service type
// ClientManagement.
func newClientManagement(vres *clientviews.ClientManagementView) *ClientManagement {
	res := &ClientManagement{}
	if vres.ClientID != nil {
		res.ClientID = *vres.ClientID
	}
	if vres.ClientName != nil {
		res.ClientName = *vres.ClientName
	}
	return res
}

// newClientManagementView projects result type ClientManagement to projected
// type ClientManagementView using the "default" view.
func newClientManagementView(res *ClientManagement) *clientviews.ClientManagementView {
	vres := &clientviews.ClientManagementView{
		ClientID:   &res.ClientID,
		ClientName: &res.ClientName,
	}
	return vres
}

// newClientManagementCollection converts projected type
// ClientManagementCollection to service type ClientManagementCollection.
func newClientManagementCollection(vres clientviews.ClientManagementCollectionView) ClientManagementCollection {
	res := make(ClientManagementCollection, len(vres))
	for i, n := range vres {
		res[i] = newClientManagement(n)
	}
	return res
}

// newClientManagementCollectionView projects result type
// ClientManagementCollection to projected type ClientManagementCollectionView
// using the "default" view.
func newClientManagementCollectionView(res ClientManagementCollection) clientviews.ClientManagementCollectionView {
	vres := make(clientviews.ClientManagementCollectionView, len(res))
	for i, n := range res {
		vres[i] = newClientManagementView(n)
	}
	return vres
}
