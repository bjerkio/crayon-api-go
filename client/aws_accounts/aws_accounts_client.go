// Code generated by go-swagger; DO NOT EDIT.

package aws_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new aws accounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for aws accounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAwsAccountByID(params *GetAwsAccountByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetAwsAccountByIDOK, error)

	GetAwsAccounts(params *GetAwsAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAwsAccountsOK, error)

	UpdateAwsAccounts(params *UpdateAwsAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAwsAccountsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAwsAccountByID get aws account by Id API
*/
func (a *Client) GetAwsAccountByID(params *GetAwsAccountByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetAwsAccountByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAwsAccountByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAwsAccountById",
		Method:             "GET",
		PathPattern:        "/api/v1/AwsAccounts/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAwsAccountByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAwsAccountByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAwsAccountById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAwsAccounts get aws accounts API
*/
func (a *Client) GetAwsAccounts(params *GetAwsAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAwsAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAwsAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAwsAccounts",
		Method:             "GET",
		PathPattern:        "/api/v1/AwsAccounts",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAwsAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAwsAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAwsAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAwsAccounts update aws accounts API
*/
func (a *Client) UpdateAwsAccounts(params *UpdateAwsAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAwsAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAwsAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateAwsAccounts",
		Method:             "PUT",
		PathPattern:        "/api/v1/AwsAccounts/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAwsAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAwsAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateAwsAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
