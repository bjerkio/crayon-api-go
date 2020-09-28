// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserProfile user profile
//
// swagger:model UserProfile
type UserProfile struct {

	// email
	Email string `json:"Email,omitempty"`

	// first name
	FirstName string `json:"FirstName,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last name
	LastName string `json:"LastName,omitempty"`

	// user name
	UserName string `json:"UserName,omitempty"`
}

// Validate validates this user profile
func (m *UserProfile) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserProfile) UnmarshalBinary(b []byte) error {
	var res UserProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}