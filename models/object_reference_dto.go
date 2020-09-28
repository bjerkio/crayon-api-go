// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ObjectReferenceDto object reference dto
//
// swagger:model ObjectReferenceDto
type ObjectReferenceDto struct {

	// Id
	ID int32 `json:"Id,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this object reference dto
func (m *ObjectReferenceDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectReferenceDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectReferenceDto) UnmarshalBinary(b []byte) error {
	var res ObjectReferenceDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}