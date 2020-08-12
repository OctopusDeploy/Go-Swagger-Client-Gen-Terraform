// Code generated by go-swagger; DO NOT EDIT.

package artifacts

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

	"models"
)

// NewUpdateArtifactSpacesParams creates a new UpdateArtifactSpacesParams object
// with the default values initialized.
func NewUpdateArtifactSpacesParams() *UpdateArtifactSpacesParams {
	var ()
	return &UpdateArtifactSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateArtifactSpacesParamsWithTimeout creates a new UpdateArtifactSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateArtifactSpacesParamsWithTimeout(timeout time.Duration) *UpdateArtifactSpacesParams {
	var ()
	return &UpdateArtifactSpacesParams{

		timeout: timeout,
	}
}

// NewUpdateArtifactSpacesParamsWithContext creates a new UpdateArtifactSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateArtifactSpacesParamsWithContext(ctx context.Context) *UpdateArtifactSpacesParams {
	var ()
	return &UpdateArtifactSpacesParams{

		Context: ctx,
	}
}

// NewUpdateArtifactSpacesParamsWithHTTPClient creates a new UpdateArtifactSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateArtifactSpacesParamsWithHTTPClient(client *http.Client) *UpdateArtifactSpacesParams {
	var ()
	return &UpdateArtifactSpacesParams{
		HTTPClient: client,
	}
}

/*UpdateArtifactSpacesParams contains all the parameters to send to the API endpoint
for the update artifact spaces operation typically these are written to a http.Request
*/
type UpdateArtifactSpacesParams struct {

	/*ArtifactResource
	  The ArtifactResource resource to create

	*/
	ArtifactResource *models.ArtifactResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the ArtifactResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) WithTimeout(timeout time.Duration) *UpdateArtifactSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) WithContext(ctx context.Context) *UpdateArtifactSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) WithHTTPClient(client *http.Client) *UpdateArtifactSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactResource adds the artifactResource to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) WithArtifactResource(artifactResource *models.ArtifactResource) *UpdateArtifactSpacesParams {
	o.SetArtifactResource(artifactResource)
	return o
}

// SetArtifactResource adds the artifactResource to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) SetArtifactResource(artifactResource *models.ArtifactResource) {
	o.ArtifactResource = artifactResource
}

// WithBaseSpaceID adds the baseSpaceID to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) WithBaseSpaceID(baseSpaceID string) *UpdateArtifactSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) WithID(id string) *UpdateArtifactSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update artifact spaces params
func (o *UpdateArtifactSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateArtifactSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ArtifactResource != nil {
		if err := r.SetBodyParam(o.ArtifactResource); err != nil {
			return err
		}
	}

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
