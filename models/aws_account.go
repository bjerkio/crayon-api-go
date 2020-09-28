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

// AwsAccount aws account
//
// swagger:model AwsAccount
type AwsAccount struct {

	// customer tenant type
	// Enum: [None T1 T2]
	CustomerTenantType string `json:"CustomerTenantType,omitempty"`

	// email
	Email string `json:"Email,omitempty"`

	// entity status
	// Enum: [None Removed]
	EntityStatus string `json:"EntityStatus,omitempty"`

	// external publisher customer Id
	ExternalPublisherCustomerID string `json:"ExternalPublisherCustomerId,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// invoice profile
	InvoiceProfile *ObjectReference `json:"InvoiceProfile,omitempty"`

	// is activated
	IsActivated bool `json:"IsActivated,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// organization
	Organization *Organization `json:"Organization,omitempty"`

	// publisher
	Publisher *ObjectReference `json:"Publisher,omitempty"`

	// reference
	Reference string `json:"Reference,omitempty"`

	// tags
	Tags *SubscriptionTags `json:"Tags,omitempty"`
}

// Validate validates this aws account
func (m *AwsAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerTenantType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublisher(formats); err != nil {
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

var awsAccountTypeCustomerTenantTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","T1","T2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsAccountTypeCustomerTenantTypePropEnum = append(awsAccountTypeCustomerTenantTypePropEnum, v)
	}
}

const (

	// AwsAccountCustomerTenantTypeNone captures enum value "None"
	AwsAccountCustomerTenantTypeNone string = "None"

	// AwsAccountCustomerTenantTypeT1 captures enum value "T1"
	AwsAccountCustomerTenantTypeT1 string = "T1"

	// AwsAccountCustomerTenantTypeT2 captures enum value "T2"
	AwsAccountCustomerTenantTypeT2 string = "T2"
)

// prop value enum
func (m *AwsAccount) validateCustomerTenantTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsAccountTypeCustomerTenantTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsAccount) validateCustomerTenantType(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerTenantType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCustomerTenantTypeEnum("CustomerTenantType", "body", m.CustomerTenantType); err != nil {
		return err
	}

	return nil
}

var awsAccountTypeEntityStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Removed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsAccountTypeEntityStatusPropEnum = append(awsAccountTypeEntityStatusPropEnum, v)
	}
}

const (

	// AwsAccountEntityStatusNone captures enum value "None"
	AwsAccountEntityStatusNone string = "None"

	// AwsAccountEntityStatusRemoved captures enum value "Removed"
	AwsAccountEntityStatusRemoved string = "Removed"
)

// prop value enum
func (m *AwsAccount) validateEntityStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsAccountTypeEntityStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsAccount) validateEntityStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityStatusEnum("EntityStatus", "body", m.EntityStatus); err != nil {
		return err
	}

	return nil
}

func (m *AwsAccount) validateInvoiceProfile(formats strfmt.Registry) error {

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

func (m *AwsAccount) validateOrganization(formats strfmt.Registry) error {

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

func (m *AwsAccount) validatePublisher(formats strfmt.Registry) error {

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

func (m *AwsAccount) validateTags(formats strfmt.Registry) error {

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
func (m *AwsAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsAccount) UnmarshalBinary(b []byte) error {
	var res AwsAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
