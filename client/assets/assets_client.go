// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new assets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for assets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddTagsAsync(params *AddTagsAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*AddTagsAsyncOK, error)

	DeleteTagsAsync(params *DeleteTagsAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTagsAsyncOK, error)

	GetAssetOrdersAsync(params *GetAssetOrdersAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssetOrdersAsyncOK, error)

	UpdateTagsAsync(params *UpdateTagsAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTagsAsyncOK, error)

	VerifyAsync(params *VerifyAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*VerifyAsyncOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddTagsAsync add tags async API
*/
func (a *Client) AddTagsAsync(params *AddTagsAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*AddTagsAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddTagsAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddTagsAsync",
		Method:             "POST",
		PathPattern:        "/api/v1/Assets/{assetId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddTagsAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddTagsAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AddTagsAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTagsAsync delete tags async API
*/
func (a *Client) DeleteTagsAsync(params *DeleteTagsAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTagsAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTagsAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteTagsAsync",
		Method:             "DELETE",
		PathPattern:        "/api/v1/Assets/{assetId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTagsAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTagsAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteTagsAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAssetOrdersAsync get asset orders async API
*/
func (a *Client) GetAssetOrdersAsync(params *GetAssetOrdersAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssetOrdersAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssetOrdersAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAssetOrdersAsync",
		Method:             "GET",
		PathPattern:        "/api/v1/Assets/orders",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAssetOrdersAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssetOrdersAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAssetOrdersAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateTagsAsync update tags async API
*/
func (a *Client) UpdateTagsAsync(params *UpdateTagsAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTagsAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTagsAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateTagsAsync",
		Method:             "PUT",
		PathPattern:        "/api/v1/Assets/{assetId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTagsAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateTagsAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateTagsAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VerifyAsync verify async API
*/
func (a *Client) VerifyAsync(params *VerifyAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*VerifyAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VerifyAsync",
		Method:             "POST",
		PathPattern:        "/api/v1/Assets/verify",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VerifyAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VerifyAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VerifyAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
