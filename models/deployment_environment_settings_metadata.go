// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeploymentEnvironmentSettingsMetadata deployment environment settings metadata
//
// swagger:model DeploymentEnvironmentSettingsMetadata
type DeploymentEnvironmentSettingsMetadata struct {

	// extension Id
	ExtensionID string `json:"ExtensionId,omitempty"`

	// metadata
	Metadata *Metadata `json:"Metadata,omitempty"`
}

// Validate validates this deployment environment settings metadata
func (m *DeploymentEnvironmentSettingsMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentEnvironmentSettingsMetadata) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentEnvironmentSettingsMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentEnvironmentSettingsMetadata) UnmarshalBinary(b []byte) error {
	var res DeploymentEnvironmentSettingsMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
