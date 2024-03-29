// Code generated by go-swagger; DO NOT EDIT.

package publishers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new publishers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for publishers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetByID(params *GetByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetByIDOK, error)

	GetPublishers(params *GetPublishersParams, authInfo runtime.ClientAuthInfoWriter) (*GetPublishersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetByID get by Id API
*/
func (a *Client) GetByID(params *GetByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetById",
		Method:             "GET",
		PathPattern:        "/api/v1/Publishers/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPublishers get publishers API
*/
func (a *Client) GetPublishers(params *GetPublishersParams, authInfo runtime.ClientAuthInfoWriter) (*GetPublishersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublishersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPublishers",
		Method:             "GET",
		PathPattern:        "/api/v1/Publishers",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPublishersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPublishersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPublishers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
