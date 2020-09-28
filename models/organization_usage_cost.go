// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationUsageCost organization usage cost
//
// swagger:model OrganizationUsageCost
type OrganizationUsageCost struct {

	// account Id
	AccountID int32 `json:"AccountId,omitempty"`

	// account name
	AccountName string `json:"AccountName,omitempty"`

	// amount
	Amount float64 `json:"Amount,omitempty"`

	// currency code
	CurrencyCode string `json:"CurrencyCode,omitempty"`

	// subscription Id
	SubscriptionID string `json:"SubscriptionId,omitempty"`

	// subscription name
	SubscriptionName string `json:"SubscriptionName,omitempty"`

	// supplier
	Supplier string `json:"Supplier,omitempty"`
}

// Validate validates this organization usage cost
func (m *OrganizationUsageCost) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationUsageCost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationUsageCost) UnmarshalBinary(b []byte) error {
	var res OrganizationUsageCost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
