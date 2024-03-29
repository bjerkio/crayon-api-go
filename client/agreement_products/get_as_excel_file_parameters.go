// Code generated by go-swagger; DO NOT EDIT.

package agreement_products

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

	"github.com/bjerkio/crayon-api-go/models"
)

// NewGetAsExcelFileParams creates a new GetAsExcelFileParams object
// with the default values initialized.
func NewGetAsExcelFileParams() *GetAsExcelFileParams {
	var ()
	return &GetAsExcelFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAsExcelFileParamsWithTimeout creates a new GetAsExcelFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAsExcelFileParamsWithTimeout(timeout time.Duration) *GetAsExcelFileParams {
	var ()
	return &GetAsExcelFileParams{

		timeout: timeout,
	}
}

// NewGetAsExcelFileParamsWithContext creates a new GetAsExcelFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAsExcelFileParamsWithContext(ctx context.Context) *GetAsExcelFileParams {
	var ()
	return &GetAsExcelFileParams{

		Context: ctx,
	}
}

// NewGetAsExcelFileParamsWithHTTPClient creates a new GetAsExcelFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAsExcelFileParamsWithHTTPClient(client *http.Client) *GetAsExcelFileParams {
	var ()
	return &GetAsExcelFileParams{
		HTTPClient: client,
	}
}

/*GetAsExcelFileParams contains all the parameters to send to the API endpoint
for the get as excel file operation typically these are written to a http.Request
*/
type GetAsExcelFileParams struct {

	/*Filter*/
	Filter *models.AgreementProductFilter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get as excel file params
func (o *GetAsExcelFileParams) WithTimeout(timeout time.Duration) *GetAsExcelFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get as excel file params
func (o *GetAsExcelFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get as excel file params
func (o *GetAsExcelFileParams) WithContext(ctx context.Context) *GetAsExcelFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get as excel file params
func (o *GetAsExcelFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get as excel file params
func (o *GetAsExcelFileParams) WithHTTPClient(client *http.Client) *GetAsExcelFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get as excel file params
func (o *GetAsExcelFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get as excel file params
func (o *GetAsExcelFileParams) WithFilter(filter *models.AgreementProductFilter) *GetAsExcelFileParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get as excel file params
func (o *GetAsExcelFileParams) SetFilter(filter *models.AgreementProductFilter) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *GetAsExcelFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {
		if err := r.SetBodyParam(o.Filter); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
