// Code generated by go-swagger; DO NOT EDIT.

package groupings

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

// NewDeleteGroupingsParams creates a new DeleteGroupingsParams object
// with the default values initialized.
func NewDeleteGroupingsParams() *DeleteGroupingsParams {
	var ()
	return &DeleteGroupingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGroupingsParamsWithTimeout creates a new DeleteGroupingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteGroupingsParamsWithTimeout(timeout time.Duration) *DeleteGroupingsParams {
	var ()
	return &DeleteGroupingsParams{

		timeout: timeout,
	}
}

// NewDeleteGroupingsParamsWithContext creates a new DeleteGroupingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteGroupingsParamsWithContext(ctx context.Context) *DeleteGroupingsParams {
	var ()
	return &DeleteGroupingsParams{

		Context: ctx,
	}
}

// NewDeleteGroupingsParamsWithHTTPClient creates a new DeleteGroupingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteGroupingsParamsWithHTTPClient(client *http.Client) *DeleteGroupingsParams {
	var ()
	return &DeleteGroupingsParams{
		HTTPClient: client,
	}
}

/*DeleteGroupingsParams contains all the parameters to send to the API endpoint
for the delete groupings operation typically these are written to a http.Request
*/
type DeleteGroupingsParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete groupings params
func (o *DeleteGroupingsParams) WithTimeout(timeout time.Duration) *DeleteGroupingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete groupings params
func (o *DeleteGroupingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete groupings params
func (o *DeleteGroupingsParams) WithContext(ctx context.Context) *DeleteGroupingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete groupings params
func (o *DeleteGroupingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete groupings params
func (o *DeleteGroupingsParams) WithHTTPClient(client *http.Client) *DeleteGroupingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete groupings params
func (o *DeleteGroupingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete groupings params
func (o *DeleteGroupingsParams) WithID(id int32) *DeleteGroupingsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete groupings params
func (o *DeleteGroupingsParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGroupingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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