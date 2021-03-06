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

// DeploymentPromomotionTenant deployment promomotion tenant
//
// swagger:model DeploymentPromomotionTenant
type DeploymentPromomotionTenant struct {

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// promote to
	PromoteTo []*DeploymentPromotionTarget `json:"PromoteTo"`
}

// Validate validates this deployment promomotion tenant
func (m *DeploymentPromomotionTenant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePromoteTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentPromomotionTenant) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentPromomotionTenant) validatePromoteTo(formats strfmt.Registry) error {

	if swag.IsZero(m.PromoteTo) { // not required
		return nil
	}

	for i := 0; i < len(m.PromoteTo); i++ {
		if swag.IsZero(m.PromoteTo[i]) { // not required
			continue
		}

		if m.PromoteTo[i] != nil {
			if err := m.PromoteTo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PromoteTo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentPromomotionTenant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentPromomotionTenant) UnmarshalBinary(b []byte) error {
	var res DeploymentPromomotionTenant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
