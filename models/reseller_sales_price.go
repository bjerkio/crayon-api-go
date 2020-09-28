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

// ResellerSalesPrice reseller sales price
//
// swagger:model ResellerSalesPrice
type ResellerSalesPrice struct {

	// from date
	// Format: date-time
	FromDate strfmt.DateTime `json:"FromDate,omitempty"`

	// object Id
	ObjectID int32 `json:"ObjectId,omitempty"`

	// object type
	// Enum: [Organization CustomerTenant Subscription SubscriptionAddon AzureSubscription]
	ObjectType string `json:"ObjectType,omitempty"`

	// price type
	// Enum: [Markup FixedPrice ListPrice]
	PriceType string `json:"PriceType,omitempty"`

	// type
	// Enum: [License Usage OneTime]
	Type string `json:"Type,omitempty"`

	// value
	Value float64 `json:"Value,omitempty"`
}

// Validate validates this reseller sales price
func (m *ResellerSalesPrice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResellerSalesPrice) validateFromDate(formats strfmt.Registry) error {

	if swag.IsZero(m.FromDate) { // not required
		return nil
	}

	if err := validate.FormatOf("FromDate", "body", "date-time", m.FromDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var resellerSalesPriceTypeObjectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Organization","CustomerTenant","Subscription","SubscriptionAddon","AzureSubscription"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resellerSalesPriceTypeObjectTypePropEnum = append(resellerSalesPriceTypeObjectTypePropEnum, v)
	}
}

const (

	// ResellerSalesPriceObjectTypeOrganization captures enum value "Organization"
	ResellerSalesPriceObjectTypeOrganization string = "Organization"

	// ResellerSalesPriceObjectTypeCustomerTenant captures enum value "CustomerTenant"
	ResellerSalesPriceObjectTypeCustomerTenant string = "CustomerTenant"

	// ResellerSalesPriceObjectTypeSubscription captures enum value "Subscription"
	ResellerSalesPriceObjectTypeSubscription string = "Subscription"

	// ResellerSalesPriceObjectTypeSubscriptionAddon captures enum value "SubscriptionAddon"
	ResellerSalesPriceObjectTypeSubscriptionAddon string = "SubscriptionAddon"

	// ResellerSalesPriceObjectTypeAzureSubscription captures enum value "AzureSubscription"
	ResellerSalesPriceObjectTypeAzureSubscription string = "AzureSubscription"
)

// prop value enum
func (m *ResellerSalesPrice) validateObjectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resellerSalesPriceTypeObjectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResellerSalesPrice) validateObjectType(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateObjectTypeEnum("ObjectType", "body", m.ObjectType); err != nil {
		return err
	}

	return nil
}

var resellerSalesPriceTypePriceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Markup","FixedPrice","ListPrice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resellerSalesPriceTypePriceTypePropEnum = append(resellerSalesPriceTypePriceTypePropEnum, v)
	}
}

const (

	// ResellerSalesPricePriceTypeMarkup captures enum value "Markup"
	ResellerSalesPricePriceTypeMarkup string = "Markup"

	// ResellerSalesPricePriceTypeFixedPrice captures enum value "FixedPrice"
	ResellerSalesPricePriceTypeFixedPrice string = "FixedPrice"

	// ResellerSalesPricePriceTypeListPrice captures enum value "ListPrice"
	ResellerSalesPricePriceTypeListPrice string = "ListPrice"
)

// prop value enum
func (m *ResellerSalesPrice) validatePriceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resellerSalesPriceTypePriceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResellerSalesPrice) validatePriceType(formats strfmt.Registry) error {

	if swag.IsZero(m.PriceType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePriceTypeEnum("PriceType", "body", m.PriceType); err != nil {
		return err
	}

	return nil
}

var resellerSalesPriceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["License","Usage","OneTime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resellerSalesPriceTypeTypePropEnum = append(resellerSalesPriceTypeTypePropEnum, v)
	}
}

const (

	// ResellerSalesPriceTypeLicense captures enum value "License"
	ResellerSalesPriceTypeLicense string = "License"

	// ResellerSalesPriceTypeUsage captures enum value "Usage"
	ResellerSalesPriceTypeUsage string = "Usage"

	// ResellerSalesPriceTypeOneTime captures enum value "OneTime"
	ResellerSalesPriceTypeOneTime string = "OneTime"
)

// prop value enum
func (m *ResellerSalesPrice) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resellerSalesPriceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResellerSalesPrice) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResellerSalesPrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResellerSalesPrice) UnmarshalBinary(b []byte) error {
	var res ResellerSalesPrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
