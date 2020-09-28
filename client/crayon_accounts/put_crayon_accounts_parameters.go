// Code generated by go-swagger; DO NOT EDIT.

package crayon_accounts

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

// NewPutCrayonAccountsParams creates a new PutCrayonAccountsParams object
// with the default values initialized.
func NewPutCrayonAccountsParams() *PutCrayonAccountsParams {
	var ()
	return &PutCrayonAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCrayonAccountsParamsWithTimeout creates a new PutCrayonAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCrayonAccountsParamsWithTimeout(timeout time.Duration) *PutCrayonAccountsParams {
	var ()
	return &PutCrayonAccountsParams{

		timeout: timeout,
	}
}

// NewPutCrayonAccountsParamsWithContext creates a new PutCrayonAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCrayonAccountsParamsWithContext(ctx context.Context) *PutCrayonAccountsParams {
	var ()
	return &PutCrayonAccountsParams{

		Context: ctx,
	}
}

// NewPutCrayonAccountsParamsWithHTTPClient creates a new PutCrayonAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCrayonAccountsParamsWithHTTPClient(client *http.Client) *PutCrayonAccountsParams {
	var ()
	return &PutCrayonAccountsParams{
		HTTPClient: client,
	}
}

/*PutCrayonAccountsParams contains all the parameters to send to the API endpoint
for the put crayon accounts operation typically these are written to a http.Request
*/
type PutCrayonAccountsParams struct {

	/*CrayonAccount*/
	CrayonAccount *models.CrayonAccount
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put crayon accounts params
func (o *PutCrayonAccountsParams) WithTimeout(timeout time.Duration) *PutCrayonAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put crayon accounts params
func (o *PutCrayonAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put crayon accounts params
func (o *PutCrayonAccountsParams) WithContext(ctx context.Context) *PutCrayonAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put crayon accounts params
func (o *PutCrayonAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put crayon accounts params
func (o *PutCrayonAccountsParams) WithHTTPClient(client *http.Client) *PutCrayonAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put crayon accounts params
func (o *PutCrayonAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrayonAccount adds the crayonAccount to the put crayon accounts params
func (o *PutCrayonAccountsParams) WithCrayonAccount(crayonAccount *models.CrayonAccount) *PutCrayonAccountsParams {
	o.SetCrayonAccount(crayonAccount)
	return o
}

// SetCrayonAccount adds the crayonAccount to the put crayon accounts params
func (o *PutCrayonAccountsParams) SetCrayonAccount(crayonAccount *models.CrayonAccount) {
	o.CrayonAccount = crayonAccount
}

// WithID adds the id to the put crayon accounts params
func (o *PutCrayonAccountsParams) WithID(id int32) *PutCrayonAccountsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put crayon accounts params
func (o *PutCrayonAccountsParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutCrayonAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CrayonAccount != nil {
		if err := r.SetBodyParam(o.CrayonAccount); err != nil {
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