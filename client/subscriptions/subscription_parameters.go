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

// NewSubscriptionParams creates a new SubscriptionParams object
// with the default values initialized.
func NewSubscriptionParams() *SubscriptionParams {
	var ()
	return &SubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionParamsWithTimeout creates a new SubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubscriptionParamsWithTimeout(timeout time.Duration) *SubscriptionParams {
	var ()
	return &SubscriptionParams{

		timeout: timeout,
	}
}

// NewSubscriptionParamsWithContext creates a new SubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubscriptionParamsWithContext(ctx context.Context) *SubscriptionParams {
	var ()
	return &SubscriptionParams{

		Context: ctx,
	}
}

// NewSubscriptionParamsWithHTTPClient creates a new SubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubscriptionParamsWithHTTPClient(client *http.Client) *SubscriptionParams {
	var ()
	return &SubscriptionParams{
		HTTPClient: client,
	}
}

/*SubscriptionParams contains all the parameters to send to the API endpoint
for the subscription operation typically these are written to a http.Request
*/
type SubscriptionParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subscription params
func (o *SubscriptionParams) WithTimeout(timeout time.Duration) *SubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription params
func (o *SubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription params
func (o *SubscriptionParams) WithContext(ctx context.Context) *SubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription params
func (o *SubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription params
func (o *SubscriptionParams) WithHTTPClient(client *http.Client) *SubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription params
func (o *SubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the subscription params
func (o *SubscriptionParams) WithID(id int32) *SubscriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the subscription params
func (o *SubscriptionParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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