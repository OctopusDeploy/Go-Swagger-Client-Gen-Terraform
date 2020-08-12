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

// MigrationPartialExportResource migration partial export resource
//
// swagger:model MigrationPartialExportResource
type MigrationPartialExportResource struct {

	// destination Api key
	DestinationAPIKey string `json:"DestinationApiKey,omitempty"`

	// destination package feed
	DestinationPackageFeed string `json:"DestinationPackageFeed,omitempty"`

	// destination package feed space Id
	DestinationPackageFeedSpaceID string `json:"DestinationPackageFeedSpaceId,omitempty"`

	// encrypt package
	EncryptPackage bool `json:"EncryptPackage,omitempty"`

	// failure callback Uri
	FailureCallbackURI string `json:"FailureCallbackUri,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// ignore certificates
	IgnoreCertificates bool `json:"IgnoreCertificates,omitempty"`

	// ignore deployments
	IgnoreDeployments bool `json:"IgnoreDeployments,omitempty"`

	// ignore machines
	IgnoreMachines bool `json:"IgnoreMachines,omitempty"`

	// ignore tenants
	IgnoreTenants bool `json:"IgnoreTenants,omitempty"`

	// include task logs
	IncludeTaskLogs bool `json:"IncludeTaskLogs,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// package Id
	PackageID string `json:"PackageId,omitempty"`

	// package version
	PackageVersion string `json:"PackageVersion,omitempty"`

	// password
	Password string `json:"Password,omitempty"`

	// projects
	Projects []string `json:"Projects"`

	// space Id
	SpaceID string `json:"SpaceId,omitempty"`

	// success callback Uri
	SuccessCallbackURI string `json:"SuccessCallbackUri,omitempty"`

	// task Id
	TaskID string `json:"TaskId,omitempty"`
}

// Validate validates this migration partial export resource
func (m *MigrationPartialExportResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrationPartialExportResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MigrationPartialExportResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrationPartialExportResource) UnmarshalBinary(b []byte) error {
	var res MigrationPartialExportResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
