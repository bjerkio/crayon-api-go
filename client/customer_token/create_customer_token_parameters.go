// Code generated by go-swagger; DO NOT EDIT.

package customer_token

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
)

// NewCreateCustomerTokenParams creates a new CreateCustomerTokenParams object
// with the default values initialized.
func NewCreateCustomerTokenParams() *CreateCustomerTokenParams {
	var ()
	return &CreateCustomerTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCustomerTokenParamsWithTimeout creates a new CreateCustomerTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCustomerTokenParamsWithTimeout(timeout time.Duration) *CreateCustomerTokenParams {
	var ()
	return &CreateCustomerTokenParams{

		timeout: timeout,
	}
}

// NewCreateCustomerTokenParamsWithContext creates a new CreateCustomerTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCustomerTokenParamsWithContext(ctx context.Context) *CreateCustomerTokenParams {
	var ()
	return &CreateCustomerTokenParams{

		Context: ctx,
	}
}

// NewCreateCustomerTokenParamsWithHTTPClient creates a new CreateCustomerTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCustomerTokenParamsWithHTTPClient(client *http.Client) *CreateCustomerTokenParams {
	var ()
	return &CreateCustomerTokenParams{
		HTTPClient: client,
	}
}

/*CreateCustomerTokenParams contains all the parameters to send to the API endpoint
for the create customer token operation typically these are written to a http.Request
*/
type CreateCustomerTokenParams struct {

	/*GrantType*/
	GrantType interface{}
	/*Password*/
	Password interface{}
	/*Scope*/
	Scope interface{}
	/*Username*/
	Username interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create customer token params
func (o *CreateCustomerTokenParams) WithTimeout(timeout time.Duration) *CreateCustomerTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create customer token params
func (o *CreateCustomerTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create customer token params
func (o *CreateCustomerTokenParams) WithContext(ctx context.Context) *CreateCustomerTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create customer token params
func (o *CreateCustomerTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create customer token params
func (o *CreateCustomerTokenParams) WithHTTPClient(client *http.Client) *CreateCustomerTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create customer token params
func (o *CreateCustomerTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGrantType adds the grantType to the create customer token params
func (o *CreateCustomerTokenParams) WithGrantType(grantType interface{}) *CreateCustomerTokenParams {
	o.SetGrantType(grantType)
	return o
}

// SetGrantType adds the grantType to the create customer token params
func (o *CreateCustomerTokenParams) SetGrantType(grantType interface{}) {
	o.GrantType = grantType
}

// WithPassword adds the password to the create customer token params
func (o *CreateCustomerTokenParams) WithPassword(password interface{}) *CreateCustomerTokenParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the create customer token params
func (o *CreateCustomerTokenParams) SetPassword(password interface{}) {
	o.Password = password
}

// WithScope adds the scope to the create customer token params
func (o *CreateCustomerTokenParams) WithScope(scope interface{}) *CreateCustomerTokenParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the create customer token params
func (o *CreateCustomerTokenParams) SetScope(scope interface{}) {
	o.Scope = scope
}

// WithUsername adds the username to the create customer token params
func (o *CreateCustomerTokenParams) WithUsername(username interface{}) *CreateCustomerTokenParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the create customer token params
func (o *CreateCustomerTokenParams) SetUsername(username interface{}) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCustomerTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GrantType != nil {
		if err := r.SetBodyParam(o.GrantType); err != nil {
			return err
		}
	}

	if o.Password != nil {
		if err := r.SetBodyParam(o.Password); err != nil {
			return err
		}
	}

	if o.Scope != nil {
		if err := r.SetBodyParam(o.Scope); err != nil {
			return err
		}
	}

	if o.Username != nil {
		if err := r.SetBodyParam(o.Username); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
