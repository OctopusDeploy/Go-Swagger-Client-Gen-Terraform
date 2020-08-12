// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PackageReference package reference
//
// swagger:model PackageReference
type PackageReference struct {

	// acquisition location
	AcquisitionLocation string `json:"AcquisitionLocation,omitempty"`

	// feed Id
	FeedID string `json:"FeedId,omitempty"`

	// Id
	// Read Only: true
	ID string `json:"Id,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// package Id
	PackageID string `json:"PackageId,omitempty"`

	// properties
	Properties map[string]string `json:"Properties,omitempty"`
}

// Validate validates this package reference
func (m *PackageReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackageReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageReference) UnmarshalBinary(b []byte) error {
	var res PackageReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}