// Code generated by go-swagger; DO NOT EDIT.

package customer_tenants

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

// NewGetCustomerTenantByIDParams creates a new GetCustomerTenantByIDParams object
// with the default values initialized.
func NewGetCustomerTenantByIDParams() *GetCustomerTenantByIDParams {
	var ()
	return &GetCustomerTenantByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomerTenantByIDParamsWithTimeout creates a new GetCustomerTenantByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomerTenantByIDParamsWithTimeout(timeout time.Duration) *GetCustomerTenantByIDParams {
	var ()
	return &GetCustomerTenantByIDParams{

		timeout: timeout,
	}
}

// NewGetCustomerTenantByIDParamsWithContext creates a new GetCustomerTenantByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomerTenantByIDParamsWithContext(ctx context.Context) *GetCustomerTenantByIDParams {
	var ()
	return &GetCustomerTenantByIDParams{

		Context: ctx,
	}
}

// NewGetCustomerTenantByIDParamsWithHTTPClient creates a new GetCustomerTenantByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomerTenantByIDParamsWithHTTPClient(client *http.Client) *GetCustomerTenantByIDParams {
	var ()
	return &GetCustomerTenantByIDParams{
		HTTPClient: client,
	}
}

/*GetCustomerTenantByIDParams contains all the parameters to send to the API endpoint
for the get customer tenant by Id operation typically these are written to a http.Request
*/
type GetCustomerTenantByIDParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get customer tenant by Id params
func (o *GetCustomerTenantByIDParams) WithTimeout(timeout time.Duration) *GetCustomerTenantByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customer tenant by Id params
func (o *GetCustomerTenantByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customer tenant by Id params
func (o *GetCustomerTenantByIDParams) WithContext(ctx context.Context) *GetCustomerTenantByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customer tenant by Id params
func (o *GetCustomerTenantByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customer tenant by Id params
func (o *GetCustomerTenantByIDParams) WithHTTPClient(client *http.Client) *GetCustomerTenantByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customer tenant by Id params
func (o *GetCustomerTenantByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get customer tenant by Id params
func (o *GetCustomerTenantByIDParams) WithID(id int32) *GetCustomerTenantByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get customer tenant by Id params
func (o *GetCustomerTenantByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomerTenantByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
