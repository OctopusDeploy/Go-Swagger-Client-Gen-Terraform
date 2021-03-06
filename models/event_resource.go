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

// EventResource event resource
//
// swagger:model EventResource
type EventResource struct {

	// category
	Category string `json:"Category,omitempty"`

	// change details
	ChangeDetails *ChangeDetails `json:"ChangeDetails,omitempty"`

	// comments
	Comments string `json:"Comments,omitempty"`

	// details
	Details string `json:"Details,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// identity established with
	IdentityEstablishedWith string `json:"IdentityEstablishedWith,omitempty"`

	// is service
	IsService bool `json:"IsService,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message Html
	MessageHTML string `json:"MessageHtml,omitempty"`

	// message references
	MessageReferences []*EventReference `json:"MessageReferences"`

	// occurred
	// Format: date-time
	Occurred strfmt.DateTime `json:"Occurred,omitempty"`

	// related document ids
	RelatedDocumentIds []string `json:"RelatedDocumentIds"`

	// space Id
	SpaceID string `json:"SpaceId,omitempty"`

	// user agent
	UserAgent string `json:"UserAgent,omitempty"`

	// user Id
	UserID string `json:"UserId,omitempty"`

	// username
	Username string `json:"Username,omitempty"`
}

// Validate validates this event resource
func (m *EventResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangeDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageReferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOccurred(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventResource) validateChangeDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.ChangeDetails) { // not required
		return nil
	}

	if m.ChangeDetails != nil {
		if err := m.ChangeDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ChangeDetails")
			}
			return err
		}
	}

	return nil
}

func (m *EventResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventResource) validateMessageReferences(formats strfmt.Registry) error {

	if swag.IsZero(m.MessageReferences) { // not required
		return nil
	}

	for i := 0; i < len(m.MessageReferences); i++ {
		if swag.IsZero(m.MessageReferences[i]) { // not required
			continue
		}

		if m.MessageReferences[i] != nil {
			if err := m.MessageReferences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("MessageReferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EventResource) validateOccurred(formats strfmt.Registry) error {

	if swag.IsZero(m.Occurred) { // not required
		return nil
	}

	if err := validate.FormatOf("Occurred", "body", "date-time", m.Occurred.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventResource) UnmarshalBinary(b []byte) error {
	var res EventResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
