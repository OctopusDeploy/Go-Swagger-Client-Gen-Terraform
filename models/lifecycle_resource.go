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

// LifecycleResource lifecycle resource
//
// swagger:model LifecycleResource
type LifecycleResource struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// phases
	// Read Only: true
	Phases []*PhaseResource `json:"Phases"`

	// release retention policy
	ReleaseRetentionPolicy *RetentionPeriod `json:"ReleaseRetentionPolicy,omitempty"`

	// space Id
	SpaceID string `json:"SpaceId,omitempty"`

	// tentacle retention policy
	TentacleRetentionPolicy *RetentionPeriod `json:"TentacleRetentionPolicy,omitempty"`
}

// Validate validates this lifecycle resource
func (m *LifecycleResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseRetentionPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTentacleRetentionPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LifecycleResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LifecycleResource) validatePhases(formats strfmt.Registry) error {

	if swag.IsZero(m.Phases) { // not required
		return nil
	}

	for i := 0; i < len(m.Phases); i++ {
		if swag.IsZero(m.Phases[i]) { // not required
			continue
		}

		if m.Phases[i] != nil {
			if err := m.Phases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Phases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LifecycleResource) validateReleaseRetentionPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseRetentionPolicy) { // not required
		return nil
	}

	if m.ReleaseRetentionPolicy != nil {
		if err := m.ReleaseRetentionPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ReleaseRetentionPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *LifecycleResource) validateTentacleRetentionPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.TentacleRetentionPolicy) { // not required
		return nil
	}

	if m.TentacleRetentionPolicy != nil {
		if err := m.TentacleRetentionPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TentacleRetentionPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LifecycleResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LifecycleResource) UnmarshalBinary(b []byte) error {
	var res LifecycleResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}