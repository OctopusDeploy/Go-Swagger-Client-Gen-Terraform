// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RunbookSnapshotUsage runbook snapshot usage
//
// swagger:model RunbookSnapshotUsage
type RunbookSnapshotUsage struct {

	// project Id
	ProjectID string `json:"ProjectId,omitempty"`

	// project name
	ProjectName string `json:"ProjectName,omitempty"`

	// runbook Id
	RunbookID string `json:"RunbookId,omitempty"`

	// runbook name
	RunbookName string `json:"RunbookName,omitempty"`

	// snapshots
	Snapshots []*RunbookSnapshotUsageEntry `json:"Snapshots"`
}

// Validate validates this runbook snapshot usage
func (m *RunbookSnapshotUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunbookSnapshotUsage) validateSnapshots(formats strfmt.Registry) error {

	if swag.IsZero(m.Snapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.Snapshots); i++ {
		if swag.IsZero(m.Snapshots[i]) { // not required
			continue
		}

		if m.Snapshots[i] != nil {
			if err := m.Snapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RunbookSnapshotUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunbookSnapshotUsage) UnmarshalBinary(b []byte) error {
	var res RunbookSnapshotUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
