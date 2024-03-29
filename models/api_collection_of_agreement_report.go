// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APICollectionOfAgreementReport Api collection of agreement report
//
// swagger:model ApiCollectionOfAgreementReport
type APICollectionOfAgreementReport struct {

	// items
	Items []*AgreementReport `json:"Items"`

	// total hits
	TotalHits int64 `json:"TotalHits,omitempty"`
}

// Validate validates this Api collection of agreement report
func (m *APICollectionOfAgreementReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APICollectionOfAgreementReport) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APICollectionOfAgreementReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICollectionOfAgreementReport) UnmarshalBinary(b []byte) error {
	var res APICollectionOfAgreementReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
