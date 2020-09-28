// Code generated by go-swagger; DO NOT EDIT.

package product_containers

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

// NewGetOrCreateShoppingCartParams creates a new GetOrCreateShoppingCartParams object
// with the default values initialized.
func NewGetOrCreateShoppingCartParams() *GetOrCreateShoppingCartParams {
	var ()
	return &GetOrCreateShoppingCartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrCreateShoppingCartParamsWithTimeout creates a new GetOrCreateShoppingCartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrCreateShoppingCartParamsWithTimeout(timeout time.Duration) *GetOrCreateShoppingCartParams {
	var ()
	return &GetOrCreateShoppingCartParams{

		timeout: timeout,
	}
}

// NewGetOrCreateShoppingCartParamsWithContext creates a new GetOrCreateShoppingCartParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrCreateShoppingCartParamsWithContext(ctx context.Context) *GetOrCreateShoppingCartParams {
	var ()
	return &GetOrCreateShoppingCartParams{

		Context: ctx,
	}
}

// NewGetOrCreateShoppingCartParamsWithHTTPClient creates a new GetOrCreateShoppingCartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrCreateShoppingCartParamsWithHTTPClient(client *http.Client) *GetOrCreateShoppingCartParams {
	var ()
	return &GetOrCreateShoppingCartParams{
		HTTPClient: client,
	}
}

/*GetOrCreateShoppingCartParams contains all the parameters to send to the API endpoint
for the get or create shopping cart operation typically these are written to a http.Request
*/
type GetOrCreateShoppingCartParams struct {

	/*OrganizationID*/
	OrganizationID *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get or create shopping cart params
func (o *GetOrCreateShoppingCartParams) WithTimeout(timeout time.Duration) *GetOrCreateShoppingCartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get or create shopping cart params
func (o *GetOrCreateShoppingCartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get or create shopping cart params
func (o *GetOrCreateShoppingCartParams) WithContext(ctx context.Context) *GetOrCreateShoppingCartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get or create shopping cart params
func (o *GetOrCreateShoppingCartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get or create shopping cart params
func (o *GetOrCreateShoppingCartParams) WithHTTPClient(client *http.Client) *GetOrCreateShoppingCartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get or create shopping cart params
func (o *GetOrCreateShoppingCartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get or create shopping cart params
func (o *GetOrCreateShoppingCartParams) WithOrganizationID(organizationID *int32) *GetOrCreateShoppingCartParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get or create shopping cart params
func (o *GetOrCreateShoppingCartParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrCreateShoppingCartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrganizationID != nil {

		// query param organizationId
		var qrOrganizationID int32
		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := swag.FormatInt32(qrOrganizationID)
		if qOrganizationID != "" {
			if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
