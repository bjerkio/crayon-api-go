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

// Asset asset
//
// swagger:model Asset
type Asset struct {

	// asset type
	// Enum: [Reservation Software Subscription]
	AssetType string `json:"AssetType,omitempty"`

	// billing cycle
	// Enum: [Unknown Monthly Annual None OneTime]
	BillingCycle string `json:"BillingCycle,omitempty"`

	// changed by
	ChangedBy string `json:"ChangedBy,omitempty"`

	// created date
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"CreatedDate,omitempty"`

	// expiration date
	// Format: date-time
	ExpirationDate strfmt.DateTime `json:"ExpirationDate,omitempty"`

	// external order Id
	ExternalOrderID string `json:"ExternalOrderId,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// modified date
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"ModifiedDate,omitempty"`

	// product Id
	ProductID int32 `json:"ProductId,omitempty"`

	// product variant
	ProductVariant *ProductVariant `json:"ProductVariant,omitempty"`

	// product variant Id
	ProductVariantID int32 `json:"ProductVariantId,omitempty"`

	// program Id
	ProgramID int32 `json:"ProgramId,omitempty"`

	// publisher Id
	PublisherID int32 `json:"PublisherId,omitempty"`

	// purchase currency
	PurchaseCurrency string `json:"PurchaseCurrency,omitempty"`

	// purchase date
	// Format: date-time
	PurchaseDate strfmt.DateTime `json:"PurchaseDate,omitempty"`

	// purchase price
	PurchasePrice float64 `json:"PurchasePrice,omitempty"`

	// quantity
	Quantity int32 `json:"Quantity,omitempty"`

	// reseller customer Id
	ResellerCustomerID int32 `json:"ResellerCustomerId,omitempty"`

	// reseller price type
	// Enum: [None Margin Markup FixedPrice ListPrice]
	ResellerPriceType string `json:"ResellerPriceType,omitempty"`

	// reseller price type value
	ResellerPriceTypeValue float64 `json:"ResellerPriceTypeValue,omitempty"`

	// reservation Id
	ReservationID string `json:"ReservationId,omitempty"`

	// reservation used in subscription
	ReservationUsedInSubscription *SubscriptionLite `json:"ReservationUsedInSubscription,omitempty"`

	// reservation used in subscription Id
	ReservationUsedInSubscriptionID int32 `json:"ReservationUsedInSubscriptionId,omitempty"`

	// reserved instance artifact resource Id
	ReservedInstanceArtifactResourceID string `json:"ReservedInstanceArtifactResourceId,omitempty"`

	// sales currency
	SalesCurrency string `json:"SalesCurrency,omitempty"`

	// sales price
	SalesPrice float64 `json:"SalesPrice,omitempty"`

	// scope
	// Enum: [Shared Single]
	Scope string `json:"Scope,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"StartDate,omitempty"`

	// status
	// Enum: [None Fulfilling Succeeded Cancelled Expired All]
	Status string `json:"Status,omitempty"`

	// tags
	Tags *AssetTags `json:"Tags,omitempty"`
}

// Validate validates this asset
func (m *Asset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductVariant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerPriceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReservationUsedInSubscription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

var assetTypeAssetTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Reservation","Software","Subscription"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetTypeAssetTypePropEnum = append(assetTypeAssetTypePropEnum, v)
	}
}

const (

	// AssetAssetTypeReservation captures enum value "Reservation"
	AssetAssetTypeReservation string = "Reservation"

	// AssetAssetTypeSoftware captures enum value "Software"
	AssetAssetTypeSoftware string = "Software"

	// AssetAssetTypeSubscription captures enum value "Subscription"
	AssetAssetTypeSubscription string = "Subscription"
)

// prop value enum
func (m *Asset) validateAssetTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assetTypeAssetTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Asset) validateAssetType(formats strfmt.Registry) error {

	if swag.IsZero(m.AssetType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAssetTypeEnum("AssetType", "body", m.AssetType); err != nil {
		return err
	}

	return nil
}

var assetTypeBillingCyclePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Monthly","Annual","None","OneTime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetTypeBillingCyclePropEnum = append(assetTypeBillingCyclePropEnum, v)
	}
}

