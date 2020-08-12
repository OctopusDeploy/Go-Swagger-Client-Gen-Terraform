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

// OctopusPackageMetadataMappedResource octopus package metadata mapped resource
//
// swagger:model OctopusPackageMetadataMappedResource
type OctopusPackageMetadataMappedResource struct {

	// build environment
	BuildEnvironment string `json:"BuildEnvironment,omitempty"`

	// build number
	BuildNumber string `json:"BuildNumber,omitempty"`

	// build Url
	BuildURL string `json:"BuildUrl,omitempty"`

	// commits
	Commits []*CommitDetails `json:"Commits"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"Created,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// incomplete data warning
	IncompleteDataWarning string `json:"IncompleteDataWarning,omitempty"`

	// issue tracker name
	IssueTrackerName string `json:"IssueTrackerName,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// package Id
	PackageID string `json:"PackageId,omitempty"`

	// vcs commit number
	VcsCommitNumber string `json:"VcsCommitNumber,omitempty"`

	// vcs commit Url
	VcsCommitURL string `json:"VcsCommitUrl,omitempty"`

	// vcs root
	VcsRoot string `json:"VcsRoot,omitempty"`

	// vcs type
	VcsType string `json:"VcsType,omitempty"`

	// version
	Version string `json:"Version,omitempty"`

	// work items
	WorkItems []*WorkItemLink `json:"WorkItems"`
}

// Validate validates this octopus package metadata mapped resource
func (m *OctopusPackageMetadataMappedResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OctopusPackageMetadataMappedResource) validateCommits(formats strfmt.Registry) error {

	if swag.IsZero(m.Commits) { // not required
		return nil
	}

	for i := 0; i < len(m.Commits); i++ {
		if swag.IsZero(m.Commits[i]) { // not required
			continue
		}

		if m.Commits[i] != nil {
			if err := m.Commits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Commits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OctopusPackageMetadataMappedResource) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("Created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OctopusPackageMetadataMappedResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OctopusPackageMetadataMappedResource) validateWorkItems(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkItems) { // not required
		return nil
	}

	for i := 0; i < len(m.WorkItems); i++ {
		if swag.IsZero(m.WorkItems[i]) { // not required
			continue
		}

		if m.WorkItems[i] != nil {
			if err := m.WorkItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("WorkItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OctopusPackageMetadataMappedResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OctopusPackageMetadataMappedResource) UnmarshalBinary(b []byte) error {
	var res OctopusPackageMetadataMappedResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
