// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Form form
//
// swagger:model Form
type Form struct {

	// elements
	Elements []*FormElement `json:"Elements"`

	// values
	Values map[string]string `json:"Values,omitempty"`
}

// Validate validates this form
func (m *Form) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Form) validateElements(formats strfmt.Registry) error {

	if swag.IsZero(m.Elements) { // not required
		return nil
	}

	for i := 0; i < len(m.Elements); i++ {
		if swag.IsZero(m.Elements[i]) { // not required
			continue
		}

		if m.Elements[i] != nil {
			if err := m.Elements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Elements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Form) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Form) UnmarshalBinary(b []byte) error {
	var res Form
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
