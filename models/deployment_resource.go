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

// DeploymentResource deployment resource
//
// swagger:model DeploymentResource
type DeploymentResource struct {

	// changes
	Changes []*ReleaseChanges `json:"Changes"`

	// changes markdown
	ChangesMarkdown string `json:"ChangesMarkdown,omitempty"`

	// channel Id
	ChannelID string `json:"ChannelId,omitempty"`

	// comments
	Comments string `json:"Comments,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"Created,omitempty"`

	// deployed by
	DeployedBy string `json:"DeployedBy,omitempty"`

	// deployed by Id
	DeployedByID string `json:"DeployedById,omitempty"`

	// deployed to machine ids
	DeployedToMachineIds []string `json:"DeployedToMachineIds"`

	// deployment process Id
	DeploymentProcessID string `json:"DeploymentProcessId,omitempty"`

	// environment Id
	// Required: true
	EnvironmentID *string `json:"EnvironmentId"`

	// excluded machine ids
	ExcludedMachineIds []string `json:"ExcludedMachineIds"`

	// failure encountered
	FailureEncountered bool `json:"FailureEncountered,omitempty"`

	// force package download
	ForcePackageDownload bool `json:"ForcePackageDownload,omitempty"`

	// force package redeployment
	ForcePackageRedeployment bool `json:"ForcePackageRedeployment,omitempty"`

	// form values
	FormValues map[string]string `json:"FormValues,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// manifest variable set Id
	ManifestVariableSetID string `json:"ManifestVariableSetId,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// project Id
	ProjectID string `json:"ProjectId,omitempty"`

	// queue time
	// Format: date-time
	QueueTime strfmt.DateTime `json:"QueueTime,omitempty"`

	// queue time expiry
	// Format: date-time
	QueueTimeExpiry strfmt.DateTime `json:"QueueTimeExpiry,omitempty"`

	// release Id
	// Required: true
	ReleaseID *string `json:"ReleaseId"`

	// skip actions
	SkipActions []string `json:"SkipActions"`

	// space Id
	SpaceID string `json:"SpaceId,omitempty"`

	// specific machine ids
	SpecificMachineIds []string `json:"SpecificMachineIds"`

	// task Id
	TaskID string `json:"TaskId,omitempty"`

	// tenant Id
	TenantID string `json:"TenantId,omitempty"`

	// tentacle retention period
	TentacleRetentionPeriod *RetentionPeriod `json:"TentacleRetentionPeriod,omitempty"`

	// use guided failure
	UseGuidedFailure bool `json:"UseGuidedFailure,omitempty"`
}

// Validate validates this deployment resource
func (m *DeploymentResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueTimeExpiry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTentacleRetentionPeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentResource) validateChanges(formats strfmt.Registry) error {

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

func (m *DeploymentResource) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("Created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResource) validateEnvironmentID(formats strfmt.Registry) error {

	if err := validate.Required("EnvironmentId", "body", m.EnvironmentID); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResource) validateQueueTime(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueTime) { // not required
		return nil
	}

	if err := validate.FormatOf("QueueTime", "body", "date-time", m.QueueTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResource) validateQueueTimeExpiry(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueTimeExpiry) { // not required
		return nil
	}

	if err := validate.FormatOf("QueueTimeExpiry", "body", "date-time", m.QueueTimeExpiry.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResource) validateReleaseID(formats strfmt.Registry) error {

	if err := validate.Required("ReleaseId", "body", m.ReleaseID); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResource) validateTentacleRetentionPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.TentacleRetentionPeriod) { // not required
		return nil
	}

	if m.TentacleRetentionPeriod != nil {
		if err := m.TentacleRetentionPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TentacleRetentionPeriod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentResource) UnmarshalBinary(b []byte) error {
	var res DeploymentResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}