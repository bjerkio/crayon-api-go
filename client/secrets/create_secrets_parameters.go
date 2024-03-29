// Code generated by go-swagger; DO NOT EDIT.

package secrets

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

	"github.com/bjerkio/crayon-api-go/models"
)

// NewCreateSecretsParams creates a new CreateSecretsParams object
// with the default values initialized.
func NewCreateSecretsParams() *CreateSecretsParams {
	var ()
	return &CreateSecretsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSecretsParamsWithTimeout creates a new CreateSecretsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSecretsParamsWithTimeout(timeout time.Duration) *CreateSecretsParams {
	var ()
	return &CreateSecretsParams{

		timeout: timeout,
	}
}

// NewCreateSecretsParamsWithContext creates a new CreateSecretsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSecretsParamsWithContext(ctx context.Context) *CreateSecretsParams {
	var ()
	return &CreateSecretsParams{

		Context: ctx,
	}
}

// NewCreateSecretsParamsWithHTTPClient creates a new CreateSecretsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSecretsParamsWithHTTPClient(client *http.Client) *CreateSecretsParams {
	var ()
	return &CreateSecretsParams{
		HTTPClient: client,
	}
}

/*CreateSecretsParams contains all the parameters to send to the API endpoint
for the create secrets operation typically these are written to a http.Request
*/
type CreateSecretsParams struct {

	/*Secret*/
	Secret *models.Secret

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create secrets params
func (o *CreateSecretsParams) WithTimeout(timeout time.Duration) *CreateSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create secrets params
func (o *CreateSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create secrets params
func (o *CreateSecretsParams) WithContext(ctx context.Context) *CreateSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create secrets params
func (o *CreateSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create secrets params
func (o *CreateSecretsParams) WithHTTPClient(client *http.Client) *CreateSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create secrets params
func (o *CreateSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSecret adds the secret to the create secrets params
func (o *CreateSecretsParams) WithSecret(secret *models.Secret) *CreateSecretsParams {
	o.SetSecret(secret)
	return o
}

// SetSecret adds the secret to the create secrets params
func (o *CreateSecretsParams) SetSecret(secret *models.Secret) {
	o.Secret = secret
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Secret != nil {
		if err := r.SetBodyParam(o.Secret); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
