// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LicenseMessageResource license message resource
//
// swagger:model LicenseMessageResource
type LicenseMessageResource struct {

	// disposition
	// Enum: [Error Information Warning]
	Disposition string `json:"Disposition,omitempty"`

	// message
	Message string `json:"Message,omitempty"`
}

// Validate validates this license message resource
func (m *LicenseMessageResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisposition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var licenseMessageResourceTypeDispositionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Error","Information","Warning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		licenseMessageResourceTypeDispositionPropEnum = append(licenseMessageResourceTypeDispositionPropEnum, v)
	}
}

const (

	// LicenseMessageResourceDispositionError captures enum value "Error"
	LicenseMessageResourceDispositionError string = "Error"

	// LicenseMessageResourceDispositionInformation captures enum value "Information"
	LicenseMessageResourceDispositionInformation string = "Information"

	// LicenseMessageResourceDispositionWarning captures enum value "Warning"
	LicenseMessageResourceDispositionWarning string = "Warning"
)

// prop value enum
func (m *LicenseMessageResource) validateDispositionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, licenseMessageResourceTypeDispositionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LicenseMessageResource) validateDisposition(formats strfmt.Registry) error {

	if swag.IsZero(m.Disposition) { // not required
		return nil
	}

	// value enum
	if err := m.validateDispositionEnum("Disposition", "body", m.Disposition); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicenseMessageResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseMessageResource) UnmarshalBinary(b []byte) error {
	var res LicenseMessageResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
