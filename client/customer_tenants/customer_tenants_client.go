// Code generated by go-swagger; DO NOT EDIT.

package customer_tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customer tenants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer tenants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddExisting(params *AddExistingParams, authInfo runtime.ClientAuthInfoWriter) (*AddExistingOK, error)

	DeleteCustomerByID(params *DeleteCustomerByIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCustomerByIDOK, error)

	GetAzurePlan(params *GetAzurePlanParams, authInfo runtime.ClientAuthInfoWriter) (*GetAzurePlanOK, error)

	GetCustomerTenantByID(params *GetCustomerTenantByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerTenantByIDOK, error)

	GetCustomerTenantDetailedByID(params *GetCustomerTenantDetailedByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerTenantDetailedByIDOK, error)

	GetCustomerTenants(params *GetCustomerTenantsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerTenantsOK, error)

	PostCustomerTenants(params *PostCustomerTenantsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomerTenantsOK, error)

	PutCustomerTenants(params *PutCustomerTenantsParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomerTenantsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddExisting add existing API
*/
func (a *Client) AddExisting(params *AddExistingParams, authInfo runtime.ClientAuthInfoWriter) (*AddExistingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddExistingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddExisting",
		Method:             "POST",
		PathPattern:        "/api/v1/CustomerTenants/existing",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddExistingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddExistingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AddExisting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCustomerByID delete customer by Id API
*/
func (a *Client) DeleteCustomerByID(params *DeleteCustomerByIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCustomerByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCustomerByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteCustomerById",
		Method:             "DELETE",
		PathPattern:        "/api/v1/CustomerTenants/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCustomerByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCustomerByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteCustomerById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAzurePlan get azure plan API
*/
func (a *Client) GetAzurePlan(params *GetAzurePlanParams, authInfo runtime.ClientAuthInfoWriter) (*GetAzurePlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAzurePlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAzurePlan",
		Method:             "GET",
		PathPattern:        "/api/v1/CustomerTenants/{customerTenantId}/azurePlan",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAzurePlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAzurePlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAzurePlan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCustomerTenantByID get customer tenant by Id API
*/
func (a *Client) GetCustomerTenantByID(params *GetCustomerTenantByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerTenantByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerTenantByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomerTenantById",
		Method:             "GET",
		PathPattern:        "/api/v1/CustomerTenants/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCustomerTenantByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomerTenantByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCustomerTenantById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCustomerTenantDetailedByID get customer tenant detailed by Id API
*/
func (a *Client) GetCustomerTenantDetailedByID(params *GetCustomerTenantDetailedByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerTenantDetailedByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerTenantDetailedByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomerTenantDetailedById",
		Method:             "GET",
		PathPattern:        "/api/v1/CustomerTenants/{id}/detailed",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCustomerTenantDetailedByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomerTenantDetailedByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCustomerTenantDetailedById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCustomerTenants get customer tenants API
*/
func (a *Client) GetCustomerTenants(params *GetCustomerTenantsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerTenantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerTenantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomerTenants",
		Method:             "GET",
		PathPattern:        "/api/v1/CustomerTenants",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCustomerTenantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomerTenantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCustomerTenants: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCustomerTenants post customer tenants API
*/
func (a *Client) PostCustomerTenants(params *PostCustomerTenantsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomerTenantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomerTenantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomerTenants",
		Method:             "POST",
		PathPattern:        "/api/v1/CustomerTenants",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostCustomerTenantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCustomerTenantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostCustomerTenants: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutCustomerTenants put customer tenants API
*/
func (a *Client) PutCustomerTenants(params *PutCustomerTenantsParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomerTenantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCustomerTenantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCustomerTenants",
		Method:             "PUT",
		PathPattern:        "/api/v1/CustomerTenants/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutCustomerTenantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCustomerTenantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutCustomerTenants: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
