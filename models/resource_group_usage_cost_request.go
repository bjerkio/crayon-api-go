// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourceGroupUsageCostRequest resource group usage cost request
//
// swagger:model ResourceGroupUsageCostRequest
type ResourceGroupUsageCostRequest struct {

	// currency code
	CurrencyCode string `json:"CurrencyCode,omitempty"`

	// from
	// Format: date-time
	From strfmt.DateTime `json:"From,omitempty"`

	// reseller customer Id
	ResellerCustomerID int32 `json:"ResellerCustomerId,omitempty"`

	// resource group
	ResourceGroup string `json:"ResourceGroup,omitempty"`

	// subscription Id
	SubscriptionID string `json:"SubscriptionId,omitempty"`

	// to
	// Format: date-time
	To strfmt.DateTime `json:"To,omitempty"`
}

// Validate validates this resource group usage cost request
func (m *ResourceGroupUsageCostRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceGroupUsageCostRequest) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if err := validate.FormatOf("From", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResourceGroupUsageCostRequest) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	if err := validate.FormatOf("To", "body", "date-time", m.To.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceGroupUsageCostRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceGroupUsageCostRequest) UnmarshalBinary(b []byte) error {
	var res ResourceGroupUsageCostRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}