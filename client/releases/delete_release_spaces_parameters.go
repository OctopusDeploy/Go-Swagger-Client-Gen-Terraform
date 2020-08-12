// Code generated by go-swagger; DO NOT EDIT.

package releases

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

// NewDeleteReleaseSpacesParams creates a new DeleteReleaseSpacesParams object
// with the default values initialized.
func NewDeleteReleaseSpacesParams() *DeleteReleaseSpacesParams {
	var ()
	return &DeleteReleaseSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteReleaseSpacesParamsWithTimeout creates a new DeleteReleaseSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteReleaseSpacesParamsWithTimeout(timeout time.Duration) *DeleteReleaseSpacesParams {
	var ()
	return &DeleteReleaseSpacesParams{

		timeout: timeout,
	}
}

// NewDeleteReleaseSpacesParamsWithContext creates a new DeleteReleaseSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteReleaseSpacesParamsWithContext(ctx context.Context) *DeleteReleaseSpacesParams {
	var ()
	return &DeleteReleaseSpacesParams{

		Context: ctx,
	}
}

// NewDeleteReleaseSpacesParamsWithHTTPClient creates a new DeleteReleaseSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteReleaseSpacesParamsWithHTTPClient(client *http.Client) *DeleteReleaseSpacesParams {
	var ()
	return &DeleteReleaseSpacesParams{
		HTTPClient: client,
	}
}

/*DeleteReleaseSpacesParams contains all the parameters to send to the API endpoint
for the delete release spaces operation typically these are written to a http.Request
*/
type DeleteReleaseSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the ReleaseResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete release spaces params
func (o *DeleteReleaseSpacesParams) WithTimeout(timeout time.Duration) *DeleteReleaseSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete release spaces params
func (o *DeleteReleaseSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete release spaces params
func (o *DeleteReleaseSpacesParams) WithContext(ctx context.Context) *DeleteReleaseSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete release spaces params
func (o *DeleteReleaseSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete release spaces params
func (o *DeleteReleaseSpacesParams) WithHTTPClient(client *http.Client) *DeleteReleaseSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete release spaces params
func (o *DeleteReleaseSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the delete release spaces params
func (o *DeleteReleaseSpacesParams) WithBaseSpaceID(baseSpaceID string) *DeleteReleaseSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the delete release spaces params
func (o *DeleteReleaseSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the delete release spaces params
func (o *DeleteReleaseSpacesParams) WithID(id string) *DeleteReleaseSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete release spaces params
func (o *DeleteReleaseSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteReleaseSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
