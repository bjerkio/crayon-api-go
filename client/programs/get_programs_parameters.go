// Code generated by go-swagger; DO NOT EDIT.

package programs

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

// NewGetProgramsParams creates a new GetProgramsParams object
// with the default values initialized.
func NewGetProgramsParams() *GetProgramsParams {
	var ()
	return &GetProgramsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProgramsParamsWithTimeout creates a new GetProgramsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProgramsParamsWithTimeout(timeout time.Duration) *GetProgramsParams {
	var ()
	return &GetProgramsParams{

		timeout: timeout,
	}
}

// NewGetProgramsParamsWithContext creates a new GetProgramsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProgramsParamsWithContext(ctx context.Context) *GetProgramsParams {
	var ()
	return &GetProgramsParams{

		Context: ctx,
	}
}

// NewGetProgramsParamsWithHTTPClient creates a new GetProgramsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProgramsParamsWithHTTPClient(client *http.Client) *GetProgramsParams {
	var ()
	return &GetProgramsParams{
		HTTPClient: client,
	}
}

/*GetProgramsParams contains all the parameters to send to the API endpoint
for the get programs operation typically these are written to a http.Request
*/
type GetProgramsParams struct {

	/*OrganizationID*/
	OrganizationID *int32
	/*Page*/
	Page *int32
	/*PageSize*/
	PageSize *int32
	/*ProgramType*/
	ProgramType *string
	/*PublisherID*/
	PublisherID *int32
	/*Search*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get programs params
func (o *GetProgramsParams) WithTimeout(timeout time.Duration) *GetProgramsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get programs params
func (o *GetProgramsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get programs params
func (o *GetProgramsParams) WithContext(ctx context.Context) *GetProgramsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get programs params
func (o *GetProgramsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get programs params
func (o *GetProgramsParams) WithHTTPClient(client *http.Client) *GetProgramsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get programs params
func (o *GetProgramsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get programs params
func (o *GetProgramsParams) WithOrganizationID(organizationID *int32) *GetProgramsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get programs params
func (o *GetProgramsParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithPage adds the page to the get programs params
func (o *GetProgramsParams) WithPage(page *int32) *GetProgramsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get programs params
func (o *GetProgramsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get programs params
func (o *GetProgramsParams) WithPageSize(pageSize *int32) *GetProgramsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get programs params
func (o *GetProgramsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithProgramType adds the programType to the get programs params
func (o *GetProgramsParams) WithProgramType(programType *string) *GetProgramsParams {
	o.SetProgramType(programType)
	return o
}

// SetProgramType adds the programType to the get programs params
func (o *GetProgramsParams) SetProgramType(programType *string) {
	o.ProgramType = programType
}

// WithPublisherID adds the publisherID to the get programs params
func (o *GetProgramsParams) WithPublisherID(publisherID *int32) *GetProgramsParams {
	o.SetPublisherID(publisherID)
	return o
}

// SetPublisherID adds the publisherId to the get programs params
func (o *GetProgramsParams) SetPublisherID(publisherID *int32) {
	o.PublisherID = publisherID
}

// WithSearch adds the search to the get programs params
func (o *GetProgramsParams) WithSearch(search *string) *GetProgramsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get programs params
func (o *GetProgramsParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetProgramsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ProgramType != nil {

		// query param ProgramType
		var qrProgramType string
		if o.ProgramType != nil {
			qrProgramType = *o.ProgramType
		}
		qProgramType := qrProgramType
		if qProgramType != "" {
			if err := r.SetQueryParam("ProgramType", qProgramType); err != nil {
				return err
			}
		}

	}

	if o.PublisherID != nil {

		// query param PublisherId
		var qrPublisherID int32
		if o.PublisherID != nil {
			qrPublisherID = *o.PublisherID
		}
		qPublisherID := swag.FormatInt32(qrPublisherID)
		if qPublisherID != "" {
			if err := r.SetQueryParam("PublisherId", qPublisherID); err != nil {
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
