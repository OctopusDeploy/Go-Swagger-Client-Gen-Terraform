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

	"models"
)

// NewUpdateEnvironmentSpacesParams creates a new UpdateEnvironmentSpacesParams object
// with the default values initialized.
func NewUpdateEnvironmentSpacesParams() *UpdateEnvironmentSpacesParams {
	var ()
	return &UpdateEnvironmentSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEnvironmentSpacesParamsWithTimeout creates a new UpdateEnvironmentSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateEnvironmentSpacesParamsWithTimeout(timeout time.Duration) *UpdateEnvironmentSpacesParams {
	var ()
	return &UpdateEnvironmentSpacesParams{

		timeout: timeout,
	}
}

// NewUpdateEnvironmentSpacesParamsWithContext creates a new UpdateEnvironmentSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateEnvironmentSpacesParamsWithContext(ctx context.Context) *UpdateEnvironmentSpacesParams {
	var ()
	return &UpdateEnvironmentSpacesParams{

		Context: ctx,
	}
}

// NewUpdateEnvironmentSpacesParamsWithHTTPClient creates a new UpdateEnvironmentSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateEnvironmentSpacesParamsWithHTTPClient(client *http.Client) *UpdateEnvironmentSpacesParams {
	var ()
	return &UpdateEnvironmentSpacesParams{
		HTTPClient: client,
	}
}

/*UpdateEnvironmentSpacesParams contains all the parameters to send to the API endpoint
for the update environment spaces operation typically these are written to a http.Request
*/
type UpdateEnvironmentSpacesParams struct {

	/*EnvironmentResource
	  The EnvironmentResource resource to create

	*/
	EnvironmentResource *models.EnvironmentResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the EnvironmentResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) WithTimeout(timeout time.Duration) *UpdateEnvironmentSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) WithContext(ctx context.Context) *UpdateEnvironmentSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) WithHTTPClient(client *http.Client) *UpdateEnvironmentSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentResource adds the environmentResource to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) WithEnvironmentResource(environmentResource *models.EnvironmentResource) *UpdateEnvironmentSpacesParams {
	o.SetEnvironmentResource(environmentResource)
	return o
}

// SetEnvironmentResource adds the environmentResource to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) SetEnvironmentResource(environmentResource *models.EnvironmentResource) {
	o.EnvironmentResource = environmentResource
}

// WithBaseSpaceID adds the baseSpaceID to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) WithBaseSpaceID(baseSpaceID string) *UpdateEnvironmentSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) WithID(id string) *UpdateEnvironmentSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update environment spaces params
func (o *UpdateEnvironmentSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEnvironmentSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnvironmentResource != nil {
		if err := r.SetBodyParam(o.EnvironmentResource); err != nil {
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
