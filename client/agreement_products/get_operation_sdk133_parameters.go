// Code generated by go-swagger; DO NOT EDIT.

package agreement_products

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

// NewGetOperationSdk133Params creates a new GetOperationSdk133Params object
// with the default values initialized.
func NewGetOperationSdk133Params() *GetOperationSdk133Params {
	var ()
	return &GetOperationSdk133Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOperationSdk133ParamsWithTimeout creates a new GetOperationSdk133Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOperationSdk133ParamsWithTimeout(timeout time.Duration) *GetOperationSdk133Params {
	var ()
	return &GetOperationSdk133Params{

		timeout: timeout,
	}
}

// NewGetOperationSdk133ParamsWithContext creates a new GetOperationSdk133Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetOperationSdk133ParamsWithContext(ctx context.Context) *GetOperationSdk133Params {
	var ()
	return &GetOperationSdk133Params{

		Context: ctx,
	}
}

// NewGetOperationSdk133ParamsWithHTTPClient creates a new GetOperationSdk133Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOperationSdk133ParamsWithHTTPClient(client *http.Client) *GetOperationSdk133Params {
	var ()
	return &GetOperationSdk133Params{
		HTTPClient: client,
	}
}

/*GetOperationSdk133Params contains all the parameters to send to the API endpoint
for the get operation sdk133 operation typically these are written to a http.Request
*/
type GetOperationSdk133Params struct {

	/*Filter*/
	Filter *models.AgreementProductFilter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get operation sdk133 params
func (o *GetOperationSdk133Params) WithTimeout(timeout time.Duration) *GetOperationSdk133Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get operation sdk133 params
func (o *GetOperationSdk133Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get operation sdk133 params
func (o *GetOperationSdk133Params) WithContext(ctx context.Context) *GetOperationSdk133Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get operation sdk133 params
func (o *GetOperationSdk133Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get operation sdk133 params
func (o *GetOperationSdk133Params) WithHTTPClient(client *http.Client) *GetOperationSdk133Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get operation sdk133 params
func (o *GetOperationSdk133Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get operation sdk133 params
func (o *GetOperationSdk133Params) WithFilter(filter *models.AgreementProductFilter) *GetOperationSdk133Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get operation sdk133 params
func (o *GetOperationSdk133Params) SetFilter(filter *models.AgreementProductFilter) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *GetOperationSdk133Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {
		if err := r.SetBodyParam(o.Filter); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
