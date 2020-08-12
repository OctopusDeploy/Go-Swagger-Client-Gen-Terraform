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

// ChannelVersionRuleResource channel version rule resource
//
// swagger:model ChannelVersionRuleResource
type ChannelVersionRuleResource struct {

	// action packages
	ActionPackages []*DeploymentActionPackageResource `json:"ActionPackages"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// tag
	Tag string `json:"Tag,omitempty"`

	// version range
	VersionRange string `json:"VersionRange,omitempty"`
}

// Validate validates this channel version rule resource
func (m *ChannelVersionRuleResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionPackages(formats); err != nil {
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

func (m *ChannelVersionRuleResource) validateActionPackages(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionPackages) { // not required
		return nil
	}

	for i := 0; i < len(m.ActionPackages); i++ {
		if swag.IsZero(m.ActionPackages[i]) { // not required
			continue
		}

		if m.ActionPackages[i] != nil {
			if err := m.ActionPackages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ActionPackages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChannelVersionRuleResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelVersionRuleResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelVersionRuleResource) UnmarshalBinary(b []byte) error {
	var res ChannelVersionRuleResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}