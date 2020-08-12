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

// NewDeleteLifecycleSpacesParams creates a new DeleteLifecycleSpacesParams object
// with the default values initialized.
func NewDeleteLifecycleSpacesParams() *DeleteLifecycleSpacesParams {
	var ()
	return &DeleteLifecycleSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLifecycleSpacesParamsWithTimeout creates a new DeleteLifecycleSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLifecycleSpacesParamsWithTimeout(timeout time.Duration) *DeleteLifecycleSpacesParams {
	var ()
	return &DeleteLifecycleSpacesParams{

		timeout: timeout,
	}
}

// NewDeleteLifecycleSpacesParamsWithContext creates a new DeleteLifecycleSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLifecycleSpacesParamsWithContext(ctx context.Context) *DeleteLifecycleSpacesParams {
	var ()
	return &DeleteLifecycleSpacesParams{

		Context: ctx,
	}
}

// NewDeleteLifecycleSpacesParamsWithHTTPClient creates a new DeleteLifecycleSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLifecycleSpacesParamsWithHTTPClient(client *http.Client) *DeleteLifecycleSpacesParams {
	var ()
	return &DeleteLifecycleSpacesParams{
		HTTPClient: client,
	}
}

/*DeleteLifecycleSpacesParams contains all the parameters to send to the API endpoint
for the delete lifecycle spaces operation typically these are written to a http.Request
*/
type DeleteLifecycleSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the LifecycleResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lifecycle spaces params
func (o *DeleteLifecycleSpacesParams) WithTimeout(timeout time.Duration) *DeleteLifecycleSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lifecycle spaces params
func (o *DeleteLifecycleSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lifecycle spaces params
func (o *DeleteLifecycleSpacesParams) WithContext(ctx context.Context) *DeleteLifecycleSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lifecycle spaces params
func (o *DeleteLifecycleSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lifecycle spaces params
func (o *DeleteLifecycleSpacesParams) WithHTTPClient(client *http.Client) *DeleteLifecycleSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lifecycle spaces params
func (o *DeleteLifecycleSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the delete lifecycle spaces params
func (o *DeleteLifecycleSpacesParams) WithBaseSpaceID(baseSpaceID string) *DeleteLifecycleSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the delete lifecycle spaces params
func (o *DeleteLifecycleSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the delete lifecycle spaces params
func (o *DeleteLifecycleSpacesParams) WithID(id string) *DeleteLifecycleSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete lifecycle spaces params
func (o *DeleteLifecycleSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLifecycleSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
