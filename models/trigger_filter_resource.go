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

// TriggerFilterResource trigger filter resource
//
// swagger:model TriggerFilterResource
type TriggerFilterResource struct {

	// filter type
	// Read Only: true
	// Enum: [ContinuousDailySchedule CronExpressionSchedule DailySchedule DaysPerMonthSchedule DaysPerWeekSchedule MachineFilter OnceDailySchedule]
	FilterType string `json:"FilterType,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`
}

// Validate validates this trigger filter resource
func (m *TriggerFilterResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterType(formats); err != nil {
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

var triggerFilterResourceTypeFilterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ContinuousDailySchedule","CronExpressionSchedule","DailySchedule","DaysPerMonthSchedule","DaysPerWeekSchedule","MachineFilter","OnceDailySchedule"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		triggerFilterResourceTypeFilterTypePropEnum = append(triggerFilterResourceTypeFilterTypePropEnum, v)
	}
}

const (

	// TriggerFilterResourceFilterTypeContinuousDailySchedule captures enum value "ContinuousDailySchedule"
	TriggerFilterResourceFilterTypeContinuousDailySchedule string = "ContinuousDailySchedule"

	// TriggerFilterResourceFilterTypeCronExpressionSchedule captures enum value "CronExpressionSchedule"
	TriggerFilterResourceFilterTypeCronExpressionSchedule string = "CronExpressionSchedule"

	// TriggerFilterResourceFilterTypeDailySchedule captures enum value "DailySchedule"
	TriggerFilterResourceFilterTypeDailySchedule string = "DailySchedule"

	// TriggerFilterResourceFilterTypeDaysPerMonthSchedule captures enum value "DaysPerMonthSchedule"
	TriggerFilterResourceFilterTypeDaysPerMonthSchedule string = "DaysPerMonthSchedule"

	// TriggerFilterResourceFilterTypeDaysPerWeekSchedule captures enum value "DaysPerWeekSchedule"
	TriggerFilterResourceFilterTypeDaysPerWeekSchedule string = "DaysPerWeekSchedule"

	// TriggerFilterResourceFilterTypeMachineFilter captures enum value "MachineFilter"
	TriggerFilterResourceFilterTypeMachineFilter string = "MachineFilter"

	// TriggerFilterResourceFilterTypeOnceDailySchedule captures enum value "OnceDailySchedule"
	TriggerFilterResourceFilterTypeOnceDailySchedule string = "OnceDailySchedule"
)

// prop value enum
func (m *TriggerFilterResource) validateFilterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, triggerFilterResourceTypeFilterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TriggerFilterResource) validateFilterType(formats strfmt.Registry) error {

	if swag.IsZero(m.FilterType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFilterTypeEnum("FilterType", "body", m.FilterType); err != nil {
		return err
	}

	return nil
}

func (m *TriggerFilterResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TriggerFilterResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TriggerFilterResource) UnmarshalBinary(b []byte) error {
	var res TriggerFilterResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}