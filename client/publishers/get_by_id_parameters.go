// Code generated by go-swagger; DO NOT EDIT.

package publishers

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

// NewGetByIDParams creates a new GetByIDParams object
// with the default values initialized.
func NewGetByIDParams() *GetByIDParams {
	var ()
	return &GetByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetByIDParamsWithTimeout creates a new GetByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetByIDParamsWithTimeout(timeout time.Duration) *GetByIDParams {
	var ()
	return &GetByIDParams{

		timeout: timeout,
	}
}

// NewGetByIDParamsWithContext creates a new GetByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetByIDParamsWithContext(ctx context.Context) *GetByIDParams {
	var ()
	return &GetByIDParams{

		Context: ctx,
	}
}

// NewGetByIDParamsWithHTTPClient creates a new GetByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetByIDParamsWithHTTPClient(client *http.Client) *GetByIDParams {
	var ()
	return &GetByIDParams{
		HTTPClient: client,
	}
}

/*GetByIDParams contains all the parameters to send to the API endpoint
for the get by Id operation typically these are written to a http.Request
*/
type GetByIDParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get by Id params
func (o *GetByIDParams) WithTimeout(timeout time.Duration) *GetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get by Id params
func (o *GetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get by Id params
func (o *GetByIDParams) WithContext(ctx context.Context) *GetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get by Id params
func (o *GetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get by Id params
func (o *GetByIDParams) WithHTTPClient(client *http.Client) *GetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get by Id params
func (o *GetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get by Id params
func (o *GetByIDParams) WithID(id int32) *GetByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get by Id params
func (o *GetByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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