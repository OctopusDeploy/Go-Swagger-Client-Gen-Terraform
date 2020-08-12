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
)

// NewUpdateArtifactContentSpacesParams creates a new UpdateArtifactContentSpacesParams object
// with the default values initialized.
func NewUpdateArtifactContentSpacesParams() *UpdateArtifactContentSpacesParams {
	var ()
	return &UpdateArtifactContentSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateArtifactContentSpacesParamsWithTimeout creates a new UpdateArtifactContentSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateArtifactContentSpacesParamsWithTimeout(timeout time.Duration) *UpdateArtifactContentSpacesParams {
	var ()
	return &UpdateArtifactContentSpacesParams{

		timeout: timeout,
	}
}

// NewUpdateArtifactContentSpacesParamsWithContext creates a new UpdateArtifactContentSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateArtifactContentSpacesParamsWithContext(ctx context.Context) *UpdateArtifactContentSpacesParams {
	var ()
	return &UpdateArtifactContentSpacesParams{

		Context: ctx,
	}
}

// NewUpdateArtifactContentSpacesParamsWithHTTPClient creates a new UpdateArtifactContentSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateArtifactContentSpacesParamsWithHTTPClient(client *http.Client) *UpdateArtifactContentSpacesParams {
	var ()
	return &UpdateArtifactContentSpacesParams{
		HTTPClient: client,
	}
}

/*UpdateArtifactContentSpacesParams contains all the parameters to send to the API endpoint
for the update artifact content spaces operation typically these are written to a http.Request
*/
type UpdateArtifactContentSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the artifact

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update artifact content spaces params
func (o *UpdateArtifactContentSpacesParams) WithTimeout(timeout time.Duration) *UpdateArtifactContentSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update artifact content spaces params
func (o *UpdateArtifactContentSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update artifact content spaces params
func (o *UpdateArtifactContentSpacesParams) WithContext(ctx context.Context) *UpdateArtifactContentSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update artifact content spaces params
func (o *UpdateArtifactContentSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update artifact content spaces params
func (o *UpdateArtifactContentSpacesParams) WithHTTPClient(client *http.Client) *UpdateArtifactContentSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update artifact content spaces params
func (o *UpdateArtifactContentSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the update artifact content spaces params
func (o *UpdateArtifactContentSpacesParams) WithBaseSpaceID(baseSpaceID string) *UpdateArtifactContentSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the update artifact content spaces params
func (o *UpdateArtifactContentSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the update artifact content spaces params
func (o *UpdateArtifactContentSpacesParams) WithID(id string) *UpdateArtifactContentSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update artifact content spaces params
func (o *UpdateArtifactContentSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateArtifactContentSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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