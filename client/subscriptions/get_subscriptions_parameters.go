// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

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

// NewGetSubscriptionsParams creates a new GetSubscriptionsParams object
// with the default values initialized.
func NewGetSubscriptionsParams() *GetSubscriptionsParams {
	var ()
	return &GetSubscriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubscriptionsParamsWithTimeout creates a new GetSubscriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubscriptionsParamsWithTimeout(timeout time.Duration) *GetSubscriptionsParams {
	var ()
	return &GetSubscriptionsParams{

		timeout: timeout,
	}
}

// NewGetSubscriptionsParamsWithContext creates a new GetSubscriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubscriptionsParamsWithContext(ctx context.Context) *GetSubscriptionsParams {
	var ()
	return &GetSubscriptionsParams{

		Context: ctx,
	}
}

// NewGetSubscriptionsParamsWithHTTPClient creates a new GetSubscriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubscriptionsParamsWithHTTPClient(client *http.Client) *GetSubscriptionsParams {
	var ()
	return &GetSubscriptionsParams{
		HTTPClient: client,
	}
}

/*GetSubscriptionsParams contains all the parameters to send to the API endpoint
for the get subscriptions operation typically these are written to a http.Request
*/
type GetSubscriptionsParams struct {

	/*CustomerTenantID*/
	CustomerTenantID *int32
	/*IsTrial*/
	IsTrial *bool
	/*OrganizationID*/
	OrganizationID *int32
	/*Page*/
	Page *int32
	/*PageSize*/
	PageSize *int32
	/*PublisherID*/
	PublisherID *int32
	/*Refresh*/
	Refresh *bool
	/*RegisteredForReservedInstance*/
	RegisteredForReservedInstance *bool
	/*Search*/
	Search *string
	/*Statuses*/
	Statuses *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get subscriptions params
func (o *GetSubscriptionsParams) WithTimeout(timeout time.Duration) *GetSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subscriptions params
func (o *GetSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subscriptions params
func (o *GetSubscriptionsParams) WithContext(ctx context.Context) *GetSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subscriptions params
func (o *GetSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subscriptions params
func (o *GetSubscriptionsParams) WithHTTPClient(client *http.Client) *GetSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subscriptions params
func (o *GetSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerTenantID adds the customerTenantID to the get subscriptions params
func (o *GetSubscriptionsParams) WithCustomerTenantID(customerTenantID *int32) *GetSubscriptionsParams {
	o.SetCustomerTenantID(customerTenantID)
	return o
}

// SetCustomerTenantID adds the customerTenantId to the get subscriptions params
func (o *GetSubscriptionsParams) SetCustomerTenantID(customerTenantID *int32) {
	o.CustomerTenantID = customerTenantID
}

// WithIsTrial adds the isTrial to the get subscriptions params
func (o *GetSubscriptionsParams) WithIsTrial(isTrial *bool) *GetSubscriptionsParams {
	o.SetIsTrial(isTrial)
	return o
}

// SetIsTrial adds the isTrial to the get subscriptions params
func (o *GetSubscriptionsParams) SetIsTrial(isTrial *bool) {
	o.IsTrial = isTrial
}

// WithOrganizationID adds the organizationID to the get subscriptions params
func (o *GetSubscriptionsParams) WithOrganizationID(organizationID *int32) *GetSubscriptionsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get subscriptions params
func (o *GetSubscriptionsParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithPage adds the page to the get subscriptions params
func (o *GetSubscriptionsParams) WithPage(page *int32) *GetSubscriptionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get subscriptions params
func (o *GetSubscriptionsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get subscriptions params
func (o *GetSubscriptionsParams) WithPageSize(pageSize *int32) *GetSubscriptionsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get subscriptions params
func (o *GetSubscriptionsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPublisherID adds the publisherID to the get subscriptions params
func (o *GetSubscriptionsParams) WithPublisherID(publisherID *int32) *GetSubscriptionsParams {
	o.SetPublisherID(publisherID)
	return o
}

// SetPublisherID adds the publisherId to the get subscriptions params
func (o *GetSubscriptionsParams) SetPublisherID(publisherID *int32) {
	o.PublisherID = publisherID
}

// WithRefresh adds the refresh to the get subscriptions params
func (o *GetSubscriptionsParams) WithRefresh(refresh *bool) *GetSubscriptionsParams {
	o.SetRefresh(refresh)
	return o
}

// SetRefresh adds the refresh to the get subscriptions params
func (o *GetSubscriptionsParams) SetRefresh(refresh *bool) {
	o.Refresh = refresh
}

// WithRegisteredForReservedInstance adds the registeredForReservedInstance to the get subscriptions params
func (o *GetSubscriptionsParams) WithRegisteredForReservedInstance(registeredForReservedInstance *bool) *GetSubscriptionsParams {
	o.SetRegisteredForReservedInstance(registeredForReservedInstance)
	return o
}

// SetRegisteredForReservedInstance adds the registeredForReservedInstance to the get subscriptions params
func (o *GetSubscriptionsParams) SetRegisteredForReservedInstance(registeredForReservedInstance *bool) {
	o.RegisteredForReservedInstance = registeredForReservedInstance
}

// WithSearch adds the search to the get subscriptions params
func (o *GetSubscriptionsParams) WithSearch(search *string) *GetSubscriptionsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get subscriptions params
func (o *GetSubscriptionsParams) SetSearch(search *string) {
	o.Search = search
}

// WithStatuses adds the statuses to the get subscriptions params
func (o *GetSubscriptionsParams) WithStatuses(statuses *string) *GetSubscriptionsParams {
	o.SetStatuses(statuses)
	return o
}

// SetStatuses adds the statuses to the get subscriptions params
func (o *GetSubscriptionsParams) SetStatuses(statuses *string) {
	o.Statuses = statuses
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CustomerTenantID != nil {

		// query param CustomerTenantId
		var qrCustomerTenantID int32
		if o.CustomerTenantID != nil {
			qrCustomerTenantID = *o.CustomerTenantID
		}
		qCustomerTenantID := swag.FormatInt32(qrCustomerTenantID)
		if qCustomerTenantID != "" {
			if err := r.SetQueryParam("CustomerTenantId", qCustomerTenantID); err != nil {
				return err
			}
		}

	}

	if o.IsTrial != nil {

		// query param IsTrial
		var qrIsTrial bool
		if o.IsTrial != nil {
			qrIsTrial = *o.IsTrial
		}
		qIsTrial := swag.FormatBool(qrIsTrial)
		if qIsTrial != "" {
			if err := r.SetQueryParam("IsTrial", qIsTrial); err != nil {
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

	if o.Refresh != nil {

		// query param Refresh
		var qrRefresh bool
		if o.Refresh != nil {
			qrRefresh = *o.Refresh
		}
		qRefresh := swag.FormatBool(qrRefresh)
		if qRefresh != "" {
			if err := r.SetQueryParam("Refresh", qRefresh); err != nil {
				return err
			}
		}

	}

	if o.RegisteredForReservedInstance != nil {

		// query param RegisteredForReservedInstance
		var qrRegisteredForReservedInstance bool
		if o.RegisteredForReservedInstance != nil {
			qrRegisteredForReservedInstance = *o.RegisteredForReservedInstance
		}
		qRegisteredForReservedInstance := swag.FormatBool(qrRegisteredForReservedInstance)
		if qRegisteredForReservedInstance != "" {
			if err := r.SetQueryParam("RegisteredForReservedInstance", qRegisteredForReservedInstance); err != nil {
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

	if o.Statuses != nil {

		// query param Statuses
		var qrStatuses string
		if o.Statuses != nil {
			qrStatuses = *o.Statuses
		}
		qStatuses := qrStatuses
		if qStatuses != "" {
			if err := r.SetQueryParam("Statuses", qStatuses); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}