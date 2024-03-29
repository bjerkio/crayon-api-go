// Code generated by go-swagger; DO NOT EDIT.

package invoice_profiles

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

// NewGetInvoiceProfilesParams creates a new GetInvoiceProfilesParams object
// with the default values initialized.
func NewGetInvoiceProfilesParams() *GetInvoiceProfilesParams {
	var ()
	return &GetInvoiceProfilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceProfilesParamsWithTimeout creates a new GetInvoiceProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoiceProfilesParamsWithTimeout(timeout time.Duration) *GetInvoiceProfilesParams {
	var ()
	return &GetInvoiceProfilesParams{

		timeout: timeout,
	}
}

// NewGetInvoiceProfilesParamsWithContext creates a new GetInvoiceProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoiceProfilesParamsWithContext(ctx context.Context) *GetInvoiceProfilesParams {
	var ()
	return &GetInvoiceProfilesParams{

		Context: ctx,
	}
}

// NewGetInvoiceProfilesParamsWithHTTPClient creates a new GetInvoiceProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoiceProfilesParamsWithHTTPClient(client *http.Client) *GetInvoiceProfilesParams {
	var ()
	return &GetInvoiceProfilesParams{
		HTTPClient: client,
	}
}

/*GetInvoiceProfilesParams contains all the parameters to send to the API endpoint
for the get invoice profiles operation typically these are written to a http.Request
*/
type GetInvoiceProfilesParams struct {

	/*OrganizationID*/
	OrganizationID *int32
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

// WithTimeout adds the timeout to the get invoice profiles params
func (o *GetInvoiceProfilesParams) WithTimeout(timeout time.Duration) *GetInvoiceProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice profiles params
func (o *GetInvoiceProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice profiles params
func (o *GetInvoiceProfilesParams) WithContext(ctx context.Context) *GetInvoiceProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice profiles params
func (o *GetInvoiceProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice profiles params
func (o *GetInvoiceProfilesParams) WithHTTPClient(client *http.Client) *GetInvoiceProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice profiles params
func (o *GetInvoiceProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get invoice profiles params
func (o *GetInvoiceProfilesParams) WithOrganizationID(organizationID *int32) *GetInvoiceProfilesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get invoice profiles params
func (o *GetInvoiceProfilesParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithPage adds the page to the get invoice profiles params
func (o *GetInvoiceProfilesParams) WithPage(page *int32) *GetInvoiceProfilesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get invoice profiles params
func (o *GetInvoiceProfilesParams) SetPage(page *int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get invoice profiles params
func (o *GetInvoiceProfilesParams) WithPageSize(pageSize *int32) *GetInvoiceProfilesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get invoice profiles params
func (o *GetInvoiceProfilesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the get invoice profiles params
func (o *GetInvoiceProfilesParams) WithSearch(search *string) *GetInvoiceProfilesParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get invoice profiles params
func (o *GetInvoiceProfilesParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrganizationID != nil {

		// query param OrganizationId
		var qrOrganizationID int32
		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := swag.FormatInt32(qrOrganizationID)
		if qOrganizationID != "" {
			if err := r.SetQueryParam("OrganizationId", qOrganizationID); err != nil {
				return err
			}
		}

	}

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
