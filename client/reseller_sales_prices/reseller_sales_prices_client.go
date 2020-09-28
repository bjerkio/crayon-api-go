// Code generated by go-swagger; DO NOT EDIT.

package reseller_sales_prices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new reseller sales prices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reseller sales prices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteByFilterAsync(params *DeleteByFilterAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteByFilterAsyncOK, error)

	GetAsync(params *GetAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*GetAsyncOK, error)

	GetCurrentAsync(params *GetCurrentAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrentAsyncOK, error)

	PostAsync(params *PostAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*PostAsyncOK, error)

	PutAsync(params *PutAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*PutAsyncOK, error)

	ToggleAsync(params *ToggleAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*ToggleAsyncOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteByFilterAsync delete by filter async API
*/
func (a *Client) DeleteByFilterAsync(params *DeleteByFilterAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteByFilterAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteByFilterAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteByFilterAsync",
		Method:             "DELETE",
		PathPattern:        "/api/v1/ResellerSalesPrices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteByFilterAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteByFilterAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteByFilterAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAsync get async API
*/
func (a *Client) GetAsync(params *GetAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*GetAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAsync",
		Method:             "GET",
		PathPattern:        "/api/v1/ResellerSalesPrices",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCurrentAsync get current async API
*/
func (a *Client) GetCurrentAsync(params *GetCurrentAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrentAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCurrentAsync",
		Method:             "GET",
		PathPattern:        "/api/v1/ResellerSalesPrices/current",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCurrentAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCurrentAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCurrentAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAsync post async API
*/
func (a *Client) PostAsync(params *PostAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*PostAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAsync",
		Method:             "POST",
		PathPattern:        "/api/v1/ResellerSalesPrices",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutAsync put async API
*/
func (a *Client) PutAsync(params *PutAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*PutAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutAsync",
		Method:             "PUT",
		PathPattern:        "/api/v1/ResellerSalesPrices/{oldFromDate}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ToggleAsync toggle async API
*/
func (a *Client) ToggleAsync(params *ToggleAsyncParams, authInfo runtime.ClientAuthInfoWriter) (*ToggleAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewToggleAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ToggleAsync",
		Method:             "POST",
		PathPattern:        "/api/v1/ResellerSalesPrices/toggle",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ToggleAsyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ToggleAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ToggleAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
