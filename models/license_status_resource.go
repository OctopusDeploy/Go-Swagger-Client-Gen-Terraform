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

// LicenseStatusResource license status resource
//
// swagger:model LicenseStatusResource
type LicenseStatusResource struct {

	// compliance summary
	ComplianceSummary string `json:"ComplianceSummary,omitempty"`

	// days to effective expiry date
	DaysToEffectiveExpiryDate int32 `json:"DaysToEffectiveExpiryDate,omitempty"`

	// does expiry block key activities
	DoesExpiryBlockKeyActivities bool `json:"DoesExpiryBlockKeyActivities,omitempty"`

	// effective cluster task limit
	EffectiveClusterTaskLimit int32 `json:"EffectiveClusterTaskLimit,omitempty"`

	// effective expiry date
	// Format: date
	EffectiveExpiryDate strfmt.Date `json:"EffectiveExpiryDate,omitempty"`

	// effective node task limit
	EffectiveNodeTaskLimit int32 `json:"EffectiveNodeTaskLimit,omitempty"`

	// hosting environment
	HostingEnvironment string `json:"HostingEnvironment,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// is cluster task limit controlled by license
	IsClusterTaskLimitControlledByLicense bool `json:"IsClusterTaskLimitControlledByLicense,omitempty"`

	// is compliant
	IsCompliant bool `json:"IsCompliant,omitempty"`

	// is node task limit controlled by license
	IsNodeTaskLimitControlledByLicense bool `json:"IsNodeTaskLimitControlledByLicense,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// limits
	Limits []*LicenseLimitStatusResource `json:"Limits"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// messages
	Messages []*LicenseMessageResource `json:"Messages"`

	// permissions mode
	// Enum: [Full Restricted Unspecified]
	PermissionsMode string `json:"PermissionsMode,omitempty"`
}

// Validate validates this license status resource
func (m *LicenseStatusResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffectiveExpiryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissionsMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicenseStatusResource) validateEffectiveExpiryDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EffectiveExpiryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("EffectiveExpiryDate", "body", "date", m.EffectiveExpiryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LicenseStatusResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LicenseStatusResource) validateLimits(formats strfmt.Registry) error {

	if swag.IsZero(m.Limits) { // not required
		return nil
	}

	for i := 0; i < len(m.Limits); i++ {
		if swag.IsZero(m.Limits[i]) { // not required
			continue
		}

		if m.Limits[i] != nil {
			if err := m.Limits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Limits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LicenseStatusResource) validateMessages(formats strfmt.Registry) error {

	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	for i := 0; i < len(m.Messages); i++ {
		if swag.IsZero(m.Messages[i]) { // not required
			continue
		}

		if m.Messages[i] != nil {
			if err := m.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var licenseStatusResourceTypePermissionsModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Full","Restricted","Unspecified"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		licenseStatusResourceTypePermissionsModePropEnum = append(licenseStatusResourceTypePermissionsModePropEnum, v)
	}
}

const (

	// LicenseStatusResourcePermissionsModeFull captures enum value "Full"
	LicenseStatusResourcePermissionsModeFull string = "Full"

	// LicenseStatusResourcePermissionsModeRestricted captures enum value "Restricted"
	LicenseStatusResourcePermissionsModeRestricted string = "Restricted"

	// LicenseStatusResourcePermissionsModeUnspecified captures enum value "Unspecified"
	LicenseStatusResourcePermissionsModeUnspecified string = "Unspecified"
)

// prop value enum
func (m *LicenseStatusResource) validatePermissionsModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, licenseStatusResourceTypePermissionsModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LicenseStatusResource) validatePermissionsMode(formats strfmt.Registry) error {

	if swag.IsZero(m.PermissionsMode) { // not required
		return nil
	}

	// value enum
	if err := m.validatePermissionsModeEnum("PermissionsMode", "body", m.PermissionsMode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicenseStatusResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseStatusResource) UnmarshalBinary(b []byte) error {
	var res LicenseStatusResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
