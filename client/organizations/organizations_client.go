// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new organizations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organizations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetOrganizationSalesContact(params *GetOrganizationSalesContactParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationSalesContactOK, error)

	GetOrganizations(params *GetOrganizationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationsOK, error)

	HasAccessAsync(params *HasAccessAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*HasAccessAsyncOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetOrganizationSalesContact get organization sales contact API
*/
func (a *Client) GetOrganizationSalesContact(params *GetOrganizationSalesContactParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationSalesContactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationSalesContactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOrganizationSalesContact",
		Method:             "GET",
		PathPattern:        "/api/v1/Organizations/{organizationId}/salescontact",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationSalesContactReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationSalesContactOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOrganizationSalesContact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizations get organizations API
*/
func (a *Client) GetOrganizations(params *GetOrganizationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOrganizations",
		Method:             "GET",
		PathPattern:        "/api/v1/Organizations",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOrganizations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  HasAccessAsync has access async API
*/
func (a *Client) HasAccessAsync(params *HasAccessAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*HasAccessAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHasAccessAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "HasAccessAsync",
		Method:             "GET",
		PathPattern:        "/api/v1/Organizations/HasAccess/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HasAccessAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HasAccessAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HasAccessAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
