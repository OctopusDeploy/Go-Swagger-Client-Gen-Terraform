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

// DeploymentPreviewResource deployment preview resource
//
// swagger:model DeploymentPreviewResource
type DeploymentPreviewResource struct {

	// changes
	Changes []*ReleaseChanges `json:"Changes"`

	// changes markdown
	ChangesMarkdown string `json:"ChangesMarkdown,omitempty"`

	// form
	Form *Form `json:"Form,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// steps to execute
	StepsToExecute []*DeploymentTemplateStep `json:"StepsToExecute"`

	// use guided failure mode by default
	UseGuidedFailureModeByDefault bool `json:"UseGuidedFailureModeByDefault,omitempty"`
}

// Validate validates this deployment preview resource
func (m *DeploymentPreviewResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStepsToExecute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentPreviewResource) validateChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.Changes) { // not required
		return nil
	}

	for i := 0; i < len(m.Changes); i++ {
		if swag.IsZero(m.Changes[i]) { // not required
			continue
		}

		if m.Changes[i] != nil {
			if err := m.Changes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentPreviewResource) validateForm(formats strfmt.Registry) error {

	if swag.IsZero(m.Form) { // not required
		return nil
	}

	if m.Form != nil {
		if err := m.Form.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Form")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentPreviewResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentPreviewResource) validateStepsToExecute(formats strfmt.Registry) error {

	if swag.IsZero(m.StepsToExecute) { // not required
		return nil
	}

	for i := 0; i < len(m.StepsToExecute); i++ {
		if swag.IsZero(m.StepsToExecute[i]) { // not required
			continue
		}

		if m.StepsToExecute[i] != nil {
			if err := m.StepsToExecute[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StepsToExecute" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentPreviewResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentPreviewResource) UnmarshalBinary(b []byte) error {
	var res DeploymentPreviewResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
