// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CategoryUsageCost category usage cost
//
// swagger:model CategoryUsageCost
type CategoryUsageCost struct {

	// amount
	Amount float64 `json:"Amount,omitempty"`

	// currency code
	CurrencyCode string `json:"CurrencyCode,omitempty"`

	// subcategory
	Subcategory string `json:"Subcategory,omitempty"`
}

// Validate validates this category usage cost
func (m *CategoryUsageCost) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CategoryUsageCost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CategoryUsageCost) UnmarshalBinary(b []byte) error {
	var res CategoryUsageCost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
