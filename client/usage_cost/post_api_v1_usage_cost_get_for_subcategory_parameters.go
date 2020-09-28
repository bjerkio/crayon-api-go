// Code generated by go-swagger; DO NOT EDIT.

package usage_cost

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

// NewPostAPIV1UsageCostGetForSubcategoryParams creates a new PostAPIV1UsageCostGetForSubcategoryParams object
// with the default values initialized.
func NewPostAPIV1UsageCostGetForSubcategoryParams() *PostAPIV1UsageCostGetForSubcategoryParams {
	var ()
	return &PostAPIV1UsageCostGetForSubcategoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV1UsageCostGetForSubcategoryParamsWithTimeout creates a new PostAPIV1UsageCostGetForSubcategoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV1UsageCostGetForSubcategoryParamsWithTimeout(timeout time.Duration) *PostAPIV1UsageCostGetForSubcategoryParams {
	var ()
	return &PostAPIV1UsageCostGetForSubcategoryParams{

		timeout: timeout,
	}
}

// NewPostAPIV1UsageCostGetForSubcategoryParamsWithContext creates a new PostAPIV1UsageCostGetForSubcategoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV1UsageCostGetForSubcategoryParamsWithContext(ctx context.Context) *PostAPIV1UsageCostGetForSubcategoryParams {
	var ()
	return &PostAPIV1UsageCostGetForSubcategoryParams{

		Context: ctx,
	}
}

// NewPostAPIV1UsageCostGetForSubcategoryParamsWithHTTPClient creates a new PostAPIV1UsageCostGetForSubcategoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV1UsageCostGetForSubcategoryParamsWithHTTPClient(client *http.Client) *PostAPIV1UsageCostGetForSubcategoryParams {
	var ()
	return &PostAPIV1UsageCostGetForSubcategoryParams{
		HTTPClient: client,
	}
}

/*PostAPIV1UsageCostGetForSubcategoryParams contains all the parameters to send to the API endpoint
for the post API v1 usage cost get for subcategory operation typically these are written to a http.Request
*/
type PostAPIV1UsageCostGetForSubcategoryParams struct {

	/*Model*/
	Model *models.SubcategoryUsageCostRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v1 usage cost get for subcategory params
func (o *PostAPIV1UsageCostGetForSubcategoryParams) WithTimeout(timeout time.Duration) *PostAPIV1UsageCostGetForSubcategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v1 usage cost get for subcategory params
func (o *PostAPIV1UsageCostGetForSubcategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v1 usage cost get for subcategory params
func (o *PostAPIV1UsageCostGetForSubcategoryParams) WithContext(ctx context.Context) *PostAPIV1UsageCostGetForSubcategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v1 usage cost get for subcategory params
func (o *PostAPIV1UsageCostGetForSubcategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v1 usage cost get for subcategory params
func (o *PostAPIV1UsageCostGetForSubcategoryParams) WithHTTPClient(client *http.Client) *PostAPIV1UsageCostGetForSubcategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v1 usage cost get for subcategory params
func (o *PostAPIV1UsageCostGetForSubcategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModel adds the model to the post API v1 usage cost get for subcategory params
func (o *PostAPIV1UsageCostGetForSubcategoryParams) WithModel(model *models.SubcategoryUsageCostRequest) *PostAPIV1UsageCostGetForSubcategoryParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the post API v1 usage cost get for subcategory params
func (o *PostAPIV1UsageCostGetForSubcategoryParams) SetModel(model *models.SubcategoryUsageCostRequest) {
	o.Model = model
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV1UsageCostGetForSubcategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Model != nil {
		if err := r.SetBodyParam(o.Model); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
