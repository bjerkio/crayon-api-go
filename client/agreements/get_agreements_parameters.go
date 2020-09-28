// Code generated by go-swagger; DO NOT EDIT.

package agreements

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

// NewGetAgreementsParams creates a new GetAgreementsParams object
// with the default values initialized.
func NewGetAgreementsParams() *GetAgreementsParams {
	var ()
	return &GetAgreementsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgreementsParamsWithTimeout creates a new GetAgreementsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAgreementsParamsWithTimeout(timeout time.Duration) *GetAgreementsParams {
	var ()
	return &GetAgreementsParams{

		timeout: timeout,
	}
}

// NewGetAgreementsParamsWithContext creates a new GetAgreementsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAgreementsParamsWithContext(ctx context.Context) *GetAgreementsParams {
	var ()
	return &GetAgreementsParams{

		Context: ctx,
	}
}

// NewGetAgreementsParamsWithHTTPClient creates a new GetAgreementsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAgreementsParamsWithHTTPClient(client *http.Client) *GetAgreementsParams {
	var ()
	return &GetAgreementsParams{
		HTTPClient: client,
	}
}

/*GetAgreementsParams contains all the parameters to send to the API endpoint
for the get agreements operation typically these are written to a http.Request
*/
type GetAgreementsParams struct {

	/*AgreementIds*/
	AgreementIds []int32
	/*AgreementTypes*/
	AgreementTypes []string
	/*EndDateFrom*/
	EndDateFrom *strfmt.DateTime
	/*EndDateTo*/
	EndDateTo *strfmt.DateTime
	/*OrganizationID*/
	OrganizationID *int32
	/*OrganizationIds*/
	OrganizationIds []int32
	/*Page*/
	Page *int32
	/*PageSize*/
	PageSize *int32
	/*PricelistIds*/
	PricelistIds []int32
	/*ProgramIds*/
	ProgramIds []int32
	/*PublisherID*/
	PublisherID *int32
	/*PublisherIds*/
	PublisherIds []int32
	/*SalesPriceCurrency*/
	SalesPriceCurrency *string
	/*Search*/
	Search *string
	/*SearchDate*/
	SearchDate *strfmt.DateTime
	/*Status*/
	Status *string
	/*TermRequired*/
	TermRequired *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get agreements params
func (o *GetAgreementsParams) WithTimeout(timeout time.Duration) *GetAgreementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agreements params
func (o *GetAgreementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agreements params
func (o *GetAgreementsParams) WithContext(ctx context.Context) *GetAgreementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agreements params
func (o *GetAgreementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agreements params
func (o *GetAgreementsParams) WithHTTPClient(client *http.Client) *GetAgreementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agreements params
func (o *GetAgreementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgreementIds adds the agreementIds to the get agreements params
func (o *GetAgreementsParams) WithAgreementIds(agreementIds []int32) *GetAgreementsParams {
	o.SetAgreementIds(agreementIds)
	return o
}

// SetAgreementIds adds the agreementIds to the get agreements params
func (o *GetAgreementsParams) SetAgreementIds(agreementIds []int32) {
	o.AgreementIds = agreementIds
}

// WithAgreementTypes adds the agreementTypes to the get agreements params
func (o *GetAgreementsParams) WithAgreementTypes(agreementTypes []string) *GetAgreementsParams {
	o.SetAgreementTypes(agreementTypes)
	return o
}

// SetAgreementTypes adds the agreementTypes to the get agreements params
func (o *GetAgreementsParams) SetAgreementTypes(agreementTypes []string) {
	o.AgreementTypes = agreementTypes
}

// WithEndDateFrom adds the endDateFrom to the get agreements params
func (o *GetAgreementsParams) WithEndDateFrom(endDateFrom *strfmt.DateTime) *GetAgreementsParams {
	o.SetEndDateFrom(endDateFrom)
	return o
}

// SetEndDateFrom adds the endDateFrom to the get agreements params
func (o *GetAgreementsParams) SetEndDateFrom(endDateFrom *strfmt.DateTime) {
	o.EndDateFrom = endDateFrom
}

// WithEndDateTo adds the endDateTo to the get agreements params
func (o *GetAgreementsParams) WithEndDateTo(endDateTo *strfmt.DateTime) *GetAgreementsParams {
	o.SetEndDateTo(endDateTo)
	return o
}

// SetEndDateTo adds the endDateTo to the get agreements params
func (o *GetAgreementsParams) SetEndDateTo(endDateTo *strfmt.DateTime) {
	o.EndDateTo = endDateTo
}

// WithOrganizationID adds the organizationID to the get agreements params
func (o *GetAgreementsParams) WithOrganizationID(organizationID *int32) *GetAgreementsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get agreements params
func (o *GetAgreementsParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithOrganizationIds adds the organizationIds to the get agreements params
func (o *GetAgreementsParams) WithOrganizationIds(organizationIds []int32) *GetAgreementsParams {
	o.SetOrganizationIds(organizationIds)
	return o
}

// SetOrganizationIds adds the organizationIds to the get agreements params
func (o *GetAgreementsParams) SetOrganizationIds(organizationIds []int32) {
	o.OrganizationIds = organizationIds
}

// WithPage adds the page to the get agreements params
func (o *GetAgreementsParams) WithPage(page *int32) *GetAgreementsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get agreements params
func (o *GetAgreementsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get agreements params
func (o *GetAgreementsParams) WithPageSize(pageSize *int32) *GetAgreementsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get agreements params
func (o *GetAgreementsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPricelistIds adds the pricelistIds to the get agreements params
func (o *GetAgreementsParams) WithPricelistIds(pricelistIds []int32) *GetAgreementsParams {
	o.SetPricelistIds(pricelistIds)
	return o
}

// SetPricelistIds adds the pricelistIds to the get agreements params
func (o *GetAgreementsParams) SetPricelistIds(pricelistIds []int32) {
	o.PricelistIds = pricelistIds
}

// WithProgramIds adds the programIds to the get agreements params
func (o *GetAgreementsParams) WithProgramIds(programIds []int32) *GetAgreementsParams {
	o.SetProgramIds(programIds)
	return o
}

// SetProgramIds adds the programIds to the get agreements params
func (o *GetAgreementsParams) SetProgramIds(programIds []int32) {
	o.ProgramIds = programIds
}

// WithPublisherID adds the publisherID to the get agreements params
func (o *GetAgreementsParams) WithPublisherID(publisherID *int32) *GetAgreementsParams {
	o.SetPublisherID(publisherID)
	return o
}

// SetPublisherID adds the publisherId to the get agreements params
func (o *GetAgreementsParams) SetPublisherID(publisherID *int32) {
	o.PublisherID = publisherID
}

// WithPublisherIds adds the publisherIds to the get agreements params
func (o *GetAgreementsParams) WithPublisherIds(publisherIds []int32) *GetAgreementsParams {
	o.SetPublisherIds(publisherIds)
	return o
}

// SetPublisherIds adds the publisherIds to the get agreements params
func (o *GetAgreementsParams) SetPublisherIds(publisherIds []int32) {
	o.PublisherIds = publisherIds
}

// WithSalesPriceCurrency adds the salesPriceCurrency to the get agreements params
func (o *GetAgreementsParams) WithSalesPriceCurrency(salesPriceCurrency *string) *GetAgreementsParams {
	o.SetSalesPriceCurrency(salesPriceCurrency)
	return o
}

// SetSalesPriceCurrency adds the salesPriceCurrency to the get agreements params
func (o *GetAgreementsParams) SetSalesPriceCurrency(salesPriceCurrency *string) {
	o.SalesPriceCurrency = salesPriceCurrency
}

// WithSearch adds the search to the get agreements params
func (o *GetAgreementsParams) WithSearch(search *string) *GetAgreementsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get agreements params
func (o *GetAgreementsParams) SetSearch(search *string) {
	o.Search = search
}

// WithSearchDate adds the searchDate to the get agreements params
func (o *GetAgreementsParams) WithSearchDate(searchDate *strfmt.DateTime) *GetAgreementsParams {
	o.SetSearchDate(searchDate)
	return o
}

// SetSearchDate adds the searchDate to the get agreements params
func (o *GetAgreementsParams) SetSearchDate(searchDate *strfmt.DateTime) {
	o.SearchDate = searchDate
}

// WithStatus adds the status to the get agreements params
func (o *GetAgreementsParams) WithStatus(status *string) *GetAgreementsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get agreements params
func (o *GetAgreementsParams) SetStatus(status *string) {
	o.Status = status
}

// WithTermRequired adds the termRequired to the get agreements params
func (o *GetAgreementsParams) WithTermRequired(termRequired *bool) *GetAgreementsParams {
	o.SetTermRequired(termRequired)
	return o
}

// SetTermRequired adds the termRequired to the get agreements params
func (o *GetAgreementsParams) SetTermRequired(termRequired *bool) {
	o.TermRequired = termRequired
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgreementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	var valuesAgreementIds []string
	for _, v := range o.AgreementIds {
		valuesAgreementIds = append(valuesAgreementIds, swag.FormatInt32(v))
	}

	joinedAgreementIds := swag.JoinByFormat(valuesAgreementIds, "multi")
	// query array param AgreementIds
	if err := r.SetQueryParam("AgreementIds", joinedAgreementIds...); err != nil {
		return err
	}

	valuesAgreementTypes := o.AgreementTypes

	joinedAgreementTypes := swag.JoinByFormat(valuesAgreementTypes, "multi")
	// query array param AgreementTypes
	if err := r.SetQueryParam("AgreementTypes", joinedAgreementTypes...); err != nil {
		return err
	}

	if o.EndDateFrom != nil {

		// query param EndDateFrom
		var qrEndDateFrom strfmt.DateTime
		if o.EndDateFrom != nil {
			qrEndDateFrom = *o.EndDateFrom
		}
		qEndDateFrom := qrEndDateFrom.String()
		if qEndDateFrom != "" {
			if err := r.SetQueryParam("EndDateFrom", qEndDateFrom); err != nil {
				return err
			}
		}

	}

	if o.EndDateTo != nil {

		// query param EndDateTo
		var qrEndDateTo strfmt.DateTime
		if o.EndDateTo != nil {
			qrEndDateTo = *o.EndDateTo
		}
		qEndDateTo := qrEndDateTo.String()
		if qEndDateTo != "" {
			if err := r.SetQueryParam("EndDateTo", qEndDateTo); err != nil {
				return err
			}
		}

	}

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

	var valuesOrganizationIds []string
	for _, v := range o.OrganizationIds {
		valuesOrganizationIds = append(valuesOrganizationIds, swag.FormatInt32(v))
	}

	joinedOrganizationIds := swag.JoinByFormat(valuesOrganizationIds, "multi")
	// query array param OrganizationIds
	if err := r.SetQueryParam("OrganizationIds", joinedOrganizationIds...); err != nil {
		return err
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

	var valuesPricelistIds []string
	for _, v := range o.PricelistIds {
		valuesPricelistIds = append(valuesPricelistIds, swag.FormatInt32(v))
	}

	joinedPricelistIds := swag.JoinByFormat(valuesPricelistIds, "multi")
	// query array param PricelistIds
	if err := r.SetQueryParam("PricelistIds", joinedPricelistIds...); err != nil {
		return err
	}

	var valuesProgramIds []string
	for _, v := range o.ProgramIds {
		valuesProgramIds = append(valuesProgramIds, swag.FormatInt32(v))
	}

	joinedProgramIds := swag.JoinByFormat(valuesProgramIds, "multi")
	// query array param ProgramIds
	if err := r.SetQueryParam("ProgramIds", joinedProgramIds...); err != nil {
		return err
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

	var valuesPublisherIds []string
	for _, v := range o.PublisherIds {
		valuesPublisherIds = append(valuesPublisherIds, swag.FormatInt32(v))
	}

	joinedPublisherIds := swag.JoinByFormat(valuesPublisherIds, "multi")
	// query array param PublisherIds
	if err := r.SetQueryParam("PublisherIds", joinedPublisherIds...); err != nil {
		return err
	}

	if o.SalesPriceCurrency != nil {

		// query param SalesPriceCurrency
		var qrSalesPriceCurrency string
		if o.SalesPriceCurrency != nil {
			qrSalesPriceCurrency = *o.SalesPriceCurrency
		}
		qSalesPriceCurrency := qrSalesPriceCurrency
		if qSalesPriceCurrency != "" {
			if err := r.SetQueryParam("SalesPriceCurrency", qSalesPriceCurrency); err != nil {
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

	if o.SearchDate != nil {

		// query param SearchDate
		var qrSearchDate strfmt.DateTime
		if o.SearchDate != nil {
			qrSearchDate = *o.SearchDate
		}
		qSearchDate := qrSearchDate.String()
		if qSearchDate != "" {
			if err := r.SetQueryParam("SearchDate", qSearchDate); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param Status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("Status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.TermRequired != nil {

		// query param TermRequired
		var qrTermRequired bool
		if o.TermRequired != nil {
			qrTermRequired = *o.TermRequired
		}
		qTermRequired := swag.FormatBool(qrTermRequired)
		if qTermRequired != "" {
			if err := r.SetQueryParam("TermRequired", qTermRequired); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}