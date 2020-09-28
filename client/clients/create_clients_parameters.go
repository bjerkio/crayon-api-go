// Code generated by go-swagger; DO NOT EDIT.

package clients

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

// NewCreateClientsParams creates a new CreateClientsParams object
// with the default values initialized.
func NewCreateClientsParams() *CreateClientsParams {
	var ()
	return &CreateClientsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClientsParamsWithTimeout creates a new CreateClientsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClientsParamsWithTimeout(timeout time.Duration) *CreateClientsParams {
	var ()
	return &CreateClientsParams{

		timeout: timeout,
	}
}

// NewCreateClientsParamsWithContext creates a new CreateClientsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateClientsParamsWithContext(ctx context.Context) *CreateClientsParams {
	var ()
	return &CreateClientsParams{

		Context: ctx,
	}
}

// NewCreateClientsParamsWithHTTPClient creates a new CreateClientsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateClientsParamsWithHTTPClient(client *http.Client) *CreateClientsParams {
	var ()
	return &CreateClientsParams{
		HTTPClient: client,
	}
}

/*CreateClientsParams contains all the parameters to send to the API endpoint
for the create clients operation typically these are written to a http.Request
*/
type CreateClientsParams struct {

	/*Client*/
	Client *models.Client

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create clients params
func (o *CreateClientsParams) WithTimeout(timeout time.Duration) *CreateClientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create clients params
func (o *CreateClientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create clients params
func (o *CreateClientsParams) WithContext(ctx context.Context) *CreateClientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create clients params
func (o *CreateClientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create clients params
func (o *CreateClientsParams) WithHTTPClient(client *http.Client) *CreateClientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create clients params
func (o *CreateClientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClient adds the client to the create clients params
func (o *CreateClientsParams) WithClient(client *models.Client) *CreateClientsParams {
	o.SetClient(client)
	return o
}

// SetClient adds the client to the create clients params
func (o *CreateClientsParams) SetClient(client *models.Client) {
	o.Client = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Client != nil {
		if err := r.SetBodyParam(o.Client); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
