// Code generated by go-swagger; DO NOT EDIT.

package google_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new google orders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for google orders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CheckoutAsync(params *CheckoutAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*CheckoutAsyncOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CheckoutAsync checkout async API
*/
func (a *Client) CheckoutAsync(params *CheckoutAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*CheckoutAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckoutAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CheckoutAsync",
		Method:             "POST",
		PathPattern:        "/api/v1/GoogleOrders/checkout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CheckoutAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CheckoutAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CheckoutAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
