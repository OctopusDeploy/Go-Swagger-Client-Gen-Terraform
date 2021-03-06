// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EventNotificationSubscriptionFilter event notification subscription filter
//
// swagger:model EventNotificationSubscriptionFilter
type EventNotificationSubscriptionFilter struct {

	// document types
	DocumentTypes []string `json:"DocumentTypes"`

	// environments
	Environments []string `json:"Environments"`

	// event agents
	EventAgents []string `json:"EventAgents"`

	// event categories
	EventCategories []string `json:"EventCategories"`

	// event groups
	EventGroups []string `json:"EventGroups"`

	// project groups
	ProjectGroups []string `json:"ProjectGroups"`

	// projects
	Projects []string `json:"Projects"`

	// tags
	Tags []string `json:"Tags"`

	// tenants
	Tenants []string `json:"Tenants"`

	// users
	Users []string `json:"Users"`
}

// Validate validates this event notification subscription filter
func (m *EventNotificationSubscriptionFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventNotificationSubscriptionFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventNotificationSubscriptionFilter) UnmarshalBinary(b []byte) error {
	var res EventNotificationSubscriptionFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
