// Code generated by go-swagger; DO NOT EDIT.

package product_containers

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

// NewPatchProductRowParams creates a new PatchProductRowParams object
// with the default values initialized.
func NewPatchProductRowParams() *PatchProductRowParams {
	var ()
	return &PatchProductRowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchProductRowParamsWithTimeout creates a new PatchProductRowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchProductRowParamsWithTimeout(timeout time.Duration) *PatchProductRowParams {
	var ()
	return &PatchProductRowParams{

		timeout: timeout,
	}
}

// NewPatchProductRowParamsWithContext creates a new PatchProductRowParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchProductRowParamsWithContext(ctx context.Context) *PatchProductRowParams {
	var ()
	return &PatchProductRowParams{

		Context: ctx,
	}
}

// NewPatchProductRowParamsWithHTTPClient creates a new PatchProductRowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchProductRowParamsWithHTTPClient(client *http.Client) *PatchProductRowParams {
	var ()
	return &PatchProductRowParams{
		HTTPClient: client,
	}
}

/*PatchProductRowParams contains all the parameters to send to the API endpoint
for the patch product row operation typically these are written to a http.Request
*/
type PatchProductRowParams struct {

	/*ProductContainerID*/
	ProductContainerID int32
	/*ProductRowID*/
	ProductRowID int32
	/*ProductRowPatch*/
	ProductRowPatch *models.ProductRowPatch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch product row params
func (o *PatchProductRowParams) WithTimeout(timeout time.Duration) *PatchProductRowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch product row params
func (o *PatchProductRowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch product row params
func (o *PatchProductRowParams) WithContext(ctx context.Context) *PatchProductRowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch product row params
func (o *PatchProductRowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch product row params
func (o *PatchProductRowParams) WithHTTPClient(client *http.Client) *PatchProductRowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch product row params
func (o *PatchProductRowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductContainerID adds the productContainerID to the patch product row params
func (o *PatchProductRowParams) WithProductContainerID(productContainerID int32) *PatchProductRowParams {
	o.SetProductContainerID(productContainerID)
	return o
}

// SetProductContainerID adds the productContainerId to the patch product row params
func (o *PatchProductRowParams) SetProductContainerID(productContainerID int32) {
	o.ProductContainerID = productContainerID
}

// WithProductRowID adds the productRowID to the patch product row params
func (o *PatchProductRowParams) WithProductRowID(productRowID int32) *PatchProductRowParams {
	o.SetProductRowID(productRowID)
	return o
}

// SetProductRowID adds the productRowId to the patch product row params
func (o *PatchProductRowParams) SetProductRowID(productRowID int32) {
	o.ProductRowID = productRowID
}

// WithProductRowPatch adds the productRowPatch to the patch product row params
func (o *PatchProductRowParams) WithProductRowPatch(productRowPatch *models.ProductRowPatch) *PatchProductRowParams {
	o.SetProductRowPatch(productRowPatch)
	return o
}

// SetProductRowPatch adds the productRowPatch to the patch product row params
func (o *PatchProductRowParams) SetProductRowPatch(productRowPatch *models.ProductRowPatch) {
	o.ProductRowPatch = productRowPatch
}

// WriteToRequest writes these params to a swagger request
func (o *PatchProductRowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param productContainerId
	if err := r.SetPathParam("productContainerId", swag.FormatInt32(o.ProductContainerID)); err != nil {
		return err
	}

	// path param productRowId
	if err := r.SetPathParam("productRowId", swag.FormatInt32(o.ProductRowID)); err != nil {
		return err
	}

	if o.ProductRowPatch != nil {
		if err := r.SetBodyParam(o.ProductRowPatch); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}