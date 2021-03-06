// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkItemLink work item link
//
// swagger:model WorkItemLink
type WorkItemLink struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// link Url
	LinkURL string `json:"LinkUrl,omitempty"`

	// source
	Source string `json:"Source,omitempty"`
}

// Validate validates this work item link
func (m *WorkItemLink) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkItemLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkItemLink) UnmarshalBinary(b []byte) error {
	var res WorkItemLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
