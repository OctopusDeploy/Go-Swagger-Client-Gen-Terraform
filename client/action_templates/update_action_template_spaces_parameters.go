// Code generated by go-swagger; DO NOT EDIT.

package action_templates

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

// NewUpdateActionTemplateSpacesParams creates a new UpdateActionTemplateSpacesParams object
// with the default values initialized.
func NewUpdateActionTemplateSpacesParams() *UpdateActionTemplateSpacesParams {
	var ()
	return &UpdateActionTemplateSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateActionTemplateSpacesParamsWithTimeout creates a new UpdateActionTemplateSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateActionTemplateSpacesParamsWithTimeout(timeout time.Duration) *UpdateActionTemplateSpacesParams {
	var ()
	return &UpdateActionTemplateSpacesParams{

		timeout: timeout,
	}
}

// NewUpdateActionTemplateSpacesParamsWithContext creates a new UpdateActionTemplateSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateActionTemplateSpacesParamsWithContext(ctx context.Context) *UpdateActionTemplateSpacesParams {
	var ()
	return &UpdateActionTemplateSpacesParams{

		Context: ctx,
	}
}

// NewUpdateActionTemplateSpacesParamsWithHTTPClient creates a new UpdateActionTemplateSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateActionTemplateSpacesParamsWithHTTPClient(client *http.Client) *UpdateActionTemplateSpacesParams {
	var ()
	return &UpdateActionTemplateSpacesParams{
		HTTPClient: client,
	}
}

/*UpdateActionTemplateSpacesParams contains all the parameters to send to the API endpoint
for the update action template spaces operation typically these are written to a http.Request
*/
type UpdateActionTemplateSpacesParams struct {

	/*ActionTemplateResource
	  The ActionTemplateResource resource to create

	*/
	ActionTemplateResource *models.ActionTemplateResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the ActionTemplateResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) WithTimeout(timeout time.Duration) *UpdateActionTemplateSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) WithContext(ctx context.Context) *UpdateActionTemplateSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) WithHTTPClient(client *http.Client) *UpdateActionTemplateSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionTemplateResource adds the actionTemplateResource to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) WithActionTemplateResource(actionTemplateResource *models.ActionTemplateResource) *UpdateActionTemplateSpacesParams {
	o.SetActionTemplateResource(actionTemplateResource)
	return o
}

// SetActionTemplateResource adds the actionTemplateResource to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) SetActionTemplateResource(actionTemplateResource *models.ActionTemplateResource) {
	o.ActionTemplateResource = actionTemplateResource
}

// WithBaseSpaceID adds the baseSpaceID to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) WithBaseSpaceID(baseSpaceID string) *UpdateActionTemplateSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) WithID(id string) *UpdateActionTemplateSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update action template spaces params
func (o *UpdateActionTemplateSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateActionTemplateSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActionTemplateResource != nil {
		if err := r.SetBodyParam(o.ActionTemplateResource); err != nil {
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
