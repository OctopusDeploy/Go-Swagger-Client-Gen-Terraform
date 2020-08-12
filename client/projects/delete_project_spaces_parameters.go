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

// NewDeleteProjectSpacesParams creates a new DeleteProjectSpacesParams object
// with the default values initialized.
func NewDeleteProjectSpacesParams() *DeleteProjectSpacesParams {
	var ()
	return &DeleteProjectSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectSpacesParamsWithTimeout creates a new DeleteProjectSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProjectSpacesParamsWithTimeout(timeout time.Duration) *DeleteProjectSpacesParams {
	var ()
	return &DeleteProjectSpacesParams{

		timeout: timeout,
	}
}

// NewDeleteProjectSpacesParamsWithContext creates a new DeleteProjectSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProjectSpacesParamsWithContext(ctx context.Context) *DeleteProjectSpacesParams {
	var ()
	return &DeleteProjectSpacesParams{

		Context: ctx,
	}
}

// NewDeleteProjectSpacesParamsWithHTTPClient creates a new DeleteProjectSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProjectSpacesParamsWithHTTPClient(client *http.Client) *DeleteProjectSpacesParams {
	var ()
	return &DeleteProjectSpacesParams{
		HTTPClient: client,
	}
}

/*DeleteProjectSpacesParams contains all the parameters to send to the API endpoint
for the delete project spaces operation typically these are written to a http.Request
*/
type DeleteProjectSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the ProjectResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete project spaces params
func (o *DeleteProjectSpacesParams) WithTimeout(timeout time.Duration) *DeleteProjectSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project spaces params
func (o *DeleteProjectSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project spaces params
func (o *DeleteProjectSpacesParams) WithContext(ctx context.Context) *DeleteProjectSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project spaces params
func (o *DeleteProjectSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project spaces params
func (o *DeleteProjectSpacesParams) WithHTTPClient(client *http.Client) *DeleteProjectSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project spaces params
func (o *DeleteProjectSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the delete project spaces params
func (o *DeleteProjectSpacesParams) WithBaseSpaceID(baseSpaceID string) *DeleteProjectSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the delete project spaces params
func (o *DeleteProjectSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the delete project spaces params
func (o *DeleteProjectSpacesParams) WithID(id string) *DeleteProjectSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete project spaces params
func (o *DeleteProjectSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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