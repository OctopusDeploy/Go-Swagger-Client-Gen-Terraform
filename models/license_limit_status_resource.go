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

// LicenseLimitStatusResource license limit status resource
//
// swagger:model LicenseLimitStatusResource
type LicenseLimitStatusResource struct {

	// current usage
	CurrentUsage int32 `json:"CurrentUsage,omitempty"`

	// disposition
	// Enum: [Error Information Warning]
	Disposition string `json:"Disposition,omitempty"`

	// effective limit
	EffectiveLimit int32 `json:"EffectiveLimit,omitempty"`

	// effective limit description
	EffectiveLimitDescription string `json:"EffectiveLimitDescription,omitempty"`

	// is unlimited
	IsUnlimited bool `json:"IsUnlimited,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this license limit status resource
func (m *LicenseLimitStatusResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisposition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var licenseLimitStatusResourceTypeDispositionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Error","Information","Warning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		licenseLimitStatusResourceTypeDispositionPropEnum = append(licenseLimitStatusResourceTypeDispositionPropEnum, v)
	}
}

const (

	// LicenseLimitStatusResourceDispositionError captures enum value "Error"
	LicenseLimitStatusResourceDispositionError string = "Error"

	// LicenseLimitStatusResourceDispositionInformation captures enum value "Information"
	LicenseLimitStatusResourceDispositionInformation string = "Information"

	// LicenseLimitStatusResourceDispositionWarning captures enum value "Warning"
	LicenseLimitStatusResourceDispositionWarning string = "Warning"
)

// prop value enum
func (m *LicenseLimitStatusResource) validateDispositionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, licenseLimitStatusResourceTypeDispositionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LicenseLimitStatusResource) validateDisposition(formats strfmt.Registry) error {

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
func (m *LicenseLimitStatusResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseLimitStatusResource) UnmarshalBinary(b []byte) error {
	var res LicenseLimitStatusResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}