// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewGetOrganizationsParams creates a new GetOrganizationsParams object
// with the default values initialized.
func NewGetOrganizationsParams() *GetOrganizationsParams {
	var ()
	return &GetOrganizationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsParamsWithTimeout creates a new GetOrganizationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationsParamsWithTimeout(timeout time.Duration) *GetOrganizationsParams {
	var ()
	return &GetOrganizationsParams{

		timeout: timeout,
	}
}

// NewGetOrganizationsParamsWithContext creates a new GetOrganizationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationsParamsWithContext(ctx context.Context) *GetOrganizationsParams {
	var ()
	return &GetOrganizationsParams{

		Context: ctx,
	}
}

// NewGetOrganizationsParamsWithHTTPClient creates a new GetOrganizationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationsParamsWithHTTPClient(client *http.Client) *GetOrganizationsParams {
	var ()
	return &GetOrganizationsParams{
		HTTPClient: client,
	}
}

/*GetOrganizationsParams contains all the parameters to send to the API endpoint
for the get organizations operation typically these are written to a http.Request
*/
type GetOrganizationsParams struct {

	/*Page*/
	Page *int32
	/*PageSize*/
	PageSize *int32
	/*Search*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organizations params
func (o *GetOrganizationsParams) WithTimeout(timeout time.Duration) *GetOrganizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations params
func (o *GetOrganizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations params
func (o *GetOrganizationsParams) WithContext(ctx context.Context) *GetOrganizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations params
func (o *GetOrganizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations params
func (o *GetOrganizationsParams) WithHTTPClient(client *http.Client) *GetOrganizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations params
func (o *GetOrganizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get organizations params
func (o *GetOrganizationsParams) WithPage(page *int32) *GetOrganizationsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get organizations params
func (o *GetOrganizationsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get organizations params
func (o *GetOrganizationsParams) WithPageSize(pageSize *int32) *GetOrganizationsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get organizations params
func (o *GetOrganizationsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the get organizations params
func (o *GetOrganizationsParams) WithSearch(search *string) *GetOrganizationsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get organizations params
func (o *GetOrganizationsParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param Page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("Page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param PageSize
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("PageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.Search != nil {

		// query param Search
		var qrSearch string
		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {
			if err := r.SetQueryParam("Search", qSearch); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
