// Code generated by go-swagger; DO NOT EDIT.

package organization_access

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

// NewPutOrganizationAccessParams creates a new PutOrganizationAccessParams object
// with the default values initialized.
func NewPutOrganizationAccessParams() *PutOrganizationAccessParams {
	var ()
	return &PutOrganizationAccessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrganizationAccessParamsWithTimeout creates a new PutOrganizationAccessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutOrganizationAccessParamsWithTimeout(timeout time.Duration) *PutOrganizationAccessParams {
	var ()
	return &PutOrganizationAccessParams{

		timeout: timeout,
	}
}

// NewPutOrganizationAccessParamsWithContext creates a new PutOrganizationAccessParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutOrganizationAccessParamsWithContext(ctx context.Context) *PutOrganizationAccessParams {
	var ()
	return &PutOrganizationAccessParams{

		Context: ctx,
	}
}

// NewPutOrganizationAccessParamsWithHTTPClient creates a new PutOrganizationAccessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutOrganizationAccessParamsWithHTTPClient(client *http.Client) *PutOrganizationAccessParams {
	var ()
	return &PutOrganizationAccessParams{
		HTTPClient: client,
	}
}

/*PutOrganizationAccessParams contains all the parameters to send to the API endpoint
for the put organization access operation typically these are written to a http.Request
*/
type PutOrganizationAccessParams struct {

	/*List*/
	List []*models.OrganizationAccess

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put organization access params
func (o *PutOrganizationAccessParams) WithTimeout(timeout time.Duration) *PutOrganizationAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put organization access params
func (o *PutOrganizationAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put organization access params
func (o *PutOrganizationAccessParams) WithContext(ctx context.Context) *PutOrganizationAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put organization access params
func (o *PutOrganizationAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put organization access params
func (o *PutOrganizationAccessParams) WithHTTPClient(client *http.Client) *PutOrganizationAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put organization access params
func (o *PutOrganizationAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithList adds the list to the put organization access params
func (o *PutOrganizationAccessParams) WithList(list []*models.OrganizationAccess) *PutOrganizationAccessParams {
	o.SetList(list)
	return o
}

// SetList adds the list to the put organization access params
func (o *PutOrganizationAccessParams) SetList(list []*models.OrganizationAccess) {
	o.List = list
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrganizationAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.List != nil {
		if err := r.SetBodyParam(o.List); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}