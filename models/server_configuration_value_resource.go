// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServerConfigurationValueResource server configuration value resource
//
// swagger:model ServerConfigurationValueResource
type ServerConfigurationValueResource struct {

	// description
	Description string `json:"Description,omitempty"`

	// key
	Key string `json:"Key,omitempty"`

	// value
	Value interface{} `json:"Value,omitempty"`
}

// Validate validates this server configuration value resource
func (m *ServerConfigurationValueResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerConfigurationValueResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerConfigurationValueResource) UnmarshalBinary(b []byte) error {
	var res ServerConfigurationValueResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
