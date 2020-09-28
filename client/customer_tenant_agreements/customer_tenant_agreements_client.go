// Code generated by go-swagger; DO NOT EDIT.

package customer_tenant_agreements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customer tenant agreements API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer tenant agreements API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	Add(params *AddParams, authInfo runtime.ClientAuthInfoWriter) (*AddOK, error)

	GetCustomerTenantAgreements(params *GetCustomerTenantAgreementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerTenantAgreementsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Add add API
*/
func (a *Client) Add(params *AddParams, authInfo runtime.ClientAuthInfoWriter) (*AddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Add",
		Method:             "POST",
		PathPattern:        "/api/v1/customertenants/{customerTenantId}/agreements",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Add: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCustomerTenantAgreements get customer tenant agreements API
*/
func (a *Client) GetCustomerTenantAgreements(params *GetCustomerTenantAgreementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerTenantAgreementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerTenantAgreementsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomerTenantAgreements",
		Method:             "GET",
		PathPattern:        "/api/v1/customertenants/{customerTenantId}/agreements",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomerTenantAgreementsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomerTenantAgreementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCustomerTenantAgreements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
