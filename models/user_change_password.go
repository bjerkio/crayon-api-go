// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserChangePassword user change password
//
// swagger:model UserChangePassword
type UserChangePassword struct {

	// new password
	NewPassword string `json:"NewPassword,omitempty"`

	// old password
	OldPassword string `json:"OldPassword,omitempty"`

	// user Id
	UserID string `json:"UserId,omitempty"`
}

// Validate validates this user change password
func (m *UserChangePassword) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserChangePassword) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserChangePassword) UnmarshalBinary(b []byte) error {
	var res UserChangePassword
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
