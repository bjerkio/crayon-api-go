// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APICollectionOfBoolean Api collection of boolean
//
// swagger:model ApiCollectionOfBoolean
type APICollectionOfBoolean struct {

	// items
	Items []bool `json:"Items"`

	// total hits
	TotalHits int64 `json:"TotalHits,omitempty"`
}

// Validate validates this Api collection of boolean
func (m *APICollectionOfBoolean) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APICollectionOfBoolean) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICollectionOfBoolean) UnmarshalBinary(b []byte) error {
	var res APICollectionOfBoolean
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
