// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnvironmentSummaryResource environment summary resource
//
// swagger:model EnvironmentSummaryResource
type EnvironmentSummaryResource struct {

	// environment
	Environment *EnvironmentResource `json:"Environment,omitempty"`

	// machine endpoint summaries
	MachineEndpointSummaries map[string]int32 `json:"MachineEndpointSummaries,omitempty"`

	// machine health status summaries
	MachineHealthStatusSummaries map[string]int32 `json:"MachineHealthStatusSummaries,omitempty"`

	// machine ids for calamari upgrade
	MachineIdsForCalamariUpgrade []string `json:"MachineIdsForCalamariUpgrade"`

	// machine ids for tentacle upgrade
	MachineIdsForTentacleUpgrade []string `json:"MachineIdsForTentacleUpgrade"`

	// machine tenant summaries
	MachineTenantSummaries map[string]int32 `json:"MachineTenantSummaries,omitempty"`

	// machine tenant tag summaries
	MachineTenantTagSummaries map[string]int32 `json:"MachineTenantTagSummaries,omitempty"`

	// tentacle upgrades required
	TentacleUpgradesRequired bool `json:"TentacleUpgradesRequired,omitempty"`

	// total disabled machines
	TotalDisabledMachines int32 `json:"TotalDisabledMachines,omitempty"`

	// total machines
	TotalMachines int32 `json:"TotalMachines,omitempty"`
}

// Validate validates this environment summary resource
func (m *EnvironmentSummaryResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentSummaryResource) validateEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Environment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentSummaryResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentSummaryResource) UnmarshalBinary(b []byte) error {
	var res EnvironmentSummaryResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
