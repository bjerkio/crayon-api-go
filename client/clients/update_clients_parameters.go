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

// NewUpdateClientsParams creates a new UpdateClientsParams object
// with the default values initialized.
func NewUpdateClientsParams() *UpdateClientsParams {
	var ()
	return &UpdateClientsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClientsParamsWithTimeout creates a new UpdateClientsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateClientsParamsWithTimeout(timeout time.Duration) *UpdateClientsParams {
	var ()
	return &UpdateClientsParams{

		timeout: timeout,
	}
}

// NewUpdateClientsParamsWithContext creates a new UpdateClientsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateClientsParamsWithContext(ctx context.Context) *UpdateClientsParams {
	var ()
	return &UpdateClientsParams{

		Context: ctx,
	}
}

// NewUpdateClientsParamsWithHTTPClient creates a new UpdateClientsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateClientsParamsWithHTTPClient(client *http.Client) *UpdateClientsParams {
	var ()
	return &UpdateClientsParams{
		HTTPClient: client,
	}
}

/*UpdateClientsParams contains all the parameters to send to the API endpoint
for the update clients operation typically these are written to a http.Request
*/
type UpdateClientsParams struct {

	/*Client*/
	Client *models.Client
	/*ClientID*/
	ClientID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update clients params
func (o *UpdateClientsParams) WithTimeout(timeout time.Duration) *UpdateClientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update clients params
func (o *UpdateClientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update clients params
func (o *UpdateClientsParams) WithContext(ctx context.Context) *UpdateClientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update clients params
func (o *UpdateClientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update clients params
func (o *UpdateClientsParams) WithHTTPClient(client *http.Client) *UpdateClientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update clients params
func (o *UpdateClientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClient adds the client to the update clients params
func (o *UpdateClientsParams) WithClient(client *models.Client) *UpdateClientsParams {
	o.SetClient(client)
	return o
}

// SetClient adds the client to the update clients params
func (o *UpdateClientsParams) SetClient(client *models.Client) {
	o.Client = client
}

// WithClientID adds the clientID to the update clients params
func (o *UpdateClientsParams) WithClientID(clientID string) *UpdateClientsParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the update clients params
func (o *UpdateClientsParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Client != nil {
		if err := r.SetBodyParam(o.Client); err != nil {
			return err
		}
	}

	// path param clientId
	if err := r.SetPathParam("clientId", o.ClientID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
