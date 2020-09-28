// Code generated by go-swagger; DO NOT EDIT.

package product_containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new product containers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for product containers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateReportAsync(params *CreateReportAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*CreateReportAsyncOK, error)

	DeleteProductContainer(params *DeleteProductContainerParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProductContainerOK, error)

	GetByIDWithRowIssues(params *GetByIDWithRowIssuesParams, authInfo runtime.ClientAuthInfoWriter) (*GetByIDWithRowIssuesOK, error)

	GetOrCreateShoppingCart(params *GetOrCreateShoppingCartParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrCreateShoppingCartOK, error)

	PatchProductRow(params *PatchProductRowParams, authInfo runtime.ClientAuthInfoWriter) (*PatchProductRowOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateReportAsync create report async API
*/
func (a *Client) CreateReportAsync(params *CreateReportAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*CreateReportAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateReportAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateReportAsync",
		Method:             "POST",
		PathPattern:        "/api/v1/ProductContainers/reportbymonth",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateReportAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateReportAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateReportAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteProductContainer delete product container API
*/
func (a *Client) DeleteProductContainer(params *DeleteProductContainerParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProductContainerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProductContainerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteProductContainer",
		Method:             "DELETE",
		PathPattern:        "/api/v1/ProductContainers/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteProductContainerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProductContainerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteProductContainer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetByIDWithRowIssues get by Id with row issues API
*/
func (a *Client) GetByIDWithRowIssues(params *GetByIDWithRowIssuesParams, authInfo runtime.ClientAuthInfoWriter) (*GetByIDWithRowIssuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetByIDWithRowIssuesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetByIdWithRowIssues",
		Method:             "GET",
		PathPattern:        "/api/v1/ProductContainers/rowissues/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetByIDWithRowIssuesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetByIDWithRowIssuesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetByIdWithRowIssues: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrCreateShoppingCart get or create shopping cart API
*/
func (a *Client) GetOrCreateShoppingCart(params *GetOrCreateShoppingCartParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrCreateShoppingCartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrCreateShoppingCartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOrCreateShoppingCart",
		Method:             "GET",
		PathPattern:        "/api/v1/ProductContainers/getorcreateshoppingcart",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOrCreateShoppingCartReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrCreateShoppingCartOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOrCreateShoppingCart: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchProductRow patch product row API
*/
func (a *Client) PatchProductRow(params *PatchProductRowParams, authInfo runtime.ClientAuthInfoWriter) (*PatchProductRowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchProductRowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchProductRow",
		Method:             "PATCH",
		PathPattern:        "/api/v1/ProductContainers/{productContainerId}/row/{productRowId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchProductRowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchProductRowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchProductRow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}