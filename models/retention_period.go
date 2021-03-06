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

// RetentionPeriod retention period
//
// swagger:model RetentionPeriod
type RetentionPeriod struct {

	// quantity to keep
	// Read Only: true
	QuantityToKeep int32 `json:"QuantityToKeep,omitempty"`

	// should keep forever
	// Read Only: true
	ShouldKeepForever *bool `json:"ShouldKeepForever,omitempty"`

	// unit
	// Read Only: true
	// Enum: [Days Items]
	Unit string `json:"Unit,omitempty"`
}

// Validate validates this retention period
func (m *RetentionPeriod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var retentionPeriodTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Days","Items"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		retentionPeriodTypeUnitPropEnum = append(retentionPeriodTypeUnitPropEnum, v)
	}
}

const (

	// RetentionPeriodUnitDays captures enum value "Days"
	RetentionPeriodUnitDays string = "Days"

	// RetentionPeriodUnitItems captures enum value "Items"
	RetentionPeriodUnitItems string = "Items"
)

// prop value enum
func (m *RetentionPeriod) validateUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, retentionPeriodTypeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RetentionPeriod) validateUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	// value enum
	if err := m.validateUnitEnum("Unit", "body", m.Unit); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RetentionPeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionPeriod) UnmarshalBinary(b []byte) error {
	var res RetentionPeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
