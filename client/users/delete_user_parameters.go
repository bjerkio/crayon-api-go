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
)

// NewDeleteUserParams creates a new DeleteUserParams object
// with the default values initialized.
func NewDeleteUserParams() *DeleteUserParams {
	var ()
	return &DeleteUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserParamsWithTimeout creates a new DeleteUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserParamsWithTimeout(timeout time.Duration) *DeleteUserParams {
	var ()
	return &DeleteUserParams{

		timeout: timeout,
	}
}

// NewDeleteUserParamsWithContext creates a new DeleteUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserParamsWithContext(ctx context.Context) *DeleteUserParams {
	var ()
	return &DeleteUserParams{

		Context: ctx,
	}
}

// NewDeleteUserParamsWithHTTPClient creates a new DeleteUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserParamsWithHTTPClient(client *http.Client) *DeleteUserParams {
	var ()
	return &DeleteUserParams{
		HTTPClient: client,
	}
}

/*DeleteUserParams contains all the parameters to send to the API endpoint
for the delete user operation typically these are written to a http.Request
*/
type DeleteUserParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user params
func (o *DeleteUserParams) WithTimeout(timeout time.Duration) *DeleteUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user params
func (o *DeleteUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user params
func (o *DeleteUserParams) WithContext(ctx context.Context) *DeleteUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user params
func (o *DeleteUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user params
func (o *DeleteUserParams) WithHTTPClient(client *http.Client) *DeleteUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user params
func (o *DeleteUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete user params
func (o *DeleteUserParams) WithID(id string) *DeleteUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete user params
func (o *DeleteUserParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
