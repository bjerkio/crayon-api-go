// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GoogleOrder google order
//
// swagger:model GoogleOrder
type GoogleOrder struct {

	// consumer city
	ConsumerCity string `json:"ConsumerCity,omitempty"`

	// consumer contact email
	ConsumerContactEmail string `json:"ConsumerContactEmail,omitempty"`

	// consumer contact name
	ConsumerContactName string `json:"ConsumerContactName,omitempty"`

	// consumer country
	ConsumerCountry string `json:"ConsumerCountry,omitempty"`

	// consumer full legal name
	ConsumerFullLegalName string `json:"ConsumerFullLegalName,omitempty"`

	// consumer phone number
	ConsumerPhoneNumber string `json:"ConsumerPhoneNumber,omitempty"`

	// consumer state
	ConsumerState string `json:"ConsumerState,omitempty"`

	// consumer street address
	ConsumerStreetAddress string `json:"ConsumerStreetAddress,omitempty"`

	// consumer zip code
	ConsumerZipCode string `json:"ConsumerZipCode,omitempty"`

	// contact email
	ContactEmail string `json:"ContactEmail,omitempty"`

	// contact name
	ContactName string `json:"ContactName,omitempty"`

	// domain name
	DomainName string `json:"DomainName,omitempty"`

	// invoice profile Id
	InvoiceProfileID int32 `json:"InvoiceProfileId,omitempty"`

	// lines
	Lines []*GoogleOrderLine `json:"Lines"`

	// organization Id
	OrganizationID int32 `json:"OrganizationId,omitempty"`

	// primary admin name
	PrimaryAdminName string `json:"PrimaryAdminName,omitempty"`

	// primary admin user name
	PrimaryAdminUserName string `json:"PrimaryAdminUserName,omitempty"`

	// subscription start date
	// Format: date-time
	SubscriptionStartDate strfmt.DateTime `json:"SubscriptionStartDate,omitempty"`
}

// Validate validates this google order
func (m *GoogleOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GoogleOrder) validateLines(formats strfmt.Registry) error {

	if swag.IsZero(m.Lines) { // not required
		return nil
	}

	for i := 0; i < len(m.Lines); i++ {
		if swag.IsZero(m.Lines[i]) { // not required
			continue
		}

		if m.Lines[i] != nil {
			if err := m.Lines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GoogleOrder) validateSubscriptionStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("SubscriptionStartDate", "body", "date-time", m.SubscriptionStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GoogleOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GoogleOrder) UnmarshalBinary(b []byte) error {
	var res GoogleOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
