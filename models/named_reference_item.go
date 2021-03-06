// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NamedReferenceItem named reference item
//
// swagger:model NamedReferenceItem
type NamedReferenceItem struct {

	// display Id and name
	DisplayIDAndName bool `json:"DisplayIdAndName,omitempty"`

	// display name
	DisplayName string `json:"DisplayName,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`
}

// Validate validates this named reference item
func (m *NamedReferenceItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NamedReferenceItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamedReferenceItem) UnmarshalBinary(b []byte) error {
	var res NamedReferenceItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
