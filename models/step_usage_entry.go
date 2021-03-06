// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StepUsageEntry step usage entry
//
// swagger:model StepUsageEntry
type StepUsageEntry struct {

	// step Id
	StepID string `json:"StepId,omitempty"`

	// step name
	StepName string `json:"StepName,omitempty"`
}

// Validate validates this step usage entry
func (m *StepUsageEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StepUsageEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StepUsageEntry) UnmarshalBinary(b []byte) error {
	var res StepUsageEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
