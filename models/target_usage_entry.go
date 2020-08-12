// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TargetUsageEntry target usage entry
//
// swagger:model TargetUsageEntry
type TargetUsageEntry struct {

	// target Id
	TargetID string `json:"TargetId,omitempty"`

	// target name
	TargetName string `json:"TargetName,omitempty"`
}

// Validate validates this target usage entry
func (m *TargetUsageEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TargetUsageEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TargetUsageEntry) UnmarshalBinary(b []byte) error {
	var res TargetUsageEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
