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

// NumericReportData numeric report data
//
// swagger:model NumericReportData
type NumericReportData struct {

	// labels
	Labels []string `json:"Labels"`

	// series
	Series []*NumericReportSeries `json:"Series"`
}

// Validate validates this numeric report data
func (m *NumericReportData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NumericReportData) validateSeries(formats strfmt.Registry) error {

	if swag.IsZero(m.Series) { // not required
		return nil
	}

	for i := 0; i < len(m.Series); i++ {
		if swag.IsZero(m.Series[i]) { // not required
			continue
		}

		if m.Series[i] != nil {
			if err := m.Series[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Series" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NumericReportData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NumericReportData) UnmarshalBinary(b []byte) error {
	var res NumericReportData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
