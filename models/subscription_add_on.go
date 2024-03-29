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

// SubscriptionAddOn subscription add on
//
// swagger:model SubscriptionAddOn
type SubscriptionAddOn struct {

	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"CreationDate,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// markup
	Markup float64 `json:"Markup,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// order Id
	OrderID string `json:"OrderId,omitempty"`

	// organization
	Organization *ObjectReference `json:"Organization,omitempty"`

	// product
	Product *ProductReference `json:"Product,omitempty"`

	// publisher
	Publisher *ObjectReference `json:"Publisher,omitempty"`

	// publisher customer Id
	PublisherCustomerID string `json:"PublisherCustomerId,omitempty"`

	// publisher subscription Id
	PublisherSubscriptionID string `json:"PublisherSubscriptionId,omitempty"`

	// quantity
	Quantity int32 `json:"Quantity,omitempty"`

	// status
	// Enum: [None Active Suspended Deleted CustomerCancellation Converted Inactive All]
	Status string `json:"Status,omitempty"`

	// subscription tags
	SubscriptionTags *SubscriptionTags `json:"SubscriptionTags,omitempty"`
}

// Validate validates this subscription add on
func (m *SubscriptionAddOn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublisher(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionAddOn) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("CreationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionAddOn) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

func (m *SubscriptionAddOn) validateProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Product")
			}
			return err
		}
	}

	return nil
}

func (m *SubscriptionAddOn) validatePublisher(formats strfmt.Registry) error {

	if swag.IsZero(m.Publisher) { // not required
		return nil
	}

	if m.Publisher != nil {
		if err := m.Publisher.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Publisher")
			}
			return err
		}
	}

	return nil
}

var subscriptionAddOnTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Active","Suspended","Deleted","CustomerCancellation","Converted","Inactive","All"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionAddOnTypeStatusPropEnum = append(subscriptionAddOnTypeStatusPropEnum, v)
	}
}

const (

	// SubscriptionAddOnStatusNone captures enum value "None"
	SubscriptionAddOnStatusNone string = "None"

	// SubscriptionAddOnStatusActive captures enum value "Active"
	SubscriptionAddOnStatusActive string = "Active"

	// SubscriptionAddOnStatusSuspended captures enum value "Suspended"
	SubscriptionAddOnStatusSuspended string = "Suspended"

	// SubscriptionAddOnStatusDeleted captures enum value "Deleted"
	SubscriptionAddOnStatusDeleted string = "Deleted"

	// SubscriptionAddOnStatusCustomerCancellation captures enum value "CustomerCancellation"
	SubscriptionAddOnStatusCustomerCancellation string = "CustomerCancellation"

	// SubscriptionAddOnStatusConverted captures enum value "Converted"
	SubscriptionAddOnStatusConverted string = "Converted"

	// SubscriptionAddOnStatusInactive captures enum value "Inactive"
	SubscriptionAddOnStatusInactive string = "Inactive"

	// SubscriptionAddOnStatusAll captures enum value "All"
	SubscriptionAddOnStatusAll string = "All"
)

// prop value enum
func (m *SubscriptionAddOn) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, subscriptionAddOnTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionAddOn) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionAddOn) validateSubscriptionTags(formats strfmt.Registry) error {

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
func (m *SubscriptionAddOn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionAddOn) UnmarshalBinary(b []byte) error {
	var res SubscriptionAddOn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
