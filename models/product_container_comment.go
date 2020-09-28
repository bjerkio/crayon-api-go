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

// ProductContainerComment product container comment
//
// swagger:model ProductContainerComment
type ProductContainerComment struct {

	// Id
	ID int32 `json:"Id,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// modified
	// Format: date-time
	Modified strfmt.DateTime `json:"Modified,omitempty"`

	// product container comment type
	// Enum: [None Comment DraftCreated DraftModified QuoteCreated QuoteModified RequestCreated RequestModifed OrderSubmitted OrderInvoiced Removed TemplateCreated TemplateModifed RequestDeclined SystemMessage]
	ProductContainerCommentType string `json:"ProductContainerCommentType,omitempty"`

	// product container Id
	ProductContainerID int32 `json:"ProductContainerId,omitempty"`

	// removed
	Removed bool `json:"Removed,omitempty"`

	// time stamp
	// Format: date-time
	TimeStamp strfmt.DateTime `json:"TimeStamp,omitempty"`

	// user
	User *ProductContainerCommentUser `json:"User,omitempty"`
}

// Validate validates this product container comment
func (m *ProductContainerComment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductContainerCommentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeStamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductContainerComment) validateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.Modified) { // not required
		return nil
	}

	if err := validate.FormatOf("Modified", "body", "date-time", m.Modified.String(), formats); err != nil {
		return err
	}

	return nil
}

var productContainerCommentTypeProductContainerCommentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Comment","DraftCreated","DraftModified","QuoteCreated","QuoteModified","RequestCreated","RequestModifed","OrderSubmitted","OrderInvoiced","Removed","TemplateCreated","TemplateModifed","RequestDeclined","SystemMessage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productContainerCommentTypeProductContainerCommentTypePropEnum = append(productContainerCommentTypeProductContainerCommentTypePropEnum, v)
	}
}

const (

	// ProductContainerCommentProductContainerCommentTypeNone captures enum value "None"
	ProductContainerCommentProductContainerCommentTypeNone string = "None"

	// ProductContainerCommentProductContainerCommentTypeComment captures enum value "Comment"
	ProductContainerCommentProductContainerCommentTypeComment string = "Comment"

	// ProductContainerCommentProductContainerCommentTypeDraftCreated captures enum value "DraftCreated"
	ProductContainerCommentProductContainerCommentTypeDraftCreated string = "DraftCreated"

	// ProductContainerCommentProductContainerCommentTypeDraftModified captures enum value "DraftModified"
	ProductContainerCommentProductContainerCommentTypeDraftModified string = "DraftModified"

	// ProductContainerCommentProductContainerCommentTypeQuoteCreated captures enum value "QuoteCreated"
	ProductContainerCommentProductContainerCommentTypeQuoteCreated string = "QuoteCreated"

	// ProductContainerCommentProductContainerCommentTypeQuoteModified captures enum value "QuoteModified"
	ProductContainerCommentProductContainerCommentTypeQuoteModified string = "QuoteModified"

	// ProductContainerCommentProductContainerCommentTypeRequestCreated captures enum value "RequestCreated"
	ProductContainerCommentProductContainerCommentTypeRequestCreated string = "RequestCreated"

	// ProductContainerCommentProductContainerCommentTypeRequestModifed captures enum value "RequestModifed"
	ProductContainerCommentProductContainerCommentTypeRequestModifed string = "RequestModifed"

	// ProductContainerCommentProductContainerCommentTypeOrderSubmitted captures enum value "OrderSubmitted"
	ProductContainerCommentProductContainerCommentTypeOrderSubmitted string = "OrderSubmitted"

	// ProductContainerCommentProductContainerCommentTypeOrderInvoiced captures enum value "OrderInvoiced"
	ProductContainerCommentProductContainerCommentTypeOrderInvoiced string = "OrderInvoiced"

	// ProductContainerCommentProductContainerCommentTypeRemoved captures enum value "Removed"
	ProductContainerCommentProductContainerCommentTypeRemoved string = "Removed"

	// ProductContainerCommentProductContainerCommentTypeTemplateCreated captures enum value "TemplateCreated"
	ProductContainerCommentProductContainerCommentTypeTemplateCreated string = "TemplateCreated"

	// ProductContainerCommentProductContainerCommentTypeTemplateModifed captures enum value "TemplateModifed"
	ProductContainerCommentProductContainerCommentTypeTemplateModifed string = "TemplateModifed"

	// ProductContainerCommentProductContainerCommentTypeRequestDeclined captures enum value "RequestDeclined"
	ProductContainerCommentProductContainerCommentTypeRequestDeclined string = "RequestDeclined"

	// ProductContainerCommentProductContainerCommentTypeSystemMessage captures enum value "SystemMessage"
	ProductContainerCommentProductContainerCommentTypeSystemMessage string = "SystemMessage"
)

// prop value enum
func (m *ProductContainerComment) validateProductContainerCommentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, productContainerCommentTypeProductContainerCommentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProductContainerComment) validateProductContainerCommentType(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductContainerCommentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProductContainerCommentTypeEnum("ProductContainerCommentType", "body", m.ProductContainerCommentType); err != nil {
		return err
	}

	return nil
}

func (m *ProductContainerComment) validateTimeStamp(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeStamp) { // not required
		return nil
	}

	if err := validate.FormatOf("TimeStamp", "body", "date-time", m.TimeStamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProductContainerComment) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("User")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductContainerComment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductContainerComment) UnmarshalBinary(b []byte) error {
	var res ProductContainerComment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
