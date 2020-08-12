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

// TeamResource team resource
//
// swagger:model TeamResource
type TeamResource struct {

	// can be deleted
	CanBeDeleted bool `json:"CanBeDeleted,omitempty"`

	// can be renamed
	CanBeRenamed bool `json:"CanBeRenamed,omitempty"`

	// can change members
	CanChangeMembers bool `json:"CanChangeMembers,omitempty"`

	// can change roles
	CanChangeRoles bool `json:"CanChangeRoles,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// external security groups
	ExternalSecurityGroups []*NamedReferenceItem `json:"ExternalSecurityGroups"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// member user ids
	MemberUserIds []string `json:"MemberUserIds"`

	// name
	Name string `json:"Name,omitempty"`

	// space Id
	SpaceID string `json:"SpaceId,omitempty"`
}

// Validate validates this team resource
func (m *TeamResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalSecurityGroups(formats); err != nil {
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

func (m *TeamResource) validateExternalSecurityGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalSecurityGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalSecurityGroups); i++ {
		if swag.IsZero(m.ExternalSecurityGroups[i]) { // not required
			continue
		}

		if m.ExternalSecurityGroups[i] != nil {
			if err := m.ExternalSecurityGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExternalSecurityGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TeamResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamResource) UnmarshalBinary(b []byte) error {
	var res TeamResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}