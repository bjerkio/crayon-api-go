// Code generated by go-swagger; DO NOT EDIT.

package crayon_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new crayon accounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for crayon accounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCrayonAccounts(params *CreateCrayonAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCrayonAccountsOK, error)

	GetCrayonAccountByID(params *GetCrayonAccountByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCrayonAccountByIDOK, error)

	GetCrayonAccounts(params *GetCrayonAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCrayonAccountsOK, error)

	UpdateCrayonAccounts(params *UpdateCrayonAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCrayonAccountsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCrayonAccounts create crayon accounts API
*/
func (a *Client) CreateCrayonAccounts(params *CreateCrayonAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCrayonAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCrayonAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateCrayonAccounts",
		Method:             "POST",
		PathPattern:        "/api/v1/CrayonAccounts",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateCrayonAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCrayonAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateCrayonAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCrayonAccountByID get crayon account by Id API
*/
func (a *Client) GetCrayonAccountByID(params *GetCrayonAccountByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCrayonAccountByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCrayonAccountByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCrayonAccountById",
		Method:             "GET",
		PathPattern:        "/api/v1/CrayonAccounts/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCrayonAccountByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCrayonAccountByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCrayonAccountById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCrayonAccounts get crayon accounts API
*/
func (a *Client) GetCrayonAccounts(params *GetCrayonAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCrayonAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCrayonAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCrayonAccounts",
		Method:             "GET",
		PathPattern:        "/api/v1/CrayonAccounts",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCrayonAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCrayonAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCrayonAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCrayonAccounts update crayon accounts API
*/
func (a *Client) UpdateCrayonAccounts(params *UpdateCrayonAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCrayonAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCrayonAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateCrayonAccounts",
		Method:             "PUT",
		PathPattern:        "/api/v1/CrayonAccounts/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCrayonAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCrayonAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateCrayonAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
