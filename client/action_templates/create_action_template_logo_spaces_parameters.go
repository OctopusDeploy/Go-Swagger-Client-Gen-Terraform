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
)

// NewCreateActionTemplateLogoSpacesParams creates a new CreateActionTemplateLogoSpacesParams object
// with the default values initialized.
func NewCreateActionTemplateLogoSpacesParams() *CreateActionTemplateLogoSpacesParams {
	var ()
	return &CreateActionTemplateLogoSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateActionTemplateLogoSpacesParamsWithTimeout creates a new CreateActionTemplateLogoSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateActionTemplateLogoSpacesParamsWithTimeout(timeout time.Duration) *CreateActionTemplateLogoSpacesParams {
	var ()
	return &CreateActionTemplateLogoSpacesParams{

		timeout: timeout,
	}
}

// NewCreateActionTemplateLogoSpacesParamsWithContext creates a new CreateActionTemplateLogoSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateActionTemplateLogoSpacesParamsWithContext(ctx context.Context) *CreateActionTemplateLogoSpacesParams {
	var ()
	return &CreateActionTemplateLogoSpacesParams{

		Context: ctx,
	}
}

// NewCreateActionTemplateLogoSpacesParamsWithHTTPClient creates a new CreateActionTemplateLogoSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateActionTemplateLogoSpacesParamsWithHTTPClient(client *http.Client) *CreateActionTemplateLogoSpacesParams {
	var ()
	return &CreateActionTemplateLogoSpacesParams{
		HTTPClient: client,
	}
}

/*CreateActionTemplateLogoSpacesParams contains all the parameters to send to the API endpoint
for the create action template logo spaces operation typically these are written to a http.Request
*/
type CreateActionTemplateLogoSpacesParams struct {

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

// WithTimeout adds the timeout to the create action template logo spaces params
func (o *CreateActionTemplateLogoSpacesParams) WithTimeout(timeout time.Duration) *CreateActionTemplateLogoSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create action template logo spaces params
func (o *CreateActionTemplateLogoSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create action template logo spaces params
func (o *CreateActionTemplateLogoSpacesParams) WithContext(ctx context.Context) *CreateActionTemplateLogoSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create action template logo spaces params
func (o *CreateActionTemplateLogoSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create action template logo spaces params
func (o *CreateActionTemplateLogoSpacesParams) WithHTTPClient(client *http.Client) *CreateActionTemplateLogoSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create action template logo spaces params
func (o *CreateActionTemplateLogoSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the create action template logo spaces params
func (o *CreateActionTemplateLogoSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateActionTemplateLogoSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create action template logo spaces params
func (o *CreateActionTemplateLogoSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the create action template logo spaces params
func (o *CreateActionTemplateLogoSpacesParams) WithID(id string) *CreateActionTemplateLogoSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create action template logo spaces params
func (o *CreateActionTemplateLogoSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateActionTemplateLogoSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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