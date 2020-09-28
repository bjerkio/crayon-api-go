// Code generated by go-swagger; DO NOT EDIT.

package organization_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new organization access API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organization access API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetOrganizationAccess(params *GetOrganizationAccessParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationAccessOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetOrganizationAccess get organization access API
*/
func (a *Client) GetOrganizationAccess(params *GetOrganizationAccessParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationAccessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationAccessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOrganizationAccess",
		Method:             "GET",
		PathPattern:        "/api/v1/OrganizationAccess",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOrganizationAccessReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationAccessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOrganizationAccess: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
