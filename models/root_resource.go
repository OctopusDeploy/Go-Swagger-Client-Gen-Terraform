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

// RootResource root resource
//
// swagger:model RootResource
type RootResource struct {

	// Api version
	APIVersion string `json:"ApiVersion,omitempty"`

	// application
	Application string `json:"Application,omitempty"`

	// has long term support
	// Read Only: true
	HasLongTermSupport *bool `json:"HasLongTermSupport,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// installation Id
	// Format: uuid
	InstallationID strfmt.UUID `json:"InstallationId,omitempty"`

	// is early access program
	IsEarlyAccessProgram *bool `json:"IsEarlyAccessProgram,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// version
	Version string `json:"Version,omitempty"`
}

// Validate validates this root resource
func (m *RootResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstallationID(formats); err != nil {
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

func (m *RootResource) validateInstallationID(formats strfmt.Registry) error {

	if swag.IsZero(m.InstallationID) { // not required
		return nil
	}

	if err := validate.FormatOf("InstallationId", "body", "uuid", m.InstallationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RootResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RootResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RootResource) UnmarshalBinary(b []byte) error {
	var res RootResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
