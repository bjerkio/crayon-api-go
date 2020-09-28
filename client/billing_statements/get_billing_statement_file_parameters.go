// Code generated by go-swagger; DO NOT EDIT.

package billing_statements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetBillingStatementFileParams creates a new GetBillingStatementFileParams object
// with the default values initialized.
func NewGetBillingStatementFileParams() *GetBillingStatementFileParams {
	var ()
	return &GetBillingStatementFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBillingStatementFileParamsWithTimeout creates a new GetBillingStatementFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBillingStatementFileParamsWithTimeout(timeout time.Duration) *GetBillingStatementFileParams {
	var ()
	return &GetBillingStatementFileParams{

		timeout: timeout,
	}
}

// NewGetBillingStatementFileParamsWithContext creates a new GetBillingStatementFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBillingStatementFileParamsWithContext(ctx context.Context) *GetBillingStatementFileParams {
	var ()
	return &GetBillingStatementFileParams{

		Context: ctx,
	}
}

// NewGetBillingStatementFileParamsWithHTTPClient creates a new GetBillingStatementFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBillingStatementFileParamsWithHTTPClient(client *http.Client) *GetBillingStatementFileParams {
	var ()
	return &GetBillingStatementFileParams{
		HTTPClient: client,
	}
}

/*GetBillingStatementFileParams contains all the parameters to send to the API endpoint
for the get billing statement file operation typically these are written to a http.Request
*/
type GetBillingStatementFileParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get billing statement file params
func (o *GetBillingStatementFileParams) WithTimeout(timeout time.Duration) *GetBillingStatementFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get billing statement file params
func (o *GetBillingStatementFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get billing statement file params
func (o *GetBillingStatementFileParams) WithContext(ctx context.Context) *GetBillingStatementFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get billing statement file params
func (o *GetBillingStatementFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get billing statement file params
func (o *GetBillingStatementFileParams) WithHTTPClient(client *http.Client) *GetBillingStatementFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get billing statement file params
func (o *GetBillingStatementFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get billing statement file params
func (o *GetBillingStatementFileParams) WithID(id int32) *GetBillingStatementFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get billing statement file params
func (o *GetBillingStatementFileParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetBillingStatementFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}