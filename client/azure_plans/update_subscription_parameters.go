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

	"github.com/bjerkio/crayon-api-go/models"
)

// NewUpdateSubscriptionParams creates a new UpdateSubscriptionParams object
// with the default values initialized.
func NewUpdateSubscriptionParams() *UpdateSubscriptionParams {
	var ()
	return &UpdateSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSubscriptionParamsWithTimeout creates a new UpdateSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSubscriptionParamsWithTimeout(timeout time.Duration) *UpdateSubscriptionParams {
	var ()
	return &UpdateSubscriptionParams{

		timeout: timeout,
	}
}

// NewUpdateSubscriptionParamsWithContext creates a new UpdateSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSubscriptionParamsWithContext(ctx context.Context) *UpdateSubscriptionParams {
	var ()
	return &UpdateSubscriptionParams{

		Context: ctx,
	}
}

// NewUpdateSubscriptionParamsWithHTTPClient creates a new UpdateSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSubscriptionParamsWithHTTPClient(client *http.Client) *UpdateSubscriptionParams {
	var ()
	return &UpdateSubscriptionParams{
		HTTPClient: client,
	}
}

/*UpdateSubscriptionParams contains all the parameters to send to the API endpoint
for the update subscription operation typically these are written to a http.Request
*/
type UpdateSubscriptionParams struct {

	/*AzurePlanID*/
	AzurePlanID int32
	/*AzureSubscription*/
	AzureSubscription *models.PutAzureSubscription
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update subscription params
func (o *UpdateSubscriptionParams) WithTimeout(timeout time.Duration) *UpdateSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update subscription params
func (o *UpdateSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update subscription params
func (o *UpdateSubscriptionParams) WithContext(ctx context.Context) *UpdateSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update subscription params
func (o *UpdateSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update subscription params
func (o *UpdateSubscriptionParams) WithHTTPClient(client *http.Client) *UpdateSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update subscription params
func (o *UpdateSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAzurePlanID adds the azurePlanID to the update subscription params
func (o *UpdateSubscriptionParams) WithAzurePlanID(azurePlanID int32) *UpdateSubscriptionParams {
	o.SetAzurePlanID(azurePlanID)
	return o
}

// SetAzurePlanID adds the azurePlanId to the update subscription params
func (o *UpdateSubscriptionParams) SetAzurePlanID(azurePlanID int32) {
	o.AzurePlanID = azurePlanID
}

// WithAzureSubscription adds the azureSubscription to the update subscription params
func (o *UpdateSubscriptionParams) WithAzureSubscription(azureSubscription *models.PutAzureSubscription) *UpdateSubscriptionParams {
	o.SetAzureSubscription(azureSubscription)
	return o
}

// SetAzureSubscription adds the azureSubscription to the update subscription params
func (o *UpdateSubscriptionParams) SetAzureSubscription(azureSubscription *models.PutAzureSubscription) {
	o.AzureSubscription = azureSubscription
}

// WithID adds the id to the update subscription params
func (o *UpdateSubscriptionParams) WithID(id int32) *UpdateSubscriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update subscription params
func (o *UpdateSubscriptionParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param azurePlanId
	if err := r.SetPathParam("azurePlanId", swag.FormatInt32(o.AzurePlanID)); err != nil {
		return err
	}

	if o.AzureSubscription != nil {
		if err := r.SetBodyParam(o.AzureSubscription); err != nil {
			return err
		}
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
