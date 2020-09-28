// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerTenantProfile customer tenant profile
//
// swagger:model CustomerTenantProfile
type CustomerTenantProfile struct {

	// address
	Address *CustomerTenantAddress `json:"Address,omitempty"`

	// agreement
	Agreement *CustomerTenantAgreement `json:"Agreement,omitempty"`

	// contact
	Contact *CustomerTenantContact `json:"Contact,omitempty"`

	// culture code
	CultureCode string `json:"CultureCode,omitempty"`

	// language code
	LanguageCode string `json:"LanguageCode,omitempty"`
}

// Validate validates this customer tenant profile
func (m *CustomerTenantProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgreement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContact(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerTenantProfile) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Address")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerTenantProfile) validateAgreement(formats strfmt.Registry) error {

	if swag.IsZero(m.Agreement) { // not required
		return nil
	}

	if m.Agreement != nil {
		if err := m.Agreement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Agreement")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerTenantProfile) validateContact(formats strfmt.Registry) error {

	if swag.IsZero(m.Contact) { // not required
		return nil
	}

	if m.Contact != nil {
		if err := m.Contact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Contact")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerTenantProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerTenantProfile) UnmarshalBinary(b []byte) error {
	var res CustomerTenantProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}