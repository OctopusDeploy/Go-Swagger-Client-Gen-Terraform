// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibraryVariableSetUsageEntry library variable set usage entry
//
// swagger:model LibraryVariableSetUsageEntry
type LibraryVariableSetUsageEntry struct {

	// library variable set Id
	LibraryVariableSetID string `json:"LibraryVariableSetId,omitempty"`

	// library variable set name
	LibraryVariableSetName string `json:"LibraryVariableSetName,omitempty"`
}

// Validate validates this library variable set usage entry
func (m *LibraryVariableSetUsageEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LibraryVariableSetUsageEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LibraryVariableSetUsageEntry) UnmarshalBinary(b []byte) error {
	var res LibraryVariableSetUsageEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
