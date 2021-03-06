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

// ReleaseProgressionResource release progression resource
//
// swagger:model ReleaseProgressionResource
type ReleaseProgressionResource struct {

	// channel
	Channel *ChannelResource `json:"Channel,omitempty"`

	// deployments
	Deployments map[string][]DashboardItemResource `json:"Deployments,omitempty"`

	// has unresolved defect
	HasUnresolvedDefect bool `json:"HasUnresolvedDefect,omitempty"`

	// next deployments
	NextDeployments []string `json:"NextDeployments"`

	// release
	Release *ReleaseResource `json:"Release,omitempty"`

	// release retention period
	ReleaseRetentionPeriod *RetentionPeriod `json:"ReleaseRetentionPeriod,omitempty"`

	// tentacle retention period
	TentacleRetentionPeriod *RetentionPeriod `json:"TentacleRetentionPeriod,omitempty"`
}

// Validate validates this release progression resource
func (m *ReleaseProgressionResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseRetentionPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTentacleRetentionPeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseProgressionResource) validateChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	if m.Channel != nil {
		if err := m.Channel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Channel")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseProgressionResource) validateDeployments(formats strfmt.Registry) error {

	if swag.IsZero(m.Deployments) { // not required
		return nil
	}

	for k := range m.Deployments {

		if err := validate.Required("Deployments"+"."+k, "body", m.Deployments[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.Deployments[k]); i++ {

			if err := m.Deployments[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Deployments" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *ReleaseProgressionResource) validateRelease(formats strfmt.Registry) error {

	if swag.IsZero(m.Release) { // not required
		return nil
	}

	if m.Release != nil {
		if err := m.Release.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Release")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseProgressionResource) validateReleaseRetentionPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseRetentionPeriod) { // not required
		return nil
	}

	if m.ReleaseRetentionPeriod != nil {
		if err := m.ReleaseRetentionPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ReleaseRetentionPeriod")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseProgressionResource) validateTentacleRetentionPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.TentacleRetentionPeriod) { // not required
		return nil
	}

	if m.TentacleRetentionPeriod != nil {
		if err := m.TentacleRetentionPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TentacleRetentionPeriod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseProgressionResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseProgressionResource) UnmarshalBinary(b []byte) error {
	var res ReleaseProgressionResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
