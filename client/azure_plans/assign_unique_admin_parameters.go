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

// NewAssignUniqueAdminParams creates a new AssignUniqueAdminParams object
// with the default values initialized.
func NewAssignUniqueAdminParams() *AssignUniqueAdminParams {
	var ()
	return &AssignUniqueAdminParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAssignUniqueAdminParamsWithTimeout creates a new AssignUniqueAdminParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAssignUniqueAdminParamsWithTimeout(timeout time.Duration) *AssignUniqueAdminParams {
	var ()
	return &AssignUniqueAdminParams{

		timeout: timeout,
	}
}

// NewAssignUniqueAdminParamsWithContext creates a new AssignUniqueAdminParams object
// with the default values initialized, and the ability to set a context for a request
func NewAssignUniqueAdminParamsWithContext(ctx context.Context) *AssignUniqueAdminParams {
	var ()
	return &AssignUniqueAdminParams{

		Context: ctx,
	}
}

// NewAssignUniqueAdminParamsWithHTTPClient creates a new AssignUniqueAdminParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAssignUniqueAdminParamsWithHTTPClient(client *http.Client) *AssignUniqueAdminParams {
	var ()
	return &AssignUniqueAdminParams{
		HTTPClient: client,
	}
}

/*AssignUniqueAdminParams contains all the parameters to send to the API endpoint
for the assign unique admin operation typically these are written to a http.Request
*/
type AssignUniqueAdminParams struct {

	/*AzurePlanID*/
	AzurePlanID int32
	/*Body*/
	Body *models.AzureSubscriptionAssignAdmin
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the assign unique admin params
func (o *AssignUniqueAdminParams) WithTimeout(timeout time.Duration) *AssignUniqueAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign unique admin params
func (o *AssignUniqueAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign unique admin params
func (o *AssignUniqueAdminParams) WithContext(ctx context.Context) *AssignUniqueAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign unique admin params
func (o *AssignUniqueAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign unique admin params
func (o *AssignUniqueAdminParams) WithHTTPClient(client *http.Client) *AssignUniqueAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign unique admin params
func (o *AssignUniqueAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAzurePlanID adds the azurePlanID to the assign unique admin params
func (o *AssignUniqueAdminParams) WithAzurePlanID(azurePlanID int32) *AssignUniqueAdminParams {
	o.SetAzurePlanID(azurePlanID)
	return o
}

// SetAzurePlanID adds the azurePlanId to the assign unique admin params
func (o *AssignUniqueAdminParams) SetAzurePlanID(azurePlanID int32) {
	o.AzurePlanID = azurePlanID
}

// WithBody adds the body to the assign unique admin params
func (o *AssignUniqueAdminParams) WithBody(body *models.AzureSubscriptionAssignAdmin) *AssignUniqueAdminParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the assign unique admin params
func (o *AssignUniqueAdminParams) SetBody(body *models.AzureSubscriptionAssignAdmin) {
	o.Body = body
}

// WithID adds the id to the assign unique admin params
func (o *AssignUniqueAdminParams) WithID(id int32) *AssignUniqueAdminParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the assign unique admin params
func (o *AssignUniqueAdminParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AssignUniqueAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param azurePlanId
	if err := r.SetPathParam("azurePlanId", swag.FormatInt32(o.AzurePlanID)); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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