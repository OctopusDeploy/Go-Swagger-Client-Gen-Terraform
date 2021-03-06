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

// DashboardResource dashboard resource
//
// swagger:model DashboardResource
type DashboardResource struct {

	// environments
	Environments []*DashboardEnvironmentResource `json:"Environments"`

	// Id
	ID string `json:"Id,omitempty"`

	// is filtered
	IsFiltered bool `json:"IsFiltered,omitempty"`

	// items
	Items []*DashboardItemResource `json:"Items"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// previous items
	PreviousItems []*DashboardItemResource `json:"PreviousItems"`

	// project groups
	ProjectGroups []*DashboardProjectGroupResource `json:"ProjectGroups"`

	// project limit
	ProjectLimit int32 `json:"ProjectLimit,omitempty"`

	// projects
	Projects []*DashboardProjectResource `json:"Projects"`

	// tenants
	Tenants []*DashboardTenantResource `json:"Tenants"`
}

// Validate validates this dashboard resource
func (m *DashboardResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviousItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenants(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardResource) validateEnvironments(formats strfmt.Registry) error {

	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	for i := 0; i < len(m.Environments); i++ {
		if swag.IsZero(m.Environments[i]) { // not required
			continue
		}

		if m.Environments[i] != nil {
			if err := m.Environments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DashboardResource) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DashboardResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardResource) validatePreviousItems(formats strfmt.Registry) error {

	if swag.IsZero(m.PreviousItems) { // not required
		return nil
	}

	for i := 0; i < len(m.PreviousItems); i++ {
		if swag.IsZero(m.PreviousItems[i]) { // not required
			continue
		}

		if m.PreviousItems[i] != nil {
			if err := m.PreviousItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PreviousItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DashboardResource) validateProjectGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.ProjectGroups); i++ {
		if swag.IsZero(m.ProjectGroups[i]) { // not required
			continue
		}

		if m.ProjectGroups[i] != nil {
			if err := m.ProjectGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ProjectGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DashboardResource) validateProjects(formats strfmt.Registry) error {

	if swag.IsZero(m.Projects) { // not required
		return nil
	}

	for i := 0; i < len(m.Projects); i++ {
		if swag.IsZero(m.Projects[i]) { // not required
			continue
		}

		if m.Projects[i] != nil {
			if err := m.Projects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DashboardResource) validateTenants(formats strfmt.Registry) error {

	if swag.IsZero(m.Tenants) { // not required
		return nil
	}

	for i := 0; i < len(m.Tenants); i++ {
		if swag.IsZero(m.Tenants[i]) { // not required
			continue
		}

		if m.Tenants[i] != nil {
			if err := m.Tenants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Tenants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardResource) UnmarshalBinary(b []byte) error {
	var res DashboardResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
