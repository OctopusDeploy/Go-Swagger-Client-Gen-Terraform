// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MachineConnectivityPolicy machine connectivity policy
//
// swagger:model MachineConnectivityPolicy
type MachineConnectivityPolicy struct {

	// machine connectivity behavior
	// Enum: [ExpectedToBeOnline MayBeOfflineAndCanBeSkipped]
	MachineConnectivityBehavior string `json:"MachineConnectivityBehavior,omitempty"`
}

// Validate validates this machine connectivity policy
func (m *MachineConnectivityPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachineConnectivityBehavior(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var machineConnectivityPolicyTypeMachineConnectivityBehaviorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ExpectedToBeOnline","MayBeOfflineAndCanBeSkipped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		machineConnectivityPolicyTypeMachineConnectivityBehaviorPropEnum = append(machineConnectivityPolicyTypeMachineConnectivityBehaviorPropEnum, v)
	}
}

const (

	// MachineConnectivityPolicyMachineConnectivityBehaviorExpectedToBeOnline captures enum value "ExpectedToBeOnline"
	MachineConnectivityPolicyMachineConnectivityBehaviorExpectedToBeOnline string = "ExpectedToBeOnline"

	// MachineConnectivityPolicyMachineConnectivityBehaviorMayBeOfflineAndCanBeSkipped captures enum value "MayBeOfflineAndCanBeSkipped"
	MachineConnectivityPolicyMachineConnectivityBehaviorMayBeOfflineAndCanBeSkipped string = "MayBeOfflineAndCanBeSkipped"
)

// prop value enum
func (m *MachineConnectivityPolicy) validateMachineConnectivityBehaviorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, machineConnectivityPolicyTypeMachineConnectivityBehaviorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MachineConnectivityPolicy) validateMachineConnectivityBehavior(formats strfmt.Registry) error {

	if swag.IsZero(m.MachineConnectivityBehavior) { // not required
		return nil
	}

	// value enum
	if err := m.validateMachineConnectivityBehaviorEnum("MachineConnectivityBehavior", "body", m.MachineConnectivityBehavior); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineConnectivityPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineConnectivityPolicy) UnmarshalBinary(b []byte) error {
	var res MachineConnectivityPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
