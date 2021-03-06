// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeploymentActionPackageResource deployment action package resource
//
// swagger:model DeploymentActionPackageResource
type DeploymentActionPackageResource struct {

	// deployment action
	DeploymentAction string `json:"DeploymentAction,omitempty"`

	// package reference
	PackageReference string `json:"PackageReference,omitempty"`
}

// Validate validates this deployment action package resource
func (m *DeploymentActionPackageResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentActionPackageResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentActionPackageResource) UnmarshalBinary(b []byte) error {
	var res DeploymentActionPackageResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
