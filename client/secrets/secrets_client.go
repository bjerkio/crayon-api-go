// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new secrets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for secrets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSecrets(params *CreateSecretsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSecretsOK, error)

	DeleteSecrets(params *DeleteSecretsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSecretsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateSecrets create secrets API
*/
func (a *Client) CreateSecrets(params *CreateSecretsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSecretsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSecretsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateSecrets",
		Method:             "POST",
		PathPattern:        "/api/v1/Secrets",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSecretsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSecretsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateSecrets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteSecrets delete secrets API
*/
func (a *Client) DeleteSecrets(params *DeleteSecretsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSecretsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSecretsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteSecrets",
		Method:             "DELETE",
		PathPattern:        "/api/v1/Secrets",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSecretsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSecretsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteSecrets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
