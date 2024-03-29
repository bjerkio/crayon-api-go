// Code generated by go-swagger; DO NOT EDIT.

package reseller_sales_prices

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

// NewDeleteByFilterAsyncParams creates a new DeleteByFilterAsyncParams object
// with the default values initialized.
func NewDeleteByFilterAsyncParams() *DeleteByFilterAsyncParams {
	var ()
	return &DeleteByFilterAsyncParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteByFilterAsyncParamsWithTimeout creates a new DeleteByFilterAsyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteByFilterAsyncParamsWithTimeout(timeout time.Duration) *DeleteByFilterAsyncParams {
	var ()
	return &DeleteByFilterAsyncParams{

		timeout: timeout,
	}
}

// NewDeleteByFilterAsyncParamsWithContext creates a new DeleteByFilterAsyncParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteByFilterAsyncParamsWithContext(ctx context.Context) *DeleteByFilterAsyncParams {
	var ()
	return &DeleteByFilterAsyncParams{

		Context: ctx,
	}
}

// NewDeleteByFilterAsyncParamsWithHTTPClient creates a new DeleteByFilterAsyncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteByFilterAsyncParamsWithHTTPClient(client *http.Client) *DeleteByFilterAsyncParams {
	var ()
	return &DeleteByFilterAsyncParams{
		HTTPClient: client,
	}
}

/*DeleteByFilterAsyncParams contains all the parameters to send to the API endpoint
for the delete by filter async operation typically these are written to a http.Request
*/
type DeleteByFilterAsyncParams struct {

	/*FromDate*/
	FromDate *strfmt.DateTime
	/*ObjectID*/
	ObjectID *int32
	/*ObjectType*/
	ObjectType *string
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete by filter async params
func (o *DeleteByFilterAsyncParams) WithTimeout(timeout time.Duration) *DeleteByFilterAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete by filter async params
func (o *DeleteByFilterAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete by filter async params
func (o *DeleteByFilterAsyncParams) WithContext(ctx context.Context) *DeleteByFilterAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete by filter async params
func (o *DeleteByFilterAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete by filter async params
func (o *DeleteByFilterAsyncParams) WithHTTPClient(client *http.Client) *DeleteByFilterAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete by filter async params
func (o *DeleteByFilterAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFromDate adds the fromDate to the delete by filter async params
func (o *DeleteByFilterAsyncParams) WithFromDate(fromDate *strfmt.DateTime) *DeleteByFilterAsyncParams {
	o.SetFromDate(fromDate)
	return o
}

// SetFromDate adds the fromDate to the delete by filter async params
func (o *DeleteByFilterAsyncParams) SetFromDate(fromDate *strfmt.DateTime) {
	o.FromDate = fromDate
}

// WithObjectID adds the objectID to the delete by filter async params
func (o *DeleteByFilterAsyncParams) WithObjectID(objectID *int32) *DeleteByFilterAsyncParams {
	o.SetObjectID(objectID)
	return o
}

// SetObjectID adds the objectId to the delete by filter async params
func (o *DeleteByFilterAsyncParams) SetObjectID(objectID *int32) {
	o.ObjectID = objectID
}

// WithObjectType adds the objectType to the delete by filter async params
func (o *DeleteByFilterAsyncParams) WithObjectType(objectType *string) *DeleteByFilterAsyncParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the delete by filter async params
func (o *DeleteByFilterAsyncParams) SetObjectType(objectType *string) {
	o.ObjectType = objectType
}

// WithType adds the typeVar to the delete by filter async params
func (o *DeleteByFilterAsyncParams) WithType(typeVar *string) *DeleteByFilterAsyncParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the delete by filter async params
func (o *DeleteByFilterAsyncParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteByFilterAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FromDate != nil {

		// query param FromDate
		var qrFromDate strfmt.DateTime
		if o.FromDate != nil {
			qrFromDate = *o.FromDate
		}
		qFromDate := qrFromDate.String()
		if qFromDate != "" {
			if err := r.SetQueryParam("FromDate", qFromDate); err != nil {
				return err
			}
		}

	}

	if o.ObjectID != nil {

		// query param ObjectId
		var qrObjectID int32
		if o.ObjectID != nil {
			qrObjectID = *o.ObjectID
		}
		qObjectID := swag.FormatInt32(qrObjectID)
		if qObjectID != "" {
			if err := r.SetQueryParam("ObjectId", qObjectID); err != nil {
				return err
			}
		}

	}

	if o.ObjectType != nil {

		// query param ObjectType
		var qrObjectType string
		if o.ObjectType != nil {
			qrObjectType = *o.ObjectType
		}
		qObjectType := qrObjectType
		if qObjectType != "" {
			if err := r.SetQueryParam("ObjectType", qObjectType); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param Type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("Type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
