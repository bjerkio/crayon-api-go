// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomerTenantAgreement customer tenant agreement
//
// swagger:model CustomerTenantAgreement
type CustomerTenantAgreement struct {

	// accepted
	Accepted bool `json:"Accepted,omitempty"`

	// agreement type
	// Enum: [MicrosoftCloudAgreement MicrosoftCustomerAgreement]
	AgreementType string `json:"AgreementType,omitempty"`

	// date agreed
	// Format: date-time
	DateAgreed strfmt.DateTime `json:"DateAgreed,omitempty"`

	// email
	Email string `json:"Email,omitempty"`

	// first name
	FirstName string `json:"FirstName,omitempty"`

	// last name
	LastName string `json:"LastName,omitempty"`

	// phone number
	PhoneNumber string `json:"PhoneNumber,omitempty"`

	// same as primary contact
	SameAsPrimaryContact bool `json:"SameAsPrimaryContact,omitempty"`
}

// Validate validates this customer tenant agreement
func (m *CustomerTenantAgreement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgreementType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateAgreed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customerTenantAgreementTypeAgreementTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MicrosoftCloudAgreement","MicrosoftCustomerAgreement"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTenantAgreementTypeAgreementTypePropEnum = append(customerTenantAgreementTypeAgreementTypePropEnum, v)
	}
}

const (

	// CustomerTenantAgreementAgreementTypeMicrosoftCloudAgreement captures enum value "MicrosoftCloudAgreement"
	CustomerTenantAgreementAgreementTypeMicrosoftCloudAgreement string = "MicrosoftCloudAgreement"

	// CustomerTenantAgreementAgreementTypeMicrosoftCustomerAgreement captures enum value "MicrosoftCustomerAgreement"
	CustomerTenantAgreementAgreementTypeMicrosoftCustomerAgreement string = "MicrosoftCustomerAgreement"
)

// prop value enum
func (m *CustomerTenantAgreement) validateAgreementTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, customerTenantAgreementTypeAgreementTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomerTenantAgreement) validateAgreementType(formats strfmt.Registry) error {

	if swag.IsZero(m.AgreementType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAgreementTypeEnum("AgreementType", "body", m.AgreementType); err != nil {
		return err
	}

	return nil
}

func (m *CustomerTenantAgreement) validateDateAgreed(formats strfmt.Registry) error {

	if swag.IsZero(m.DateAgreed) { // not required
		return nil
	}

	if err := validate.FormatOf("DateAgreed", "body", "date-time", m.DateAgreed.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerTenantAgreement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerTenantAgreement) UnmarshalBinary(b []byte) error {
	var res CustomerTenantAgreement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
