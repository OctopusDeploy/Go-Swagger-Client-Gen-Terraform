// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MachineConnectionStatus machine connection status
//
// swagger:model MachineConnectionStatus
type MachineConnectionStatus struct {

	// current tentacle version
	CurrentTentacleVersion string `json:"CurrentTentacleVersion,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last checked
	// Format: date-time
	LastChecked strfmt.DateTime `json:"LastChecked,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// logs
	Logs []*ActivityLogElement `json:"Logs"`

	// machine Id
	MachineID string `json:"MachineId,omitempty"`

	// status
	Status string `json:"Status,omitempty"`
}

// Validate validates this machine connection status
func (m *MachineConnectionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastChecked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineConnectionStatus) validateLastChecked(formats strfmt.Registry) error {

	if swag.IsZero(m.LastChecked) { // not required
		return nil
	}

	if err := validate.FormatOf("LastChecked", "body", "date-time", m.LastChecked.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MachineConnectionStatus) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MachineConnectionStatus) validateLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.Logs) { // not required
		return nil
	}

	for i := 0; i < len(m.Logs); i++ {
		if swag.IsZero(m.Logs[i]) { // not required
			continue
		}

		if m.Logs[i] != nil {
			if err := m.Logs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineConnectionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineConnectionStatus) UnmarshalBinary(b []byte) error {
	var res MachineConnectionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
