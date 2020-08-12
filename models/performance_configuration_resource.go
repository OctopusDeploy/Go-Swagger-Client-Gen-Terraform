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

// PerformanceConfigurationResource performance configuration resource
//
// swagger:model PerformanceConfigurationResource
type PerformanceConfigurationResource struct {

	// default dashboard render mode
	// Enum: [VirtualizeColumns VirtualizeRowsAndColumns]
	DefaultDashboardRenderMode string `json:"DefaultDashboardRenderMode,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`
}

// Validate validates this performance configuration resource
func (m *PerformanceConfigurationResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultDashboardRenderMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var performanceConfigurationResourceTypeDefaultDashboardRenderModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VirtualizeColumns","VirtualizeRowsAndColumns"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceConfigurationResourceTypeDefaultDashboardRenderModePropEnum = append(performanceConfigurationResourceTypeDefaultDashboardRenderModePropEnum, v)
	}
}

const (

	// PerformanceConfigurationResourceDefaultDashboardRenderModeVirtualizeColumns captures enum value "VirtualizeColumns"
	PerformanceConfigurationResourceDefaultDashboardRenderModeVirtualizeColumns string = "VirtualizeColumns"

	// PerformanceConfigurationResourceDefaultDashboardRenderModeVirtualizeRowsAndColumns captures enum value "VirtualizeRowsAndColumns"
	PerformanceConfigurationResourceDefaultDashboardRenderModeVirtualizeRowsAndColumns string = "VirtualizeRowsAndColumns"
)

// prop value enum
func (m *PerformanceConfigurationResource) validateDefaultDashboardRenderModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceConfigurationResourceTypeDefaultDashboardRenderModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceConfigurationResource) validateDefaultDashboardRenderMode(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultDashboardRenderMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateDefaultDashboardRenderModeEnum("DefaultDashboardRenderMode", "body", m.DefaultDashboardRenderMode); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceConfigurationResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceConfigurationResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceConfigurationResource) UnmarshalBinary(b []byte) error {
	var res PerformanceConfigurationResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
