// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerTenantUser customer tenant user
//
// swagger:model CustomerTenantUser
type CustomerTenantUser struct {

	// password
	Password string `json:"Password,omitempty"`

	// user name
	UserName string `json:"UserName,omitempty"`
}

// Validate validates this customer tenant user
func (m *CustomerTenantUser) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerTenantUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerTenantUser) UnmarshalBinary(b []byte) error {
	var res CustomerTenantUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}