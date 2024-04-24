// Code generated by goa v3.16.1, DO NOT EDIT.
//
// client HTTP server types
//
// Command:
// $ goa gen learngo.io/firstgoa/design

package server

import (
	goa "goa.design/goa/v3/pkg"
	client "learngo.io/firstgoa/gen/client"
	clientviews "learngo.io/firstgoa/gen/client/views"
)

// AddRequestBody is the type of the "client" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Client ID
	ClientName *string `form:"ClientName,omitempty" json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// Contact Name
	ContactName *string `form:"ContactName,omitempty" json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// Contact Email
	ContactEmail *string `form:"ContactEmail,omitempty" json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// Contact Mobile Number
	ContactMobile *int `form:"ContactMobile,omitempty" json:"ContactMobile,omitempty" xml:"ContactMobile,omitempty"`
}

// GetResponseBody is the type of the "client" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// ID is the unique id of the Client.
	ClientID string `form:"ClientID" json:"ClientID" xml:"ClientID"`
	// Name of the Client
	ClientName string `form:"ClientName" json:"ClientName" xml:"ClientName"`
	// Name of the Contact.
	ContactName string `form:"ContactName" json:"ContactName" xml:"ContactName"`
	// Email of the Client Contact
	ContactEmail string `form:"ContactEmail" json:"ContactEmail" xml:"ContactEmail"`
	// Mobile number of the Client Contact
	ContactMobile int `form:"ContactMobile" json:"ContactMobile" xml:"ContactMobile"`
}

// ClientManagementResponseCollection is the type of the "client" service
// "show" endpoint HTTP response body.
type ClientManagementResponseCollection []*ClientManagementResponse

// ClientManagementResponse is used to define fields on response body types.
type ClientManagementResponse struct {
	// ID is the unique id of the Client.
	ClientID string `form:"ClientID" json:"ClientID" xml:"ClientID"`
	// Name of the Client
	ClientName string `form:"ClientName" json:"ClientName" xml:"ClientName"`
	// Name of the Contact.
	ContactName string `form:"ContactName" json:"ContactName" xml:"ContactName"`
	// Email of the Client Contact
	ContactEmail string `form:"ContactEmail" json:"ContactEmail" xml:"ContactEmail"`
	// Mobile number of the Client Contact
	ContactMobile int `form:"ContactMobile" json:"ContactMobile" xml:"ContactMobile"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "client" service.
func NewGetResponseBody(res *clientviews.ClientManagementView) *GetResponseBody {
	body := &GetResponseBody{
		ClientID:      *res.ClientID,
		ClientName:    *res.ClientName,
		ContactName:   *res.ContactName,
		ContactEmail:  *res.ContactEmail,
		ContactMobile: *res.ContactMobile,
	}
	return body
}

// NewClientManagementResponseCollection builds the HTTP response body from the
// result of the "show" endpoint of the "client" service.
func NewClientManagementResponseCollection(res clientviews.ClientManagementCollectionView) ClientManagementResponseCollection {
	body := make([]*ClientManagementResponse, len(res))
	for i, val := range res {
		body[i] = marshalClientviewsClientManagementViewToClientManagementResponse(val)
	}
	return body
}

// NewAddPayload builds a client service add endpoint payload.
func NewAddPayload(body *AddRequestBody, clientID string, token string) *client.AddPayload {
	v := &client.AddPayload{
		ClientName:    *body.ClientName,
		ContactName:   *body.ContactName,
		ContactEmail:  *body.ContactEmail,
		ContactMobile: *body.ContactMobile,
	}
	v.ClientID = clientID
	v.Token = token

	return v
}

// NewGetPayload builds a client service get endpoint payload.
func NewGetPayload(clientID string, token string) *client.GetPayload {
	v := &client.GetPayload{}
	v.ClientID = clientID
	v.Token = token

	return v
}

// NewShowPayload builds a client service show endpoint payload.
func NewShowPayload(token string) *client.ShowPayload {
	v := &client.ShowPayload{}
	v.Token = token

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.ClientName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ClientName", "body"))
	}
	if body.ContactName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ContactName", "body"))
	}
	if body.ContactEmail == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ContactEmail", "body"))
	}
	if body.ContactMobile == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ContactMobile", "body"))
	}
	return
}
