// Code generated by go-swagger; DO NOT EDIT.

package azure_plans

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

// NewPostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams creates a new PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams object
// with the default values initialized.
func NewPostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams() *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams {
	var ()
	return &PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParamsWithTimeout creates a new PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParamsWithTimeout(timeout time.Duration) *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams {
	var ()
	return &PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams{

		timeout: timeout,
	}
}

// NewPostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParamsWithContext creates a new PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParamsWithContext(ctx context.Context) *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams {
	var ()
	return &PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams{

		Context: ctx,
	}
}

// NewPostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParamsWithHTTPClient creates a new PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParamsWithHTTPClient(client *http.Client) *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams {
	var ()
	return &PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams{
		HTTPClient: client,
	}
}

/*PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams contains all the parameters to send to the API endpoint
for the post API v1 azure plans azure plan ID azure subscriptions ID enable operation typically these are written to a http.Request
*/
type PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams struct {

	/*AzurePlanID*/
	AzurePlanID int32
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v1 azure plans azure plan ID azure subscriptions ID enable params
func (o *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams) WithTimeout(timeout time.Duration) *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v1 azure plans azure plan ID azure subscriptions ID enable params
func (o *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v1 azure plans azure plan ID azure subscriptions ID enable params
func (o *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams) WithContext(ctx context.Context) *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v1 azure plans azure plan ID azure subscriptions ID enable params
func (o *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v1 azure plans azure plan ID azure subscriptions ID enable params
func (o *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams) WithHTTPClient(client *http.Client) *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v1 azure plans azure plan ID azure subscriptions ID enable params
func (o *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAzurePlanID adds the azurePlanID to the post API v1 azure plans azure plan ID azure subscriptions ID enable params
func (o *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams) WithAzurePlanID(azurePlanID int32) *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams {
	o.SetAzurePlanID(azurePlanID)
	return o
}

// SetAzurePlanID adds the azurePlanId to the post API v1 azure plans azure plan ID azure subscriptions ID enable params
func (o *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams) SetAzurePlanID(azurePlanID int32) {
	o.AzurePlanID = azurePlanID
}

// WithID adds the id to the post API v1 azure plans azure plan ID azure subscriptions ID enable params
func (o *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams) WithID(id int32) *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post API v1 azure plans azure plan ID azure subscriptions ID enable params
func (o *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV1AzurePlansAzurePlanIDAzureSubscriptionsIDEnableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param azurePlanId
	if err := r.SetPathParam("azurePlanId", swag.FormatInt32(o.AzurePlanID)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
