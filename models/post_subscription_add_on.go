// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostSubscriptionAddOn post subscription add on
//
// swagger:model PostSubscriptionAddOn
type PostSubscriptionAddOn struct {

	// offer Id
	OfferID string `json:"OfferId,omitempty"`

	// quantity
	Quantity int32 `json:"Quantity,omitempty"`

	// subscription tags
	SubscriptionTags *SubscriptionTags `json:"SubscriptionTags,omitempty"`
}

// Validate validates this post subscription add on
func (m *PostSubscriptionAddOn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptionTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostSubscriptionAddOn) validateSubscriptionTags(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionTags) { // not required
		return nil
	}

	if m.SubscriptionTags != nil {
		if err := m.SubscriptionTags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SubscriptionTags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostSubscriptionAddOn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostSubscriptionAddOn) UnmarshalBinary(b []byte) error {
	var res PostSubscriptionAddOn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
