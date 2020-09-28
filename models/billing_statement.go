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

// BillingStatement billing statement
//
// swagger:model BillingStatement
type BillingStatement struct {

	// end date
	// Format: date-time
	EndDate strfmt.DateTime `json:"EndDate,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// invoice profile
	InvoiceProfile *ObjectReference `json:"InvoiceProfile,omitempty"`

	// order Id
	OrderID string `json:"OrderId,omitempty"`

	// organization
	Organization *ObjectReference `json:"Organization,omitempty"`

	// provision type
	// Enum: [None Seat Usage OneTime Crayon AzureMarketplace]
	ProvisionType string `json:"ProvisionType,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"StartDate,omitempty"`

	// total sales price
	TotalSalesPrice *Price `json:"TotalSalesPrice,omitempty"`
}

// Validate validates this billing statement
func (m *BillingStatement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalSalesPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingStatement) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingStatement) validateInvoiceProfile(formats strfmt.Registry) error {

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

func (m *BillingStatement) validateOrganization(formats strfmt.Registry) error {

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

var billingStatementTypeProvisionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Seat","Usage","OneTime","Crayon","AzureMarketplace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingStatementTypeProvisionTypePropEnum = append(billingStatementTypeProvisionTypePropEnum, v)
	}
}

const (

	// BillingStatementProvisionTypeNone captures enum value "None"
	BillingStatementProvisionTypeNone string = "None"

	// BillingStatementProvisionTypeSeat captures enum value "Seat"
	BillingStatementProvisionTypeSeat string = "Seat"

	// BillingStatementProvisionTypeUsage captures enum value "Usage"
	BillingStatementProvisionTypeUsage string = "Usage"

	// BillingStatementProvisionTypeOneTime captures enum value "OneTime"
	BillingStatementProvisionTypeOneTime string = "OneTime"

	// BillingStatementProvisionTypeCrayon captures enum value "Crayon"
	BillingStatementProvisionTypeCrayon string = "Crayon"

	// BillingStatementProvisionTypeAzureMarketplace captures enum value "AzureMarketplace"
	BillingStatementProvisionTypeAzureMarketplace string = "AzureMarketplace"
)

// prop value enum
func (m *BillingStatement) validateProvisionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billingStatementTypeProvisionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillingStatement) validateProvisionType(formats strfmt.Registry) error {

	if swag.IsZero(m.ProvisionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProvisionTypeEnum("ProvisionType", "body", m.ProvisionType); err != nil {
		return err
	}

	return nil
}

func (m *BillingStatement) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingStatement) validateTotalSalesPrice(formats strfmt.Registry) error {

	if swag.IsZero(m.TotalSalesPrice) { // not required
		return nil
	}

	if m.TotalSalesPrice != nil {
		if err := m.TotalSalesPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TotalSalesPrice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingStatement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingStatement) UnmarshalBinary(b []byte) error {
	var res BillingStatement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
