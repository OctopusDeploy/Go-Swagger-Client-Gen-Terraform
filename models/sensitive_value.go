// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SensitiveValue sensitive value
//
// swagger:model SensitiveValue
type SensitiveValue struct {

	// has value
	HasValue bool `json:"HasValue,omitempty"`

	// new value
	NewValue string `json:"NewValue,omitempty"`
}

// Validate validates this sensitive value
func (m *SensitiveValue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SensitiveValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SensitiveValue) UnmarshalBinary(b []byte) error {
	var res SensitiveValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
