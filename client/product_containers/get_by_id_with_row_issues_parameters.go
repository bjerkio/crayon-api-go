// Code generated by go-swagger; DO NOT EDIT.

package product_containers

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

// NewGetByIDWithRowIssuesParams creates a new GetByIDWithRowIssuesParams object
// with the default values initialized.
func NewGetByIDWithRowIssuesParams() *GetByIDWithRowIssuesParams {
	var ()
	return &GetByIDWithRowIssuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetByIDWithRowIssuesParamsWithTimeout creates a new GetByIDWithRowIssuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetByIDWithRowIssuesParamsWithTimeout(timeout time.Duration) *GetByIDWithRowIssuesParams {
	var ()
	return &GetByIDWithRowIssuesParams{

		timeout: timeout,
	}
}

// NewGetByIDWithRowIssuesParamsWithContext creates a new GetByIDWithRowIssuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetByIDWithRowIssuesParamsWithContext(ctx context.Context) *GetByIDWithRowIssuesParams {
	var ()
	return &GetByIDWithRowIssuesParams{

		Context: ctx,
	}
}

// NewGetByIDWithRowIssuesParamsWithHTTPClient creates a new GetByIDWithRowIssuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetByIDWithRowIssuesParamsWithHTTPClient(client *http.Client) *GetByIDWithRowIssuesParams {
	var ()
	return &GetByIDWithRowIssuesParams{
		HTTPClient: client,
	}
}

/*GetByIDWithRowIssuesParams contains all the parameters to send to the API endpoint
for the get by Id with row issues operation typically these are written to a http.Request
*/
type GetByIDWithRowIssuesParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get by Id with row issues params
func (o *GetByIDWithRowIssuesParams) WithTimeout(timeout time.Duration) *GetByIDWithRowIssuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get by Id with row issues params
func (o *GetByIDWithRowIssuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get by Id with row issues params
func (o *GetByIDWithRowIssuesParams) WithContext(ctx context.Context) *GetByIDWithRowIssuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get by Id with row issues params
func (o *GetByIDWithRowIssuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get by Id with row issues params
func (o *GetByIDWithRowIssuesParams) WithHTTPClient(client *http.Client) *GetByIDWithRowIssuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get by Id with row issues params
func (o *GetByIDWithRowIssuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get by Id with row issues params
func (o *GetByIDWithRowIssuesParams) WithID(id int32) *GetByIDWithRowIssuesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get by Id with row issues params
func (o *GetByIDWithRowIssuesParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetByIDWithRowIssuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
