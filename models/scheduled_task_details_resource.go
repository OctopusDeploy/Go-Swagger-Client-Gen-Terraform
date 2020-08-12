// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScheduledTaskDetailsResource scheduled task details resource
//
// swagger:model ScheduledTaskDetailsResource
type ScheduledTaskDetailsResource struct {

	// activity log
	ActivityLog *ActivityLogTreeNode `json:"ActivityLog,omitempty"`

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

// Validate validates this scheduled task details resource
func (m *ScheduledTaskDetailsResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityLog(formats); err != nil {
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

func (m *ScheduledTaskDetailsResource) validateActivityLog(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityLog) { // not required
		return nil
	}

	if m.ActivityLog != nil {
		if err := m.ActivityLog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActivityLog")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduledTaskDetailsResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduledTaskDetailsResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduledTaskDetailsResource) UnmarshalBinary(b []byte) error {
	var res ScheduledTaskDetailsResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}