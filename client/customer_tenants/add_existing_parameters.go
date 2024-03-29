// Code generated by go-swagger; DO NOT EDIT.

package customer_tenants

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

// NewAddExistingParams creates a new AddExistingParams object
// with the default values initialized.
func NewAddExistingParams() *AddExistingParams {
	var (
		syncFromPublisherDefault = bool(false)
	)
	return &AddExistingParams{
		SyncFromPublisher: &syncFromPublisherDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewAddExistingParamsWithTimeout creates a new AddExistingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddExistingParamsWithTimeout(timeout time.Duration) *AddExistingParams {
	var (
		syncFromPublisherDefault = bool(false)
	)
	return &AddExistingParams{
		SyncFromPublisher: &syncFromPublisherDefault,

		timeout: timeout,
	}
}

// NewAddExistingParamsWithContext creates a new AddExistingParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddExistingParamsWithContext(ctx context.Context) *AddExistingParams {
	var (
		syncFromPublisherDefault = bool(false)
	)
	return &AddExistingParams{
		SyncFromPublisher: &syncFromPublisherDefault,

		Context: ctx,
	}
}

// NewAddExistingParamsWithHTTPClient creates a new AddExistingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddExistingParamsWithHTTPClient(client *http.Client) *AddExistingParams {
	var (
		syncFromPublisherDefault = bool(false)
	)
	return &AddExistingParams{
		SyncFromPublisher: &syncFromPublisherDefault,
		HTTPClient:        client,
	}
}

/*AddExistingParams contains all the parameters to send to the API endpoint
for the add existing operation typically these are written to a http.Request
*/
type AddExistingParams struct {

	/*ExistingTenant*/
	ExistingTenant *models.CustomerTenantDetailed
	/*SyncFromPublisher*/
	SyncFromPublisher *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add existing params
func (o *AddExistingParams) WithTimeout(timeout time.Duration) *AddExistingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add existing params
func (o *AddExistingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add existing params
func (o *AddExistingParams) WithContext(ctx context.Context) *AddExistingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add existing params
func (o *AddExistingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add existing params
func (o *AddExistingParams) WithHTTPClient(client *http.Client) *AddExistingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add existing params
func (o *AddExistingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExistingTenant adds the existingTenant to the add existing params
func (o *AddExistingParams) WithExistingTenant(existingTenant *models.CustomerTenantDetailed) *AddExistingParams {
	o.SetExistingTenant(existingTenant)
	return o
}

// SetExistingTenant adds the existingTenant to the add existing params
func (o *AddExistingParams) SetExistingTenant(existingTenant *models.CustomerTenantDetailed) {
	o.ExistingTenant = existingTenant
}

// WithSyncFromPublisher adds the syncFromPublisher to the add existing params
func (o *AddExistingParams) WithSyncFromPublisher(syncFromPublisher *bool) *AddExistingParams {
	o.SetSyncFromPublisher(syncFromPublisher)
	return o
}

// SetSyncFromPublisher adds the syncFromPublisher to the add existing params
func (o *AddExistingParams) SetSyncFromPublisher(syncFromPublisher *bool) {
	o.SyncFromPublisher = syncFromPublisher
}

// WriteToRequest writes these params to a swagger request
func (o *AddExistingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExistingTenant != nil {
		if err := r.SetBodyParam(o.ExistingTenant); err != nil {
			return err
		}
	}

	if o.SyncFromPublisher != nil {

		// query param syncFromPublisher
		var qrSyncFromPublisher bool
		if o.SyncFromPublisher != nil {
			qrSyncFromPublisher = *o.SyncFromPublisher
		}
		qSyncFromPublisher := swag.FormatBool(qrSyncFromPublisher)
		if qSyncFromPublisher != "" {
			if err := r.SetQueryParam("syncFromPublisher", qSyncFromPublisher); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
