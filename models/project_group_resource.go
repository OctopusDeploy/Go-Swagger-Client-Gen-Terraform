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

// ProjectGroupResource project group resource
//
// swagger:model ProjectGroupResource
type ProjectGroupResource struct {

	// description
	Description string `json:"Description,omitempty"`

	// environment ids
	EnvironmentIds []string `json:"EnvironmentIds"`

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

	// retention policy Id
	RetentionPolicyID string `json:"RetentionPolicyId,omitempty"`

	// space Id
	SpaceID string `json:"SpaceId,omitempty"`
}

// Validate validates this project group resource
func (m *ProjectGroupResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectGroupResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectGroupResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectGroupResource) UnmarshalBinary(b []byte) error {
	var res ProjectGroupResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
