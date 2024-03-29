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

// NewRegisterReservedInstanceParams creates a new RegisterReservedInstanceParams object
// with the default values initialized.
func NewRegisterReservedInstanceParams() *RegisterReservedInstanceParams {
	var ()
	return &RegisterReservedInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterReservedInstanceParamsWithTimeout creates a new RegisterReservedInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegisterReservedInstanceParamsWithTimeout(timeout time.Duration) *RegisterReservedInstanceParams {
	var ()
	return &RegisterReservedInstanceParams{

		timeout: timeout,
	}
}

// NewRegisterReservedInstanceParamsWithContext creates a new RegisterReservedInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewRegisterReservedInstanceParamsWithContext(ctx context.Context) *RegisterReservedInstanceParams {
	var ()
	return &RegisterReservedInstanceParams{

		Context: ctx,
	}
}

// NewRegisterReservedInstanceParamsWithHTTPClient creates a new RegisterReservedInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegisterReservedInstanceParamsWithHTTPClient(client *http.Client) *RegisterReservedInstanceParams {
	var ()
	return &RegisterReservedInstanceParams{
		HTTPClient: client,
	}
}

/*RegisterReservedInstanceParams contains all the parameters to send to the API endpoint
for the register reserved instance operation typically these are written to a http.Request
*/
type RegisterReservedInstanceParams struct {

	/*ID*/
	ID int32
	/*ReservedInstance*/
	ReservedInstance bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the register reserved instance params
func (o *RegisterReservedInstanceParams) WithTimeout(timeout time.Duration) *RegisterReservedInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register reserved instance params
func (o *RegisterReservedInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register reserved instance params
func (o *RegisterReservedInstanceParams) WithContext(ctx context.Context) *RegisterReservedInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register reserved instance params
func (o *RegisterReservedInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register reserved instance params
func (o *RegisterReservedInstanceParams) WithHTTPClient(client *http.Client) *RegisterReservedInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register reserved instance params
func (o *RegisterReservedInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the register reserved instance params
func (o *RegisterReservedInstanceParams) WithID(id int32) *RegisterReservedInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the register reserved instance params
func (o *RegisterReservedInstanceParams) SetID(id int32) {
	o.ID = id
}

// WithReservedInstance adds the reservedInstance to the register reserved instance params
func (o *RegisterReservedInstanceParams) WithReservedInstance(reservedInstance bool) *RegisterReservedInstanceParams {
	o.SetReservedInstance(reservedInstance)
	return o
}

// SetReservedInstance adds the reservedInstance to the register reserved instance params
func (o *RegisterReservedInstanceParams) SetReservedInstance(reservedInstance bool) {
	o.ReservedInstance = reservedInstance
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterReservedInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
