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

// PackageFromBuiltInFeedResource package from built in feed resource
//
// swagger:model PackageFromBuiltInFeedResource
type PackageFromBuiltInFeedResource struct {

	// description
	Description string `json:"Description,omitempty"`

	// feed Id
	FeedID string `json:"FeedId,omitempty"`

	// file extension
	FileExtension string `json:"FileExtension,omitempty"`

	// hash
	Hash string `json:"Hash,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// nu get feed Id
	NuGetFeedID string `json:"NuGetFeedId,omitempty"`

	// nu get package Id
	NuGetPackageID string `json:"NuGetPackageId,omitempty"`

	// package Id
	PackageID string `json:"PackageId,omitempty"`

	// package size bytes
	PackageSizeBytes int64 `json:"PackageSizeBytes,omitempty"`

	// package version build information
	PackageVersionBuildInformation *OctopusPackageVersionBuildInformationMappedResource `json:"PackageVersionBuildInformation,omitempty"`

	// published
	// Format: date-time
	Published strfmt.DateTime `json:"Published,omitempty"`

	// release notes
	ReleaseNotes string `json:"ReleaseNotes,omitempty"`

	// summary
	Summary string `json:"Summary,omitempty"`

	// title
	Title string `json:"Title,omitempty"`

	// version
	Version string `json:"Version,omitempty"`
}

// Validate validates this package from built in feed resource
func (m *PackageFromBuiltInFeedResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageVersionBuildInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublished(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageFromBuiltInFeedResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PackageFromBuiltInFeedResource) validatePackageVersionBuildInformation(formats strfmt.Registry) error {

	if swag.IsZero(m.PackageVersionBuildInformation) { // not required
		return nil
	}

	if m.PackageVersionBuildInformation != nil {
		if err := m.PackageVersionBuildInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PackageVersionBuildInformation")
			}
			return err
		}
	}

	return nil
}

func (m *PackageFromBuiltInFeedResource) validatePublished(formats strfmt.Registry) error {

	if swag.IsZero(m.Published) { // not required
		return nil
	}

	if err := validate.FormatOf("Published", "body", "date-time", m.Published.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackageFromBuiltInFeedResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageFromBuiltInFeedResource) UnmarshalBinary(b []byte) error {
	var res PackageFromBuiltInFeedResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
