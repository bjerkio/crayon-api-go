// Code generated by go-swagger; DO NOT EDIT.

package blog_items

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

// NewGetForDirectCustomersParams creates a new GetForDirectCustomersParams object
// with the default values initialized.
func NewGetForDirectCustomersParams() *GetForDirectCustomersParams {
	var (
		countDefault          = int32(6)
		organizationIDDefault = int32(0)
	)
	return &GetForDirectCustomersParams{
		Count:          &countDefault,
		OrganizationID: &organizationIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetForDirectCustomersParamsWithTimeout creates a new GetForDirectCustomersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetForDirectCustomersParamsWithTimeout(timeout time.Duration) *GetForDirectCustomersParams {
	var (
		countDefault          = int32(6)
		organizationIDDefault = int32(0)
	)
	return &GetForDirectCustomersParams{
		Count:          &countDefault,
		OrganizationID: &organizationIDDefault,

		timeout: timeout,
	}
}

// NewGetForDirectCustomersParamsWithContext creates a new GetForDirectCustomersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetForDirectCustomersParamsWithContext(ctx context.Context) *GetForDirectCustomersParams {
	var (
		countDefault          = int32(6)
		organizationIdDefault = int32(0)
	)
	return &GetForDirectCustomersParams{
		Count:          &countDefault,
		OrganizationID: &organizationIdDefault,

		Context: ctx,
	}
}

// NewGetForDirectCustomersParamsWithHTTPClient creates a new GetForDirectCustomersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetForDirectCustomersParamsWithHTTPClient(client *http.Client) *GetForDirectCustomersParams {
	var (
		countDefault          = int32(6)
		organizationIdDefault = int32(0)
	)
	return &GetForDirectCustomersParams{
		Count:          &countDefault,
		OrganizationID: &organizationIdDefault,
		HTTPClient:     client,
	}
}

/*GetForDirectCustomersParams contains all the parameters to send to the API endpoint
for the get for direct customers operation typically these are written to a http.Request
*/
type GetForDirectCustomersParams struct {

	/*Count*/
	Count *int32
	/*OrganizationID*/
	OrganizationID *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get for direct customers params
func (o *GetForDirectCustomersParams) WithTimeout(timeout time.Duration) *GetForDirectCustomersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get for direct customers params
func (o *GetForDirectCustomersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get for direct customers params
func (o *GetForDirectCustomersParams) WithContext(ctx context.Context) *GetForDirectCustomersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get for direct customers params
func (o *GetForDirectCustomersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get for direct customers params
func (o *GetForDirectCustomersParams) WithHTTPClient(client *http.Client) *GetForDirectCustomersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get for direct customers params
func (o *GetForDirectCustomersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get for direct customers params
func (o *GetForDirectCustomersParams) WithCount(count *int32) *GetForDirectCustomersParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get for direct customers params
func (o *GetForDirectCustomersParams) SetCount(count *int32) {
	o.Count = count
}

// WithOrganizationID adds the organizationID to the get for direct customers params
func (o *GetForDirectCustomersParams) WithOrganizationID(organizationID *int32) *GetForDirectCustomersParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get for direct customers params
func (o *GetForDirectCustomersParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetForDirectCustomersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int32
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

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
