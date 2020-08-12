// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VersioningStrategyResource versioning strategy resource
//
// swagger:model VersioningStrategyResource
type VersioningStrategyResource struct {

	// donor package
	DonorPackage *DeploymentActionPackageResource `json:"DonorPackage,omitempty"`

	// template
	Template string `json:"Template,omitempty"`
}

// Validate validates this versioning strategy resource
func (m *VersioningStrategyResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDonorPackage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersioningStrategyResource) validateDonorPackage(formats strfmt.Registry) error {

	if swag.IsZero(m.DonorPackage) { // not required
		return nil
	}

	if m.DonorPackage != nil {
		if err := m.DonorPackage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DonorPackage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersioningStrategyResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersioningStrategyResource) UnmarshalBinary(b []byte) error {
	var res VersioningStrategyResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
