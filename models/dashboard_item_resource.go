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

// DashboardItemResource dashboard item resource
//
// swagger:model DashboardItemResource
type DashboardItemResource struct {

	// channel Id
	ChannelID string `json:"ChannelId,omitempty"`

	// completed time
	// Format: date-time
	CompletedTime strfmt.DateTime `json:"CompletedTime,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"Created,omitempty"`

	// deployment Id
	DeploymentID string `json:"DeploymentId,omitempty"`

	// duration
	Duration string `json:"Duration,omitempty"`

	// environment Id
	EnvironmentID string `json:"EnvironmentId,omitempty"`

	// error message
	ErrorMessage string `json:"ErrorMessage,omitempty"`

	// has pending interruptions
	HasPendingInterruptions bool `json:"HasPendingInterruptions,omitempty"`

	// has warnings or errors
	HasWarningsOrErrors bool `json:"HasWarningsOrErrors,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// is completed
	IsCompleted bool `json:"IsCompleted,omitempty"`

	// is current
	IsCurrent bool `json:"IsCurrent,omitempty"`

	// is previous
	IsPrevious bool `json:"IsPrevious,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// project Id
	ProjectID string `json:"ProjectId,omitempty"`

	// queue time
	// Format: date-time
	QueueTime strfmt.DateTime `json:"QueueTime,omitempty"`

	// release Id
	ReleaseID string `json:"ReleaseId,omitempty"`

	// release version
	ReleaseVersion string `json:"ReleaseVersion,omitempty"`

	// start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"StartTime,omitempty"`

	// state
	// Enum: [Canceled Cancelling Executing Failed Queued Success TimedOut]
	State string `json:"State,omitempty"`

	// task Id
	TaskID string `json:"TaskId,omitempty"`

	// tenant Id
	TenantID string `json:"TenantId,omitempty"`
}

// Validate validates this dashboard item resource
func (m *DashboardItemResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardItemResource) validateCompletedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CompletedTime", "body", "date-time", m.CompletedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardItemResource) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("Created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardItemResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardItemResource) validateQueueTime(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueTime) { // not required
		return nil
	}

	if err := validate.FormatOf("QueueTime", "body", "date-time", m.QueueTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardItemResource) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var dashboardItemResourceTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Canceled","Cancelling","Executing","Failed","Queued","Success","TimedOut"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dashboardItemResourceTypeStatePropEnum = append(dashboardItemResourceTypeStatePropEnum, v)
	}
}

const (

	// DashboardItemResourceStateCanceled captures enum value "Canceled"
	DashboardItemResourceStateCanceled string = "Canceled"

	// DashboardItemResourceStateCancelling captures enum value "Cancelling"
	DashboardItemResourceStateCancelling string = "Cancelling"

	// DashboardItemResourceStateExecuting captures enum value "Executing"
	DashboardItemResourceStateExecuting string = "Executing"

	// DashboardItemResourceStateFailed captures enum value "Failed"
	DashboardItemResourceStateFailed string = "Failed"

	// DashboardItemResourceStateQueued captures enum value "Queued"
	DashboardItemResourceStateQueued string = "Queued"

	// DashboardItemResourceStateSuccess captures enum value "Success"
	DashboardItemResourceStateSuccess string = "Success"

	// DashboardItemResourceStateTimedOut captures enum value "TimedOut"
	DashboardItemResourceStateTimedOut string = "TimedOut"
)

// prop value enum
func (m *DashboardItemResource) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dashboardItemResourceTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DashboardItemResource) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("State", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardItemResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardItemResource) UnmarshalBinary(b []byte) error {
	var res DashboardItemResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
