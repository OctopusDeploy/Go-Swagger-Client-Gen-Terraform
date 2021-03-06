// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewCreateProjectLogoSpacesParams creates a new CreateProjectLogoSpacesParams object
// with the default values initialized.
func NewCreateProjectLogoSpacesParams() *CreateProjectLogoSpacesParams {
	var ()
	return &CreateProjectLogoSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectLogoSpacesParamsWithTimeout creates a new CreateProjectLogoSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateProjectLogoSpacesParamsWithTimeout(timeout time.Duration) *CreateProjectLogoSpacesParams {
	var ()
	return &CreateProjectLogoSpacesParams{

		timeout: timeout,
	}
}

// NewCreateProjectLogoSpacesParamsWithContext creates a new CreateProjectLogoSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateProjectLogoSpacesParamsWithContext(ctx context.Context) *CreateProjectLogoSpacesParams {
	var ()
	return &CreateProjectLogoSpacesParams{

		Context: ctx,
	}
}

// NewCreateProjectLogoSpacesParamsWithHTTPClient creates a new CreateProjectLogoSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateProjectLogoSpacesParamsWithHTTPClient(client *http.Client) *CreateProjectLogoSpacesParams {
	var ()
	return &CreateProjectLogoSpacesParams{
		HTTPClient: client,
	}
}

/*CreateProjectLogoSpacesParams contains all the parameters to send to the API endpoint
for the create project logo spaces operation typically these are written to a http.Request
*/
type CreateProjectLogoSpacesParams struct {

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

// WithTimeout adds the timeout to the create project logo spaces params
func (o *CreateProjectLogoSpacesParams) WithTimeout(timeout time.Duration) *CreateProjectLogoSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project logo spaces params
func (o *CreateProjectLogoSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project logo spaces params
func (o *CreateProjectLogoSpacesParams) WithContext(ctx context.Context) *CreateProjectLogoSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project logo spaces params
func (o *CreateProjectLogoSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project logo spaces params
func (o *CreateProjectLogoSpacesParams) WithHTTPClient(client *http.Client) *CreateProjectLogoSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project logo spaces params
func (o *CreateProjectLogoSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the create project logo spaces params
func (o *CreateProjectLogoSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateProjectLogoSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create project logo spaces params
func (o *CreateProjectLogoSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the create project logo spaces params
func (o *CreateProjectLogoSpacesParams) WithID(id string) *CreateProjectLogoSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create project logo spaces params
func (o *CreateProjectLogoSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProjectLogoSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
