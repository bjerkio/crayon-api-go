// Code generated by go-swagger; DO NOT EDIT.

package invoice_profiles

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

// NewDeleteInvoiceProfilesParams creates a new DeleteInvoiceProfilesParams object
// with the default values initialized.
func NewDeleteInvoiceProfilesParams() *DeleteInvoiceProfilesParams {
	var ()
	return &DeleteInvoiceProfilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInvoiceProfilesParamsWithTimeout creates a new DeleteInvoiceProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInvoiceProfilesParamsWithTimeout(timeout time.Duration) *DeleteInvoiceProfilesParams {
	var ()
	return &DeleteInvoiceProfilesParams{

		timeout: timeout,
	}
}

// NewDeleteInvoiceProfilesParamsWithContext creates a new DeleteInvoiceProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInvoiceProfilesParamsWithContext(ctx context.Context) *DeleteInvoiceProfilesParams {
	var ()
	return &DeleteInvoiceProfilesParams{

		Context: ctx,
	}
}

// NewDeleteInvoiceProfilesParamsWithHTTPClient creates a new DeleteInvoiceProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInvoiceProfilesParamsWithHTTPClient(client *http.Client) *DeleteInvoiceProfilesParams {
	var ()
	return &DeleteInvoiceProfilesParams{
		HTTPClient: client,
	}
}

/*DeleteInvoiceProfilesParams contains all the parameters to send to the API endpoint
for the delete invoice profiles operation typically these are written to a http.Request
*/
type DeleteInvoiceProfilesParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete invoice profiles params
func (o *DeleteInvoiceProfilesParams) WithTimeout(timeout time.Duration) *DeleteInvoiceProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete invoice profiles params
func (o *DeleteInvoiceProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete invoice profiles params
func (o *DeleteInvoiceProfilesParams) WithContext(ctx context.Context) *DeleteInvoiceProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete invoice profiles params
func (o *DeleteInvoiceProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete invoice profiles params
func (o *DeleteInvoiceProfilesParams) WithHTTPClient(client *http.Client) *DeleteInvoiceProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete invoice profiles params
func (o *DeleteInvoiceProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete invoice profiles params
func (o *DeleteInvoiceProfilesParams) WithID(id int32) *DeleteInvoiceProfilesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete invoice profiles params
func (o *DeleteInvoiceProfilesParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInvoiceProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
