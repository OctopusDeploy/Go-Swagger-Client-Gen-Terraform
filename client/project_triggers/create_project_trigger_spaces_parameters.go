// Code generated by go-swagger; DO NOT EDIT.

package project_triggers

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

// NewCreateProjectTriggerSpacesParams creates a new CreateProjectTriggerSpacesParams object
// with the default values initialized.
func NewCreateProjectTriggerSpacesParams() *CreateProjectTriggerSpacesParams {
	var ()
	return &CreateProjectTriggerSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectTriggerSpacesParamsWithTimeout creates a new CreateProjectTriggerSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateProjectTriggerSpacesParamsWithTimeout(timeout time.Duration) *CreateProjectTriggerSpacesParams {
	var ()
	return &CreateProjectTriggerSpacesParams{

		timeout: timeout,
	}
}

// NewCreateProjectTriggerSpacesParamsWithContext creates a new CreateProjectTriggerSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateProjectTriggerSpacesParamsWithContext(ctx context.Context) *CreateProjectTriggerSpacesParams {
	var ()
	return &CreateProjectTriggerSpacesParams{

		Context: ctx,
	}
}

// NewCreateProjectTriggerSpacesParamsWithHTTPClient creates a new CreateProjectTriggerSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateProjectTriggerSpacesParamsWithHTTPClient(client *http.Client) *CreateProjectTriggerSpacesParams {
	var ()
	return &CreateProjectTriggerSpacesParams{
		HTTPClient: client,
	}
}

/*CreateProjectTriggerSpacesParams contains all the parameters to send to the API endpoint
for the create project trigger spaces operation typically these are written to a http.Request
*/
type CreateProjectTriggerSpacesParams struct {

	/*ProjectTriggerResource
	  The ProjectTriggerResource resource to create

	*/
	ProjectTriggerResource *models.ProjectTriggerResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create project trigger spaces params
func (o *CreateProjectTriggerSpacesParams) WithTimeout(timeout time.Duration) *CreateProjectTriggerSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project trigger spaces params
func (o *CreateProjectTriggerSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project trigger spaces params
func (o *CreateProjectTriggerSpacesParams) WithContext(ctx context.Context) *CreateProjectTriggerSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project trigger spaces params
func (o *CreateProjectTriggerSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project trigger spaces params
func (o *CreateProjectTriggerSpacesParams) WithHTTPClient(client *http.Client) *CreateProjectTriggerSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project trigger spaces params
func (o *CreateProjectTriggerSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectTriggerResource adds the projectTriggerResource to the create project trigger spaces params
func (o *CreateProjectTriggerSpacesParams) WithProjectTriggerResource(projectTriggerResource *models.ProjectTriggerResource) *CreateProjectTriggerSpacesParams {
	o.SetProjectTriggerResource(projectTriggerResource)
	return o
}

// SetProjectTriggerResource adds the projectTriggerResource to the create project trigger spaces params
func (o *CreateProjectTriggerSpacesParams) SetProjectTriggerResource(projectTriggerResource *models.ProjectTriggerResource) {
	o.ProjectTriggerResource = projectTriggerResource
}

// WithBaseSpaceID adds the baseSpaceID to the create project trigger spaces params
func (o *CreateProjectTriggerSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateProjectTriggerSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create project trigger spaces params
func (o *CreateProjectTriggerSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProjectTriggerSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ProjectTriggerResource != nil {
		if err := r.SetBodyParam(o.ProjectTriggerResource); err != nil {
			return err
		}
	}

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
