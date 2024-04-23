// Code generated by goa v3.16.1, DO NOT EDIT.
//
// client HTTP client CLI support package
//
// Command:
// $ goa gen learngo.io/firstgoa/design

package client

import (
	"encoding/json"
	"fmt"

	client "learngo.io/firstgoa/gen/client"
)

// BuildAddPayload builds the payload for the client add endpoint from CLI
// flags.
func BuildAddPayload(clientAddBody string, clientAddClientID string) (*client.AddPayload, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(clientAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"ClientName\": \"Pariatur consequatur reiciendis rerum similique aut dolore.\"\n   }'")
		}
	}
	var clientID string
	{
		clientID = clientAddClientID
	}
	v := &client.AddPayload{
		ClientName: body.ClientName,
	}
	v.ClientID = clientID

	return v, nil
}

// BuildGetPayload builds the payload for the client get endpoint from CLI
// flags.
func BuildGetPayload(clientGetClientID string) (*client.GetPayload, error) {
	var clientID string
	{
		clientID = clientGetClientID
	}
	v := &client.GetPayload{}
	v.ClientID = clientID

	return v, nil
}
