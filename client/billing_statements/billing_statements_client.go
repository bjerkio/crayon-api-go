// Code generated by go-swagger; DO NOT EDIT.

package billing_statements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new billing statements API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for billing statements API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetBillingRecordsFile(params *GetBillingRecordsFileParams, authInfo runtime.ClientAuthInfoWriter) (*GetBillingRecordsFileOK, error)

	GetBillingStatementFile(params *GetBillingStatementFileParams, authInfo runtime.ClientAuthInfoWriter) (*GetBillingStatementFileOK, error)

	GetBillingStatements(params *GetBillingStatementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetBillingStatementsOK, error)

	GetGroupedBillingStatements(params *GetGroupedBillingStatementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetGroupedBillingStatementsOK, error)

	GetReconciliationFile(params *GetReconciliationFileParams, authInfo runtime.ClientAuthInfoWriter) (*GetReconciliationFileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetBillingRecordsFile get billing records file API
*/
func (a *Client) GetBillingRecordsFile(params *GetBillingRecordsFileParams, authInfo runtime.ClientAuthInfoWriter) (*GetBillingRecordsFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBillingRecordsFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBillingRecordsFile",
		Method:             "GET",
		PathPattern:        "/api/v1/BillingStatements/{id}/billingrecordsfile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBillingRecordsFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBillingRecordsFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBillingRecordsFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBillingStatementFile get billing statement file API
*/
func (a *Client) GetBillingStatementFile(params *GetBillingStatementFileParams, authInfo runtime.ClientAuthInfoWriter) (*GetBillingStatementFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBillingStatementFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBillingStatementFile",
		Method:             "GET",
		PathPattern:        "/api/v1/BillingStatements/file/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBillingStatementFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBillingStatementFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBillingStatementFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBillingStatements get billing statements API
*/
func (a *Client) GetBillingStatements(params *GetBillingStatementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetBillingStatementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBillingStatementsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBillingStatements",
		Method:             "GET",
		PathPattern:        "/api/v1/BillingStatements",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBillingStatementsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBillingStatementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBillingStatements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGroupedBillingStatements get grouped billing statements API
*/
func (a *Client) GetGroupedBillingStatements(params *GetGroupedBillingStatementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetGroupedBillingStatementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupedBillingStatementsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetGroupedBillingStatements",
		Method:             "GET",
		PathPattern:        "/api/v1/BillingStatements/grouped",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupedBillingStatementsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGroupedBillingStatementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetGroupedBillingStatements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetReconciliationFile get reconciliation file API
*/
func (a *Client) GetReconciliationFile(params *GetReconciliationFileParams, authInfo runtime.ClientAuthInfoWriter) (*GetReconciliationFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReconciliationFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetReconciliationFile",
		Method:             "GET",
		PathPattern:        "/api/v1/BillingStatements/{id}/reconciliationfile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReconciliationFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReconciliationFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetReconciliationFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
