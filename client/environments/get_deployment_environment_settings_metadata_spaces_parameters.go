// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetDeploymentEnvironmentSettingsMetadataSpacesParams creates a new GetDeploymentEnvironmentSettingsMetadataSpacesParams object
// with the default values initialized.
func NewGetDeploymentEnvironmentSettingsMetadataSpacesParams() *GetDeploymentEnvironmentSettingsMetadataSpacesParams {
	var ()
	return &GetDeploymentEnvironmentSettingsMetadataSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentEnvironmentSettingsMetadataSpacesParamsWithTimeout creates a new GetDeploymentEnvironmentSettingsMetadataSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentEnvironmentSettingsMetadataSpacesParamsWithTimeout(timeout time.Duration) *GetDeploymentEnvironmentSettingsMetadataSpacesParams {
	var ()
	return &GetDeploymentEnvironmentSettingsMetadataSpacesParams{

		timeout: timeout,
	}
}

// NewGetDeploymentEnvironmentSettingsMetadataSpacesParamsWithContext creates a new GetDeploymentEnvironmentSettingsMetadataSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentEnvironmentSettingsMetadataSpacesParamsWithContext(ctx context.Context) *GetDeploymentEnvironmentSettingsMetadataSpacesParams {
	var ()
	return &GetDeploymentEnvironmentSettingsMetadataSpacesParams{

		Context: ctx,
	}
}

// NewGetDeploymentEnvironmentSettingsMetadataSpacesParamsWithHTTPClient creates a new GetDeploymentEnvironmentSettingsMetadataSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentEnvironmentSettingsMetadataSpacesParamsWithHTTPClient(client *http.Client) *GetDeploymentEnvironmentSettingsMetadataSpacesParams {
	var ()
	return &GetDeploymentEnvironmentSettingsMetadataSpacesParams{
		HTTPClient: client,
	}
}

/*GetDeploymentEnvironmentSettingsMetadataSpacesParams contains all the parameters to send to the API endpoint
for the get deployment environment settings metadata spaces operation typically these are written to a http.Request
*/
type GetDeploymentEnvironmentSettingsMetadataSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment environment settings metadata spaces params
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesParams) WithTimeout(timeout time.Duration) *GetDeploymentEnvironmentSettingsMetadataSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment environment settings metadata spaces params
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment environment settings metadata spaces params
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesParams) WithContext(ctx context.Context) *GetDeploymentEnvironmentSettingsMetadataSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment environment settings metadata spaces params
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment environment settings metadata spaces params
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesParams) WithHTTPClient(client *http.Client) *GetDeploymentEnvironmentSettingsMetadataSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment environment settings metadata spaces params
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get deployment environment settings metadata spaces params
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetDeploymentEnvironmentSettingsMetadataSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get deployment environment settings metadata spaces params
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get deployment environment settings metadata spaces params
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesParams) WithID(id string) *GetDeploymentEnvironmentSettingsMetadataSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get deployment environment settings metadata spaces params
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
