// Code generated by go-swagger; DO NOT EDIT.

package management_links

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new management links API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for management links API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetGrouped(params *GetGroupedParams, authInfo runtime.ClientAuthInfoWriter) (*GetGroupedOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetGrouped get grouped API
*/
func (a *Client) GetGrouped(params *GetGroupedParams, authInfo runtime.ClientAuthInfoWriter) (*GetGroupedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetGrouped",
		Method:             "GET",
		PathPattern:        "/api/v1/ManagementLinks/grouped",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGroupedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGroupedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetGrouped: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}