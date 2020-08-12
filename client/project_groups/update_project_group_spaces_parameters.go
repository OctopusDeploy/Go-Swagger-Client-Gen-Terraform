// Code generated by go-swagger; DO NOT EDIT.

package project_groups

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

// NewUpdateProjectGroupSpacesParams creates a new UpdateProjectGroupSpacesParams object
// with the default values initialized.
func NewUpdateProjectGroupSpacesParams() *UpdateProjectGroupSpacesParams {
	var ()
	return &UpdateProjectGroupSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProjectGroupSpacesParamsWithTimeout creates a new UpdateProjectGroupSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateProjectGroupSpacesParamsWithTimeout(timeout time.Duration) *UpdateProjectGroupSpacesParams {
	var ()
	return &UpdateProjectGroupSpacesParams{

		timeout: timeout,
	}
}

// NewUpdateProjectGroupSpacesParamsWithContext creates a new UpdateProjectGroupSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateProjectGroupSpacesParamsWithContext(ctx context.Context) *UpdateProjectGroupSpacesParams {
	var ()
	return &UpdateProjectGroupSpacesParams{

		Context: ctx,
	}
}

// NewUpdateProjectGroupSpacesParamsWithHTTPClient creates a new UpdateProjectGroupSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateProjectGroupSpacesParamsWithHTTPClient(client *http.Client) *UpdateProjectGroupSpacesParams {
	var ()
	return &UpdateProjectGroupSpacesParams{
		HTTPClient: client,
	}
}

/*UpdateProjectGroupSpacesParams contains all the parameters to send to the API endpoint
for the update project group spaces operation typically these are written to a http.Request
*/
type UpdateProjectGroupSpacesParams struct {

	/*ProjectGroupResource
	  The ProjectGroupResource resource to create

	*/
	ProjectGroupResource *models.ProjectGroupResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the ProjectGroupResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) WithTimeout(timeout time.Duration) *UpdateProjectGroupSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) WithContext(ctx context.Context) *UpdateProjectGroupSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) WithHTTPClient(client *http.Client) *UpdateProjectGroupSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectGroupResource adds the projectGroupResource to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) WithProjectGroupResource(projectGroupResource *models.ProjectGroupResource) *UpdateProjectGroupSpacesParams {
	o.SetProjectGroupResource(projectGroupResource)
	return o
}

// SetProjectGroupResource adds the projectGroupResource to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) SetProjectGroupResource(projectGroupResource *models.ProjectGroupResource) {
	o.ProjectGroupResource = projectGroupResource
}

// WithBaseSpaceID adds the baseSpaceID to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) WithBaseSpaceID(baseSpaceID string) *UpdateProjectGroupSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) WithID(id string) *UpdateProjectGroupSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update project group spaces params
func (o *UpdateProjectGroupSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProjectGroupSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ProjectGroupResource != nil {
		if err := r.SetBodyParam(o.ProjectGroupResource); err != nil {
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
