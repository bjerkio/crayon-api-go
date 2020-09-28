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

// GroupedBillingStatement grouped billing statement
//
// swagger:model GroupedBillingStatement
type GroupedBillingStatement struct {

	// end date
	// Format: date-time
	EndDate strfmt.DateTime `json:"EndDate,omitempty"`

	// group Id
	GroupID int32 `json:"GroupId,omitempty"`

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

// Validate validates this grouped billing statement
func (m *GroupedBillingStatement) Validate(formats strfmt.Registry) error {
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

func (m *GroupedBillingStatement) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GroupedBillingStatement) validateInvoiceProfile(formats strfmt.Registry) error {

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

func (m *GroupedBillingStatement) validateOrganization(formats strfmt.Registry) error {

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

var groupedBillingStatementTypeProvisionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Seat","Usage","OneTime","Crayon","AzureMarketplace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupedBillingStatementTypeProvisionTypePropEnum = append(groupedBillingStatementTypeProvisionTypePropEnum, v)
	}
}

const (

	// GroupedBillingStatementProvisionTypeNone captures enum value "None"
	GroupedBillingStatementProvisionTypeNone string = "None"

	// GroupedBillingStatementProvisionTypeSeat captures enum value "Seat"
	GroupedBillingStatementProvisionTypeSeat string = "Seat"

	// GroupedBillingStatementProvisionTypeUsage captures enum value "Usage"
	GroupedBillingStatementProvisionTypeUsage string = "Usage"

	// GroupedBillingStatementProvisionTypeOneTime captures enum value "OneTime"
	GroupedBillingStatementProvisionTypeOneTime string = "OneTime"

	// GroupedBillingStatementProvisionTypeCrayon captures enum value "Crayon"
	GroupedBillingStatementProvisionTypeCrayon string = "Crayon"

	// GroupedBillingStatementProvisionTypeAzureMarketplace captures enum value "AzureMarketplace"
	GroupedBillingStatementProvisionTypeAzureMarketplace string = "AzureMarketplace"
)

// prop value enum
func (m *GroupedBillingStatement) validateProvisionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupedBillingStatementTypeProvisionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupedBillingStatement) validateProvisionType(formats strfmt.Registry) error {

	if swag.IsZero(m.ProvisionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProvisionTypeEnum("ProvisionType", "body", m.ProvisionType); err != nil {
		return err
	}

	return nil
}

func (m *GroupedBillingStatement) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GroupedBillingStatement) validateTotalSalesPrice(formats strfmt.Registry) error {

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
func (m *GroupedBillingStatement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupedBillingStatement) UnmarshalBinary(b []byte) error {
	var res GroupedBillingStatement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
