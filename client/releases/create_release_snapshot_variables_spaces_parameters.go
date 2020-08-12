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

// NewCreateReleaseSnapshotVariablesSpacesParams creates a new CreateReleaseSnapshotVariablesSpacesParams object
// with the default values initialized.
func NewCreateReleaseSnapshotVariablesSpacesParams() *CreateReleaseSnapshotVariablesSpacesParams {
	var ()
	return &CreateReleaseSnapshotVariablesSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateReleaseSnapshotVariablesSpacesParamsWithTimeout creates a new CreateReleaseSnapshotVariablesSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateReleaseSnapshotVariablesSpacesParamsWithTimeout(timeout time.Duration) *CreateReleaseSnapshotVariablesSpacesParams {
	var ()
	return &CreateReleaseSnapshotVariablesSpacesParams{

		timeout: timeout,
	}
}

// NewCreateReleaseSnapshotVariablesSpacesParamsWithContext creates a new CreateReleaseSnapshotVariablesSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateReleaseSnapshotVariablesSpacesParamsWithContext(ctx context.Context) *CreateReleaseSnapshotVariablesSpacesParams {
	var ()
	return &CreateReleaseSnapshotVariablesSpacesParams{

		Context: ctx,
	}
}

// NewCreateReleaseSnapshotVariablesSpacesParamsWithHTTPClient creates a new CreateReleaseSnapshotVariablesSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateReleaseSnapshotVariablesSpacesParamsWithHTTPClient(client *http.Client) *CreateReleaseSnapshotVariablesSpacesParams {
	var ()
	return &CreateReleaseSnapshotVariablesSpacesParams{
		HTTPClient: client,
	}
}

/*CreateReleaseSnapshotVariablesSpacesParams contains all the parameters to send to the API endpoint
for the create release snapshot variables spaces operation typically these are written to a http.Request
*/
type CreateReleaseSnapshotVariablesSpacesParams struct {

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

// WithTimeout adds the timeout to the create release snapshot variables spaces params
func (o *CreateReleaseSnapshotVariablesSpacesParams) WithTimeout(timeout time.Duration) *CreateReleaseSnapshotVariablesSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create release snapshot variables spaces params
func (o *CreateReleaseSnapshotVariablesSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create release snapshot variables spaces params
func (o *CreateReleaseSnapshotVariablesSpacesParams) WithContext(ctx context.Context) *CreateReleaseSnapshotVariablesSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create release snapshot variables spaces params
func (o *CreateReleaseSnapshotVariablesSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create release snapshot variables spaces params
func (o *CreateReleaseSnapshotVariablesSpacesParams) WithHTTPClient(client *http.Client) *CreateReleaseSnapshotVariablesSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create release snapshot variables spaces params
func (o *CreateReleaseSnapshotVariablesSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the create release snapshot variables spaces params
func (o *CreateReleaseSnapshotVariablesSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateReleaseSnapshotVariablesSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create release snapshot variables spaces params
func (o *CreateReleaseSnapshotVariablesSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the create release snapshot variables spaces params
func (o *CreateReleaseSnapshotVariablesSpacesParams) WithID(id string) *CreateReleaseSnapshotVariablesSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create release snapshot variables spaces params
func (o *CreateReleaseSnapshotVariablesSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateReleaseSnapshotVariablesSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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