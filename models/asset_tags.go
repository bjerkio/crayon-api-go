// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AssetTags asset tags
//
// swagger:model AssetTags
type AssetTags struct {

	// asset Id
	AssetID int32 `json:"AssetId,omitempty"`

	// cost center
	CostCenter string `json:"CostCenter,omitempty"`

	// custom
	Custom string `json:"Custom,omitempty"`

	// department
	Department string `json:"Department,omitempty"`

	// owner
	Owner string `json:"Owner,omitempty"`

	// project
	Project string `json:"Project,omitempty"`
}

// Validate validates this asset tags
func (m *AssetTags) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssetTags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetTags) UnmarshalBinary(b []byte) error {
	var res AssetTags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
