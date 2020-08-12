// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

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

// NewGetLifecycleProjectsSpacesParams creates a new GetLifecycleProjectsSpacesParams object
// with the default values initialized.
func NewGetLifecycleProjectsSpacesParams() *GetLifecycleProjectsSpacesParams {
	var ()
	return &GetLifecycleProjectsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLifecycleProjectsSpacesParamsWithTimeout creates a new GetLifecycleProjectsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLifecycleProjectsSpacesParamsWithTimeout(timeout time.Duration) *GetLifecycleProjectsSpacesParams {
	var ()
	return &GetLifecycleProjectsSpacesParams{

		timeout: timeout,
	}
}

// NewGetLifecycleProjectsSpacesParamsWithContext creates a new GetLifecycleProjectsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLifecycleProjectsSpacesParamsWithContext(ctx context.Context) *GetLifecycleProjectsSpacesParams {
	var ()
	return &GetLifecycleProjectsSpacesParams{

		Context: ctx,
	}
}

// NewGetLifecycleProjectsSpacesParamsWithHTTPClient creates a new GetLifecycleProjectsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLifecycleProjectsSpacesParamsWithHTTPClient(client *http.Client) *GetLifecycleProjectsSpacesParams {
	var ()
	return &GetLifecycleProjectsSpacesParams{
		HTTPClient: client,
	}
}

/*GetLifecycleProjectsSpacesParams contains all the parameters to send to the API endpoint
for the get lifecycle projects spaces operation typically these are written to a http.Request
*/
type GetLifecycleProjectsSpacesParams struct {

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

// WithTimeout adds the timeout to the get lifecycle projects spaces params
func (o *GetLifecycleProjectsSpacesParams) WithTimeout(timeout time.Duration) *GetLifecycleProjectsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lifecycle projects spaces params
func (o *GetLifecycleProjectsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lifecycle projects spaces params
func (o *GetLifecycleProjectsSpacesParams) WithContext(ctx context.Context) *GetLifecycleProjectsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lifecycle projects spaces params
func (o *GetLifecycleProjectsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lifecycle projects spaces params
func (o *GetLifecycleProjectsSpacesParams) WithHTTPClient(client *http.Client) *GetLifecycleProjectsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lifecycle projects spaces params
func (o *GetLifecycleProjectsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get lifecycle projects spaces params
func (o *GetLifecycleProjectsSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetLifecycleProjectsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get lifecycle projects spaces params
func (o *GetLifecycleProjectsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get lifecycle projects spaces params
func (o *GetLifecycleProjectsSpacesParams) WithID(id string) *GetLifecycleProjectsSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get lifecycle projects spaces params
func (o *GetLifecycleProjectsSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLifecycleProjectsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