const (

	// AssetBillingCycleUnknown captures enum value "Unknown"
	AssetBillingCycleUnknown string = "Unknown"

	// AssetBillingCycleMonthly captures enum value "Monthly"
	AssetBillingCycleMonthly string = "Monthly"

	// AssetBillingCycleAnnual captures enum value "Annual"
	AssetBillingCycleAnnual string = "Annual"

	// AssetBillingCycleNone captures enum value "None"
	AssetBillingCycleNone string = "None"

	// AssetBillingCycleOneTime captures enum value "OneTime"
	AssetBillingCycleOneTime string = "OneTime"
)

// prop value enum
func (m *Asset) validateBillingCycleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assetTypeBillingCyclePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Asset) validateBillingCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingCycle) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingCycleEnum("BillingCycle", "body", m.BillingCycle); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateExpirationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpirationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ExpirationDate", "body", "date-time", m.ExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ModifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateProductVariant(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductVariant) { // not required
		return nil
	}

	if m.ProductVariant != nil {
		if err := m.ProductVariant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProductVariant")
			}
			return err
		}
	}

	return nil
}

func (m *Asset) validatePurchaseDate(formats strfmt.Registry) error {

	if swag.IsZero(m.PurchaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("PurchaseDate", "body", "date-time", m.PurchaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var assetTypeResellerPriceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Margin","Markup","FixedPrice","ListPrice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetTypeResellerPriceTypePropEnum = append(assetTypeResellerPriceTypePropEnum, v)
	}
}

const (

	// AssetResellerPriceTypeNone captures enum value "None"
	AssetResellerPriceTypeNone string = "None"

	// AssetResellerPriceTypeMargin captures enum value "Margin"
	AssetResellerPriceTypeMargin string = "Margin"

	// AssetResellerPriceTypeMarkup captures enum value "Markup"
	AssetResellerPriceTypeMarkup string = "Markup"

	// AssetResellerPriceTypeFixedPrice captures enum value "FixedPrice"
	AssetResellerPriceTypeFixedPrice string = "FixedPrice"

	// AssetResellerPriceTypeListPrice captures enum value "ListPrice"
	AssetResellerPriceTypeListPrice string = "ListPrice"
)

// prop value enum
func (m *Asset) validateResellerPriceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assetTypeResellerPriceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Asset) validateResellerPriceType(formats strfmt.Registry) error {

	if swag.IsZero(m.ResellerPriceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateResellerPriceTypeEnum("ResellerPriceType", "body", m.ResellerPriceType); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateReservationUsedInSubscription(formats strfmt.Registry) error {

	if swag.IsZero(m.ReservationUsedInSubscription) { // not required
		return nil
	}

	if m.ReservationUsedInSubscription != nil {
		if err := m.ReservationUsedInSubscription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ReservationUsedInSubscription")
			}
			return err
		}
	}

	return nil
}

var assetTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Shared","Single"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetTypeScopePropEnum = append(assetTypeScopePropEnum, v)
	}
}

const (

	// AssetScopeShared captures enum value "Shared"
	AssetScopeShared string = "Shared"

	// AssetScopeSingle captures enum value "Single"
	AssetScopeSingle string = "Single"
)

// prop value enum
func (m *Asset) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assetTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Asset) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("Scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var assetTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Fulfilling","Succeeded","Cancelled","Expired","All"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetTypeStatusPropEnum = append(assetTypeStatusPropEnum, v)
	}
}

const (

	// AssetStatusNone captures enum value "None"
	AssetStatusNone string = "None"

	// AssetStatusFulfilling captures enum value "Fulfilling"
	AssetStatusFulfilling string = "Fulfilling"

	// AssetStatusSucceeded captures enum value "Succeeded"
	AssetStatusSucceeded string = "Succeeded"

	// AssetStatusCancelled captures enum value "Cancelled"
	AssetStatusCancelled string = "Cancelled"

	// AssetStatusExpired captures enum value "Expired"
	AssetStatusExpired string = "Expired"

	// AssetStatusAll captures enum value "All"
	AssetStatusAll string = "All"
)

// prop value enum
func (m *Asset) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assetTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Asset) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateTags(formats strfmt.Registry) error {

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
func (m *Asset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Asset) UnmarshalBinary(b []byte) error {
	var res Asset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
