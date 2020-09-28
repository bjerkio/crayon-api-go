// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewChangePasswordParams creates a new ChangePasswordParams object
// with the default values initialized.
func NewChangePasswordParams() *ChangePasswordParams {
	var ()
	return &ChangePasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangePasswordParamsWithTimeout creates a new ChangePasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangePasswordParamsWithTimeout(timeout time.Duration) *ChangePasswordParams {
	var ()
	return &ChangePasswordParams{

		timeout: timeout,
	}
}

// NewChangePasswordParamsWithContext creates a new ChangePasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangePasswordParamsWithContext(ctx context.Context) *ChangePasswordParams {
	var ()
	return &ChangePasswordParams{

		Context: ctx,
	}
}

// NewChangePasswordParamsWithHTTPClient creates a new ChangePasswordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangePasswordParamsWithHTTPClient(client *http.Client) *ChangePasswordParams {
	var ()
	return &ChangePasswordParams{
		HTTPClient: client,
	}
}

/*ChangePasswordParams contains all the parameters to send to the API endpoint
for the change password operation typically these are written to a http.Request
*/
type ChangePasswordParams struct {

	/*ChangePassword*/
	ChangePassword *models.UserChangePassword
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change password params
func (o *ChangePasswordParams) WithTimeout(timeout time.Duration) *ChangePasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change password params
func (o *ChangePasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change password params
func (o *ChangePasswordParams) WithContext(ctx context.Context) *ChangePasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change password params
func (o *ChangePasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change password params
func (o *ChangePasswordParams) WithHTTPClient(client *http.Client) *ChangePasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change password params
func (o *ChangePasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangePassword adds the changePassword to the change password params
func (o *ChangePasswordParams) WithChangePassword(changePassword *models.UserChangePassword) *ChangePasswordParams {
	o.SetChangePassword(changePassword)
	return o
}

// SetChangePassword adds the changePassword to the change password params
func (o *ChangePasswordParams) SetChangePassword(changePassword *models.UserChangePassword) {
	o.ChangePassword = changePassword
}

// WithID adds the id to the change password params
func (o *ChangePasswordParams) WithID(id string) *ChangePasswordParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the change password params
func (o *ChangePasswordParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangePasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChangePassword != nil {
		if err := r.SetBodyParam(o.ChangePassword); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
