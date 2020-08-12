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
)

// NewGetProjectTriggerByIDSpacesParams creates a new GetProjectTriggerByIDSpacesParams object
// with the default values initialized.
func NewGetProjectTriggerByIDSpacesParams() *GetProjectTriggerByIDSpacesParams {
	var ()
	return &GetProjectTriggerByIDSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectTriggerByIDSpacesParamsWithTimeout creates a new GetProjectTriggerByIDSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectTriggerByIDSpacesParamsWithTimeout(timeout time.Duration) *GetProjectTriggerByIDSpacesParams {
	var ()
	return &GetProjectTriggerByIDSpacesParams{

		timeout: timeout,
	}
}

// NewGetProjectTriggerByIDSpacesParamsWithContext creates a new GetProjectTriggerByIDSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectTriggerByIDSpacesParamsWithContext(ctx context.Context) *GetProjectTriggerByIDSpacesParams {
	var ()
	return &GetProjectTriggerByIDSpacesParams{

		Context: ctx,
	}
}

// NewGetProjectTriggerByIDSpacesParamsWithHTTPClient creates a new GetProjectTriggerByIDSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProjectTriggerByIDSpacesParamsWithHTTPClient(client *http.Client) *GetProjectTriggerByIDSpacesParams {
	var ()
	return &GetProjectTriggerByIDSpacesParams{
		HTTPClient: client,
	}
}

/*GetProjectTriggerByIDSpacesParams contains all the parameters to send to the API endpoint
for the get project trigger by Id spaces operation typically these are written to a http.Request
*/
type GetProjectTriggerByIDSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the ProjectTriggerResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get project trigger by Id spaces params
func (o *GetProjectTriggerByIDSpacesParams) WithTimeout(timeout time.Duration) *GetProjectTriggerByIDSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project trigger by Id spaces params
func (o *GetProjectTriggerByIDSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project trigger by Id spaces params
func (o *GetProjectTriggerByIDSpacesParams) WithContext(ctx context.Context) *GetProjectTriggerByIDSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project trigger by Id spaces params
func (o *GetProjectTriggerByIDSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project trigger by Id spaces params
func (o *GetProjectTriggerByIDSpacesParams) WithHTTPClient(client *http.Client) *GetProjectTriggerByIDSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project trigger by Id spaces params
func (o *GetProjectTriggerByIDSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get project trigger by Id spaces params
func (o *GetProjectTriggerByIDSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetProjectTriggerByIDSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get project trigger by Id spaces params
func (o *GetProjectTriggerByIDSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get project trigger by Id spaces params
func (o *GetProjectTriggerByIDSpacesParams) WithID(id string) *GetProjectTriggerByIDSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get project trigger by Id spaces params
func (o *GetProjectTriggerByIDSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectTriggerByIDSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
