// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventNotificationSubscription event notification subscription
//
// swagger:model EventNotificationSubscription
type EventNotificationSubscription struct {

	// email digest last processed
	// Format: date-time
	EmailDigestLastProcessed strfmt.DateTime `json:"EmailDigestLastProcessed,omitempty"`

	// email digest last processed event auto Id
	EmailDigestLastProcessedEventAutoID int64 `json:"EmailDigestLastProcessedEventAutoId,omitempty"`

	// email frequency period
	EmailFrequencyPeriod string `json:"EmailFrequencyPeriod,omitempty"`

	// email priority
	// Enum: [High Low Normal]
	EmailPriority string `json:"EmailPriority,omitempty"`

	// email show dates in time zone Id
	EmailShowDatesInTimeZoneID string `json:"EmailShowDatesInTimeZoneId,omitempty"`

	// email teams
	EmailTeams []string `json:"EmailTeams"`

	// filter
	Filter *EventNotificationSubscriptionFilter `json:"Filter,omitempty"`

	// webhook header key
	WebhookHeaderKey string `json:"WebhookHeaderKey,omitempty"`

	// webhook header value
	WebhookHeaderValue string `json:"WebhookHeaderValue,omitempty"`

	// webhook last processed
	// Format: date-time
	WebhookLastProcessed strfmt.DateTime `json:"WebhookLastProcessed,omitempty"`

	// webhook last processed event auto Id
	WebhookLastProcessedEventAutoID int64 `json:"WebhookLastProcessedEventAutoId,omitempty"`

	// webhook teams
	WebhookTeams []string `json:"WebhookTeams"`

	// webhook timeout
	WebhookTimeout string `json:"WebhookTimeout,omitempty"`

	// webhook URI
	WebhookURI string `json:"WebhookURI,omitempty"`
}

// Validate validates this event notification subscription
func (m *EventNotificationSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailDigestLastProcessed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookLastProcessed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventNotificationSubscription) validateEmailDigestLastProcessed(formats strfmt.Registry) error {

	if swag.IsZero(m.EmailDigestLastProcessed) { // not required
		return nil
	}

	if err := validate.FormatOf("EmailDigestLastProcessed", "body", "date-time", m.EmailDigestLastProcessed.String(), formats); err != nil {
		return err
	}

	return nil
}

var eventNotificationSubscriptionTypeEmailPriorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["High","Low","Normal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventNotificationSubscriptionTypeEmailPriorityPropEnum = append(eventNotificationSubscriptionTypeEmailPriorityPropEnum, v)
	}
}

const (

	// EventNotificationSubscriptionEmailPriorityHigh captures enum value "High"
	EventNotificationSubscriptionEmailPriorityHigh string = "High"

	// EventNotificationSubscriptionEmailPriorityLow captures enum value "Low"
	EventNotificationSubscriptionEmailPriorityLow string = "Low"

	// EventNotificationSubscriptionEmailPriorityNormal captures enum value "Normal"
	EventNotificationSubscriptionEmailPriorityNormal string = "Normal"
)

// prop value enum
func (m *EventNotificationSubscription) validateEmailPriorityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, eventNotificationSubscriptionTypeEmailPriorityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EventNotificationSubscription) validateEmailPriority(formats strfmt.Registry) error {

	if swag.IsZero(m.EmailPriority) { // not required
		return nil
	}

	// value enum
	if err := m.validateEmailPriorityEnum("EmailPriority", "body", m.EmailPriority); err != nil {
		return err
	}

	return nil
}

func (m *EventNotificationSubscription) validateFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Filter")
			}
			return err
		}
	}

	return nil
}

func (m *EventNotificationSubscription) validateWebhookLastProcessed(formats strfmt.Registry) error {

	if swag.IsZero(m.WebhookLastProcessed) { // not required
		return nil
	}

	if err := validate.FormatOf("WebhookLastProcessed", "body", "date-time", m.WebhookLastProcessed.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventNotificationSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventNotificationSubscription) UnmarshalBinary(b []byte) error {
	var res EventNotificationSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
