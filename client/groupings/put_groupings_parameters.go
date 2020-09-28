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

	"github.com/bjerkio/crayon-api-go/models"
)

// NewPutGroupingsParams creates a new PutGroupingsParams object
// with the default values initialized.
func NewPutGroupingsParams() *PutGroupingsParams {
	var ()
	return &PutGroupingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutGroupingsParamsWithTimeout creates a new PutGroupingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutGroupingsParamsWithTimeout(timeout time.Duration) *PutGroupingsParams {
	var ()
	return &PutGroupingsParams{

		timeout: timeout,
	}
}

// NewPutGroupingsParamsWithContext creates a new PutGroupingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutGroupingsParamsWithContext(ctx context.Context) *PutGroupingsParams {
	var ()
	return &PutGroupingsParams{

		Context: ctx,
	}
}

// NewPutGroupingsParamsWithHTTPClient creates a new PutGroupingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutGroupingsParamsWithHTTPClient(client *http.Client) *PutGroupingsParams {
	var ()
	return &PutGroupingsParams{
		HTTPClient: client,
	}
}

/*PutGroupingsParams contains all the parameters to send to the API endpoint
for the put groupings operation typically these are written to a http.Request
*/
type PutGroupingsParams struct {

	/*Grouping*/
	Grouping *models.Grouping
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put groupings params
func (o *PutGroupingsParams) WithTimeout(timeout time.Duration) *PutGroupingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put groupings params
func (o *PutGroupingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put groupings params
func (o *PutGroupingsParams) WithContext(ctx context.Context) *PutGroupingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put groupings params
func (o *PutGroupingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put groupings params
func (o *PutGroupingsParams) WithHTTPClient(client *http.Client) *PutGroupingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put groupings params
func (o *PutGroupingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGrouping adds the grouping to the put groupings params
func (o *PutGroupingsParams) WithGrouping(grouping *models.Grouping) *PutGroupingsParams {
	o.SetGrouping(grouping)
	return o
}

// SetGrouping adds the grouping to the put groupings params
func (o *PutGroupingsParams) SetGrouping(grouping *models.Grouping) {
	o.Grouping = grouping
}

// WithID adds the id to the put groupings params
func (o *PutGroupingsParams) WithID(id int32) *PutGroupingsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put groupings params
func (o *PutGroupingsParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutGroupingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Grouping != nil {
		if err := r.SetBodyParam(o.Grouping); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}