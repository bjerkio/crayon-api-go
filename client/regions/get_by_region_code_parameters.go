// Code generated by go-swagger; DO NOT EDIT.

package regions

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

// NewGetByRegionCodeParams creates a new GetByRegionCodeParams object
// with the default values initialized.
func NewGetByRegionCodeParams() *GetByRegionCodeParams {
	var ()
	return &GetByRegionCodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetByRegionCodeParamsWithTimeout creates a new GetByRegionCodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetByRegionCodeParamsWithTimeout(timeout time.Duration) *GetByRegionCodeParams {
	var ()
	return &GetByRegionCodeParams{

		timeout: timeout,
	}
}

// NewGetByRegionCodeParamsWithContext creates a new GetByRegionCodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetByRegionCodeParamsWithContext(ctx context.Context) *GetByRegionCodeParams {
	var ()
	return &GetByRegionCodeParams{

		Context: ctx,
	}
}

// NewGetByRegionCodeParamsWithHTTPClient creates a new GetByRegionCodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetByRegionCodeParamsWithHTTPClient(client *http.Client) *GetByRegionCodeParams {
	var ()
	return &GetByRegionCodeParams{
		HTTPClient: client,
	}
}

/*GetByRegionCodeParams contains all the parameters to send to the API endpoint
for the get by region code operation typically these are written to a http.Request
*/
type GetByRegionCodeParams struct {

	/*RegionCode*/
	RegionCode *string
	/*RegionList*/
	RegionList *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get by region code params
func (o *GetByRegionCodeParams) WithTimeout(timeout time.Duration) *GetByRegionCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get by region code params
func (o *GetByRegionCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get by region code params
func (o *GetByRegionCodeParams) WithContext(ctx context.Context) *GetByRegionCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get by region code params
func (o *GetByRegionCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get by region code params
func (o *GetByRegionCodeParams) WithHTTPClient(client *http.Client) *GetByRegionCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get by region code params
func (o *GetByRegionCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegionCode adds the regionCode to the get by region code params
func (o *GetByRegionCodeParams) WithRegionCode(regionCode *string) *GetByRegionCodeParams {
	o.SetRegionCode(regionCode)
	return o
}

// SetRegionCode adds the regionCode to the get by region code params
func (o *GetByRegionCodeParams) SetRegionCode(regionCode *string) {
	o.RegionCode = regionCode
}

// WithRegionList adds the regionList to the get by region code params
func (o *GetByRegionCodeParams) WithRegionList(regionList *string) *GetByRegionCodeParams {
	o.SetRegionList(regionList)
	return o
}

// SetRegionList adds the regionList to the get by region code params
func (o *GetByRegionCodeParams) SetRegionList(regionList *string) {
	o.RegionList = regionList
}

// WriteToRequest writes these params to a swagger request
func (o *GetByRegionCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RegionCode != nil {

		// query param regionCode
		var qrRegionCode string
		if o.RegionCode != nil {
			qrRegionCode = *o.RegionCode
		}
		qRegionCode := qrRegionCode
		if qRegionCode != "" {
			if err := r.SetQueryParam("regionCode", qRegionCode); err != nil {
				return err
			}
		}

	}

	if o.RegionList != nil {

		// query param regionList
		var qrRegionList string
		if o.RegionList != nil {
			qrRegionList = *o.RegionList
		}
		qRegionList := qrRegionList
		if qRegionList != "" {
			if err := r.SetQueryParam("regionList", qRegionList); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}