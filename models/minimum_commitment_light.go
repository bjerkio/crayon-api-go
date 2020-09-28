// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MinimumCommitmentLight minimum commitment light
//
// swagger:model MinimumCommitmentLight
type MinimumCommitmentLight struct {

	// Id
	ID int32 `json:"Id,omitempty"`

	// months
	Months int32 `json:"Months,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this minimum commitment light
func (m *MinimumCommitmentLight) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MinimumCommitmentLight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MinimumCommitmentLight) UnmarshalBinary(b []byte) error {
	var res MinimumCommitmentLight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
