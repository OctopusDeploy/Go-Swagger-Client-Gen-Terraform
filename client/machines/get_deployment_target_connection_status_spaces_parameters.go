// Code generated by go-swagger; DO NOT EDIT.

package machines

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

// NewGetDeploymentTargetConnectionStatusSpacesParams creates a new GetDeploymentTargetConnectionStatusSpacesParams object
// with the default values initialized.
func NewGetDeploymentTargetConnectionStatusSpacesParams() *GetDeploymentTargetConnectionStatusSpacesParams {
	var ()
	return &GetDeploymentTargetConnectionStatusSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentTargetConnectionStatusSpacesParamsWithTimeout creates a new GetDeploymentTargetConnectionStatusSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentTargetConnectionStatusSpacesParamsWithTimeout(timeout time.Duration) *GetDeploymentTargetConnectionStatusSpacesParams {
	var ()
	return &GetDeploymentTargetConnectionStatusSpacesParams{

		timeout: timeout,
	}
}

// NewGetDeploymentTargetConnectionStatusSpacesParamsWithContext creates a new GetDeploymentTargetConnectionStatusSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentTargetConnectionStatusSpacesParamsWithContext(ctx context.Context) *GetDeploymentTargetConnectionStatusSpacesParams {
	var ()
	return &GetDeploymentTargetConnectionStatusSpacesParams{

		Context: ctx,
	}
}

// NewGetDeploymentTargetConnectionStatusSpacesParamsWithHTTPClient creates a new GetDeploymentTargetConnectionStatusSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentTargetConnectionStatusSpacesParamsWithHTTPClient(client *http.Client) *GetDeploymentTargetConnectionStatusSpacesParams {
	var ()
	return &GetDeploymentTargetConnectionStatusSpacesParams{
		HTTPClient: client,
	}
}

/*GetDeploymentTargetConnectionStatusSpacesParams contains all the parameters to send to the API endpoint
for the get deployment target connection status spaces operation typically these are written to a http.Request
*/
type GetDeploymentTargetConnectionStatusSpacesParams struct {

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

// WithTimeout adds the timeout to the get deployment target connection status spaces params
func (o *GetDeploymentTargetConnectionStatusSpacesParams) WithTimeout(timeout time.Duration) *GetDeploymentTargetConnectionStatusSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment target connection status spaces params
func (o *GetDeploymentTargetConnectionStatusSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment target connection status spaces params
func (o *GetDeploymentTargetConnectionStatusSpacesParams) WithContext(ctx context.Context) *GetDeploymentTargetConnectionStatusSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment target connection status spaces params
func (o *GetDeploymentTargetConnectionStatusSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment target connection status spaces params
func (o *GetDeploymentTargetConnectionStatusSpacesParams) WithHTTPClient(client *http.Client) *GetDeploymentTargetConnectionStatusSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment target connection status spaces params
func (o *GetDeploymentTargetConnectionStatusSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get deployment target connection status spaces params
func (o *GetDeploymentTargetConnectionStatusSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetDeploymentTargetConnectionStatusSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get deployment target connection status spaces params
func (o *GetDeploymentTargetConnectionStatusSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get deployment target connection status spaces params
func (o *GetDeploymentTargetConnectionStatusSpacesParams) WithID(id string) *GetDeploymentTargetConnectionStatusSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get deployment target connection status spaces params
func (o *GetDeploymentTargetConnectionStatusSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentTargetConnectionStatusSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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