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

// MachineResource machine resource
//
// swagger:model MachineResource
type MachineResource struct {

	// endpoint
	Endpoint *EndpointResource `json:"Endpoint,omitempty"`

	// has latest calamari
	HasLatestCalamari bool `json:"HasLatestCalamari,omitempty"`

	// health status
	// Enum: [HasWarnings Healthy Unavailable Unhealthy Unknown]
	HealthStatus string `json:"HealthStatus,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// is disabled
	IsDisabled bool `json:"IsDisabled,omitempty"`

	// is in process
	IsInProcess bool `json:"IsInProcess,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// machine policy Id
	MachinePolicyID string `json:"MachinePolicyId,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// operating system
	OperatingSystem string `json:"OperatingSystem,omitempty"`

	// shell name
	ShellName string `json:"ShellName,omitempty"`

	// shell version
	ShellVersion string `json:"ShellVersion,omitempty"`

	// status
	// Enum: [CalamariNeedsUpgrade Disabled NeedsUpgrade Offline Online Unknown]
	Status string `json:"Status,omitempty"`

	// status summary
	StatusSummary string `json:"StatusSummary,omitempty"`

	// thumbprint
	Thumbprint string `json:"Thumbprint,omitempty"`

	// Uri
	URI string `json:"Uri,omitempty"`
}

// Validate validates this machine resource
func (m *MachineResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineResource) validateEndpoint(formats strfmt.Registry) error {

	if swag.IsZero(m.Endpoint) { // not required
		return nil
	}

	if m.Endpoint != nil {
		if err := m.Endpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Endpoint")
			}
			return err
		}
	}

	return nil
}

var machineResourceTypeHealthStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HasWarnings","Healthy","Unavailable","Unhealthy","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		machineResourceTypeHealthStatusPropEnum = append(machineResourceTypeHealthStatusPropEnum, v)
	}
}

const (

	// MachineResourceHealthStatusHasWarnings captures enum value "HasWarnings"
	MachineResourceHealthStatusHasWarnings string = "HasWarnings"

	// MachineResourceHealthStatusHealthy captures enum value "Healthy"
	MachineResourceHealthStatusHealthy string = "Healthy"

	// MachineResourceHealthStatusUnavailable captures enum value "Unavailable"
	MachineResourceHealthStatusUnavailable string = "Unavailable"

	// MachineResourceHealthStatusUnhealthy captures enum value "Unhealthy"
	MachineResourceHealthStatusUnhealthy string = "Unhealthy"

	// MachineResourceHealthStatusUnknown captures enum value "Unknown"
	MachineResourceHealthStatusUnknown string = "Unknown"
)

// prop value enum
func (m *MachineResource) validateHealthStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, machineResourceTypeHealthStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MachineResource) validateHealthStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateHealthStatusEnum("HealthStatus", "body", m.HealthStatus); err != nil {
		return err
	}

	return nil
}

func (m *MachineResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

var machineResourceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CalamariNeedsUpgrade","Disabled","NeedsUpgrade","Offline","Online","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		machineResourceTypeStatusPropEnum = append(machineResourceTypeStatusPropEnum, v)
	}
}

const (

	// MachineResourceStatusCalamariNeedsUpgrade captures enum value "CalamariNeedsUpgrade"
	MachineResourceStatusCalamariNeedsUpgrade string = "CalamariNeedsUpgrade"

	// MachineResourceStatusDisabled captures enum value "Disabled"
	MachineResourceStatusDisabled string = "Disabled"

	// MachineResourceStatusNeedsUpgrade captures enum value "NeedsUpgrade"
	MachineResourceStatusNeedsUpgrade string = "NeedsUpgrade"

	// MachineResourceStatusOffline captures enum value "Offline"
	MachineResourceStatusOffline string = "Offline"

	// MachineResourceStatusOnline captures enum value "Online"
	MachineResourceStatusOnline string = "Online"

	// MachineResourceStatusUnknown captures enum value "Unknown"
	MachineResourceStatusUnknown string = "Unknown"
)

// prop value enum
func (m *MachineResource) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, machineResourceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MachineResource) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineResource) UnmarshalBinary(b []byte) error {
	var res MachineResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
