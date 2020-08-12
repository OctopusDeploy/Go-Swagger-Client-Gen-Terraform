// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VariablesScopedToEnvironmentResponse variables scoped to environment response
//
// swagger:model VariablesScopedToEnvironmentResponse
type VariablesScopedToEnvironmentResponse struct {

	// has unauthorized library variable set variables
	HasUnauthorizedLibraryVariableSetVariables bool `json:"HasUnauthorizedLibraryVariableSetVariables,omitempty"`

	// has unauthorized project variables
	HasUnauthorizedProjectVariables bool `json:"HasUnauthorizedProjectVariables,omitempty"`

	// variable map
	VariableMap map[string]map[string]int32 `json:"VariableMap,omitempty"`
}

// Validate validates this variables scoped to environment response
func (m *VariablesScopedToEnvironmentResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VariablesScopedToEnvironmentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariablesScopedToEnvironmentResponse) UnmarshalBinary(b []byte) error {
	var res VariablesScopedToEnvironmentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
