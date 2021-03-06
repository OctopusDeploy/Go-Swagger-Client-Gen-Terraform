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

// PackageVersionResource package version resource
//
// swagger:model PackageVersionResource
type PackageVersionResource struct {

	// feed Id
	FeedID string `json:"FeedId,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// package Id
	PackageID string `json:"PackageId,omitempty"`

	// published
	// Format: date-time
	Published strfmt.DateTime `json:"Published,omitempty"`

	// release notes
	ReleaseNotes string `json:"ReleaseNotes,omitempty"`

	// size bytes
	SizeBytes int64 `json:"SizeBytes,omitempty"`

	// title
	Title string `json:"Title,omitempty"`

	// version
	Version string `json:"Version,omitempty"`
}

// Validate validates this package version resource
func (m *PackageVersionResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublished(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageVersionResource) validatePublished(formats strfmt.Registry) error {

	if swag.IsZero(m.Published) { // not required
		return nil
	}

	if err := validate.FormatOf("Published", "body", "date-time", m.Published.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackageVersionResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageVersionResource) UnmarshalBinary(b []byte) error {
	var res PackageVersionResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
