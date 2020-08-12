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

// SchedulerStatusResource scheduler status resource
//
// swagger:model SchedulerStatusResource
type SchedulerStatusResource struct {

	// Id
	ID string `json:"Id,omitempty"`

	// is running
	IsRunning bool `json:"IsRunning,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// task status
	TaskStatus []*ScheduledTaskStatusResource `json:"TaskStatus"`
}

// Validate validates this scheduler status resource
func (m *SchedulerStatusResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerStatusResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchedulerStatusResource) validateTaskStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskStatus) { // not required
		return nil
	}

	for i := 0; i < len(m.TaskStatus); i++ {
		if swag.IsZero(m.TaskStatus[i]) { // not required
			continue
		}

		if m.TaskStatus[i] != nil {
			if err := m.TaskStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TaskStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulerStatusResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulerStatusResource) UnmarshalBinary(b []byte) error {
	var res SchedulerStatusResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}