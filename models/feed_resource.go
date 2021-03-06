// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FeedResource feed resource
//
// swagger:model FeedResource
type FeedResource struct {

	// feed type
	// Read Only: true
	// Enum: [AwsElasticContainerRegistry BuiltIn Docker GitHub Helm Maven None NuGet OctopusProject]
	FeedType string `json:"FeedType,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// package acquisition location options
	PackageAcquisitionLocationOptions []string `json:"PackageAcquisitionLocationOptions"`

	// space Id
	SpaceID string `json:"SpaceId,omitempty"`
}

// Validate validates this feed resource
func (m *FeedResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeedType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageAcquisitionLocationOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var feedResourceTypeFeedTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AwsElasticContainerRegistry","BuiltIn","Docker","GitHub","Helm","Maven","None","NuGet","OctopusProject"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		feedResourceTypeFeedTypePropEnum = append(feedResourceTypeFeedTypePropEnum, v)
	}
}

const (

	// FeedResourceFeedTypeAwsElasticContainerRegistry captures enum value "AwsElasticContainerRegistry"
	FeedResourceFeedTypeAwsElasticContainerRegistry string = "AwsElasticContainerRegistry"

	// FeedResourceFeedTypeBuiltIn captures enum value "BuiltIn"
	FeedResourceFeedTypeBuiltIn string = "BuiltIn"

	// FeedResourceFeedTypeDocker captures enum value "Docker"
	FeedResourceFeedTypeDocker string = "Docker"

	// FeedResourceFeedTypeGitHub captures enum value "GitHub"
	FeedResourceFeedTypeGitHub string = "GitHub"

	// FeedResourceFeedTypeHelm captures enum value "Helm"
	FeedResourceFeedTypeHelm string = "Helm"

	// FeedResourceFeedTypeMaven captures enum value "Maven"
	FeedResourceFeedTypeMaven string = "Maven"

	// FeedResourceFeedTypeNone captures enum value "None"
	FeedResourceFeedTypeNone string = "None"

	// FeedResourceFeedTypeNuGet captures enum value "NuGet"
	FeedResourceFeedTypeNuGet string = "NuGet"

	// FeedResourceFeedTypeOctopusProject captures enum value "OctopusProject"
	FeedResourceFeedTypeOctopusProject string = "OctopusProject"
)

// prop value enum
func (m *FeedResource) validateFeedTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, feedResourceTypeFeedTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FeedResource) validateFeedType(formats strfmt.Registry) error {

	if swag.IsZero(m.FeedType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFeedTypeEnum("FeedType", "body", m.FeedType); err != nil {
		return err
	}

	return nil
}

func (m *FeedResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

var feedResourcePackageAcquisitionLocationOptionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ExecutionTarget","NotAcquired","Server"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		feedResourcePackageAcquisitionLocationOptionsItemsEnum = append(feedResourcePackageAcquisitionLocationOptionsItemsEnum, v)
	}
}

func (m *FeedResource) validatePackageAcquisitionLocationOptionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, feedResourcePackageAcquisitionLocationOptionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FeedResource) validatePackageAcquisitionLocationOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.PackageAcquisitionLocationOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.PackageAcquisitionLocationOptions); i++ {

		// value enum
		if err := m.validatePackageAcquisitionLocationOptionsItemsEnum("PackageAcquisitionLocationOptions"+"."+strconv.Itoa(i), "body", m.PackageAcquisitionLocationOptions[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeedResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedResource) UnmarshalBinary(b []byte) error {
	var res FeedResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
