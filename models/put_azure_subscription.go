// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutAzureSubscription put azure subscription
//
// swagger:model PutAzureSubscription
type PutAzureSubscription struct {

	// Id
	ID int32 `json:"Id,omitempty"`

	// invoice profile
	InvoiceProfile *ObjectReference `json:"InvoiceProfile,omitempty"`

	// tags
	Tags *AzureSubscriptionTags `json:"Tags,omitempty"`
}

// Validate validates this put azure subscription
func (m *PutAzureSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvoiceProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutAzureSubscription) validateInvoiceProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoiceProfile) { // not required
		return nil
	}

	if m.InvoiceProfile != nil {
		if err := m.InvoiceProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InvoiceProfile")
			}
			return err
		}
	}

	return nil
}

func (m *PutAzureSubscription) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if m.Tags != nil {
		if err := m.Tags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Tags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutAzureSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutAzureSubscription) UnmarshalBinary(b []byte) error {
	var res PutAzureSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
