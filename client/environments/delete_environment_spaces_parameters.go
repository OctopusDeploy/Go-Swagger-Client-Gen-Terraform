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
)

// NewDeleteEnvironmentSpacesParams creates a new DeleteEnvironmentSpacesParams object
// with the default values initialized.
func NewDeleteEnvironmentSpacesParams() *DeleteEnvironmentSpacesParams {
	var ()
	return &DeleteEnvironmentSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEnvironmentSpacesParamsWithTimeout creates a new DeleteEnvironmentSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEnvironmentSpacesParamsWithTimeout(timeout time.Duration) *DeleteEnvironmentSpacesParams {
	var ()
	return &DeleteEnvironmentSpacesParams{

		timeout: timeout,
	}
}

// NewDeleteEnvironmentSpacesParamsWithContext creates a new DeleteEnvironmentSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEnvironmentSpacesParamsWithContext(ctx context.Context) *DeleteEnvironmentSpacesParams {
	var ()
	return &DeleteEnvironmentSpacesParams{

		Context: ctx,
	}
}

// NewDeleteEnvironmentSpacesParamsWithHTTPClient creates a new DeleteEnvironmentSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEnvironmentSpacesParamsWithHTTPClient(client *http.Client) *DeleteEnvironmentSpacesParams {
	var ()
	return &DeleteEnvironmentSpacesParams{
		HTTPClient: client,
	}
}

/*DeleteEnvironmentSpacesParams contains all the parameters to send to the API endpoint
for the delete environment spaces operation typically these are written to a http.Request
*/
type DeleteEnvironmentSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the EnvironmentResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete environment spaces params
func (o *DeleteEnvironmentSpacesParams) WithTimeout(timeout time.Duration) *DeleteEnvironmentSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete environment spaces params
func (o *DeleteEnvironmentSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete environment spaces params
func (o *DeleteEnvironmentSpacesParams) WithContext(ctx context.Context) *DeleteEnvironmentSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete environment spaces params
func (o *DeleteEnvironmentSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete environment spaces params
func (o *DeleteEnvironmentSpacesParams) WithHTTPClient(client *http.Client) *DeleteEnvironmentSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete environment spaces params
func (o *DeleteEnvironmentSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the delete environment spaces params
func (o *DeleteEnvironmentSpacesParams) WithBaseSpaceID(baseSpaceID string) *DeleteEnvironmentSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the delete environment spaces params
func (o *DeleteEnvironmentSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the delete environment spaces params
func (o *DeleteEnvironmentSpacesParams) WithID(id string) *DeleteEnvironmentSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete environment spaces params
func (o *DeleteEnvironmentSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEnvironmentSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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