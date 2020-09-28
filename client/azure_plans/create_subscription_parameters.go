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

// NewCreateSubscriptionParams creates a new CreateSubscriptionParams object
// with the default values initialized.
func NewCreateSubscriptionParams() *CreateSubscriptionParams {
	var ()
	return &CreateSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSubscriptionParamsWithTimeout creates a new CreateSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSubscriptionParamsWithTimeout(timeout time.Duration) *CreateSubscriptionParams {
	var ()
	return &CreateSubscriptionParams{

		timeout: timeout,
	}
}

// NewCreateSubscriptionParamsWithContext creates a new CreateSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSubscriptionParamsWithContext(ctx context.Context) *CreateSubscriptionParams {
	var ()
	return &CreateSubscriptionParams{

		Context: ctx,
	}
}

// NewCreateSubscriptionParamsWithHTTPClient creates a new CreateSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSubscriptionParamsWithHTTPClient(client *http.Client) *CreateSubscriptionParams {
	var ()
	return &CreateSubscriptionParams{
		HTTPClient: client,
	}
}

/*CreateSubscriptionParams contains all the parameters to send to the API endpoint
for the create subscription operation typically these are written to a http.Request
*/
type CreateSubscriptionParams struct {

	/*AzurePlanID*/
	AzurePlanID int32
	/*RequestBody*/
	RequestBody *models.CreateAzureSubscriptionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create subscription params
func (o *CreateSubscriptionParams) WithTimeout(timeout time.Duration) *CreateSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create subscription params
func (o *CreateSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create subscription params
func (o *CreateSubscriptionParams) WithContext(ctx context.Context) *CreateSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create subscription params
func (o *CreateSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create subscription params
func (o *CreateSubscriptionParams) WithHTTPClient(client *http.Client) *CreateSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create subscription params
func (o *CreateSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAzurePlanID adds the azurePlanID to the create subscription params
func (o *CreateSubscriptionParams) WithAzurePlanID(azurePlanID int32) *CreateSubscriptionParams {
	o.SetAzurePlanID(azurePlanID)
	return o
}

// SetAzurePlanID adds the azurePlanId to the create subscription params
func (o *CreateSubscriptionParams) SetAzurePlanID(azurePlanID int32) {
	o.AzurePlanID = azurePlanID
}

// WithRequestBody adds the requestBody to the create subscription params
func (o *CreateSubscriptionParams) WithRequestBody(requestBody *models.CreateAzureSubscriptionRequest) *CreateSubscriptionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create subscription params
func (o *CreateSubscriptionParams) SetRequestBody(requestBody *models.CreateAzureSubscriptionRequest) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param azurePlanId
	if err := r.SetPathParam("azurePlanId", swag.FormatInt32(o.AzurePlanID)); err != nil {
		return err
	}

	if o.RequestBody != nil {
		if err := r.SetBodyParam(o.RequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
