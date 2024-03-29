// Code generated by go-swagger; DO NOT EDIT.

package management_links

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

// NewGetGroupedParams creates a new GetGroupedParams object
// with the default values initialized.
func NewGetGroupedParams() *GetGroupedParams {
	var ()
	return &GetGroupedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupedParamsWithTimeout creates a new GetGroupedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupedParamsWithTimeout(timeout time.Duration) *GetGroupedParams {
	var ()
	return &GetGroupedParams{

		timeout: timeout,
	}
}

// NewGetGroupedParamsWithContext creates a new GetGroupedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupedParamsWithContext(ctx context.Context) *GetGroupedParams {
	var ()
	return &GetGroupedParams{

		Context: ctx,
	}
}

// NewGetGroupedParamsWithHTTPClient creates a new GetGroupedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupedParamsWithHTTPClient(client *http.Client) *GetGroupedParams {
	var ()
	return &GetGroupedParams{
		HTTPClient: client,
	}
}

/*GetGroupedParams contains all the parameters to send to the API endpoint
for the get grouped operation typically these are written to a http.Request
*/
type GetGroupedParams struct {

	/*Page*/
	Page *int32
	/*PageSize*/
	PageSize *int32
	/*ResellerCustomerIds*/
	ResellerCustomerIds []int32
	/*SubscriptionIds*/
	SubscriptionIds []int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get grouped params
func (o *GetGroupedParams) WithTimeout(timeout time.Duration) *GetGroupedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get grouped params
func (o *GetGroupedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get grouped params
func (o *GetGroupedParams) WithContext(ctx context.Context) *GetGroupedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get grouped params
func (o *GetGroupedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get grouped params
func (o *GetGroupedParams) WithHTTPClient(client *http.Client) *GetGroupedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get grouped params
func (o *GetGroupedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get grouped params
func (o *GetGroupedParams) WithPage(page *int32) *GetGroupedParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get grouped params
func (o *GetGroupedParams) SetPage(page *int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get grouped params
func (o *GetGroupedParams) WithPageSize(pageSize *int32) *GetGroupedParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get grouped params
func (o *GetGroupedParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithResellerCustomerIds adds the resellerCustomerIds to the get grouped params
func (o *GetGroupedParams) WithResellerCustomerIds(resellerCustomerIds []int32) *GetGroupedParams {
	o.SetResellerCustomerIds(resellerCustomerIds)
	return o
}

// SetResellerCustomerIds adds the resellerCustomerIds to the get grouped params
func (o *GetGroupedParams) SetResellerCustomerIds(resellerCustomerIds []int32) {
	o.ResellerCustomerIds = resellerCustomerIds
}

// WithSubscriptionIds adds the subscriptionIds to the get grouped params
func (o *GetGroupedParams) WithSubscriptionIds(subscriptionIds []int32) *GetGroupedParams {
	o.SetSubscriptionIds(subscriptionIds)
	return o
}

// SetSubscriptionIds adds the subscriptionIds to the get grouped params
func (o *GetGroupedParams) SetSubscriptionIds(subscriptionIds []int32) {
	o.SubscriptionIds = subscriptionIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	var valuesResellerCustomerIds []string
	for _, v := range o.ResellerCustomerIds {
		valuesResellerCustomerIds = append(valuesResellerCustomerIds, swag.FormatInt32(v))
	}

	joinedResellerCustomerIds := swag.JoinByFormat(valuesResellerCustomerIds, "multi")
	// query array param ResellerCustomerIds
	if err := r.SetQueryParam("ResellerCustomerIds", joinedResellerCustomerIds...); err != nil {
		return err
	}

	var valuesSubscriptionIds []string
	for _, v := range o.SubscriptionIds {
		valuesSubscriptionIds = append(valuesSubscriptionIds, swag.FormatInt32(v))
	}

	joinedSubscriptionIds := swag.JoinByFormat(valuesSubscriptionIds, "multi")
	// query array param SubscriptionIds
	if err := r.SetQueryParam("SubscriptionIds", joinedSubscriptionIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
