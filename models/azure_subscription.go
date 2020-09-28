// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureSubscription azure subscription
//
// swagger:model AzureSubscription
type AzureSubscription struct {

	// azure plan Id
	AzurePlanID int32 `json:"AzurePlanId,omitempty"`

	// friendly name
	FriendlyName string `json:"FriendlyName,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// invoice profile
	InvoiceProfile *ObjectReference `json:"InvoiceProfile,omitempty"`

	// publisher subscription Id
	PublisherSubscriptionID string `json:"PublisherSubscriptionId,omitempty"`

	// status
	Status string `json:"Status,omitempty"`

	// tags
	Tags *AzureSubscriptionTags `json:"Tags,omitempty"`
}

// Validate validates this azure subscription
func (m *AzureSubscription) Validate(formats strfmt.Registry) error {
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

func (m *AzureSubscription) validateInvoiceProfile(formats strfmt.Registry) error {

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

func (m *AzureSubscription) validateTags(formats strfmt.Registry) error {

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
func (m *AzureSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureSubscription) UnmarshalBinary(b []byte) error {
	var res AzureSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
