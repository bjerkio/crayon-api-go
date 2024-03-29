// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

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

// NewGetRegisterReservedInstanceParams creates a new GetRegisterReservedInstanceParams object
// with the default values initialized.
func NewGetRegisterReservedInstanceParams() *GetRegisterReservedInstanceParams {
	var ()
	return &GetRegisterReservedInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegisterReservedInstanceParamsWithTimeout creates a new GetRegisterReservedInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRegisterReservedInstanceParamsWithTimeout(timeout time.Duration) *GetRegisterReservedInstanceParams {
	var ()
	return &GetRegisterReservedInstanceParams{

		timeout: timeout,
	}
}

// NewGetRegisterReservedInstanceParamsWithContext creates a new GetRegisterReservedInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRegisterReservedInstanceParamsWithContext(ctx context.Context) *GetRegisterReservedInstanceParams {
	var ()
	return &GetRegisterReservedInstanceParams{

		Context: ctx,
	}
}

// NewGetRegisterReservedInstanceParamsWithHTTPClient creates a new GetRegisterReservedInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRegisterReservedInstanceParamsWithHTTPClient(client *http.Client) *GetRegisterReservedInstanceParams {
	var ()
	return &GetRegisterReservedInstanceParams{
		HTTPClient: client,
	}
}

/*GetRegisterReservedInstanceParams contains all the parameters to send to the API endpoint
for the get register reserved instance operation typically these are written to a http.Request
*/
type GetRegisterReservedInstanceParams struct {

	/*ID*/
	ID int32
	/*ReservedInstance*/
	ReservedInstance bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get register reserved instance params
func (o *GetRegisterReservedInstanceParams) WithTimeout(timeout time.Duration) *GetRegisterReservedInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get register reserved instance params
func (o *GetRegisterReservedInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get register reserved instance params
func (o *GetRegisterReservedInstanceParams) WithContext(ctx context.Context) *GetRegisterReservedInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get register reserved instance params
func (o *GetRegisterReservedInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get register reserved instance params
func (o *GetRegisterReservedInstanceParams) WithHTTPClient(client *http.Client) *GetRegisterReservedInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get register reserved instance params
func (o *GetRegisterReservedInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get register reserved instance params
func (o *GetRegisterReservedInstanceParams) WithID(id int32) *GetRegisterReservedInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get register reserved instance params
func (o *GetRegisterReservedInstanceParams) SetID(id int32) {
	o.ID = id
}

// WithReservedInstance adds the reservedInstance to the get register reserved instance params
func (o *GetRegisterReservedInstanceParams) WithReservedInstance(reservedInstance bool) *GetRegisterReservedInstanceParams {
	o.SetReservedInstance(reservedInstance)
	return o
}

// SetReservedInstance adds the reservedInstance to the get register reserved instance params
func (o *GetRegisterReservedInstanceParams) SetReservedInstance(reservedInstance bool) {
	o.ReservedInstance = reservedInstance
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegisterReservedInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	// path param reservedInstance
	if err := r.SetPathParam("reservedInstance", swag.FormatBool(o.ReservedInstance)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
