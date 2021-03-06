// Code generated by go-swagger; DO NOT EDIT.

package runbook_snapshots

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

// NewCreateRunbookSnapshotSnapshotVariablesSpacesParams creates a new CreateRunbookSnapshotSnapshotVariablesSpacesParams object
// with the default values initialized.
func NewCreateRunbookSnapshotSnapshotVariablesSpacesParams() *CreateRunbookSnapshotSnapshotVariablesSpacesParams {
	var ()
	return &CreateRunbookSnapshotSnapshotVariablesSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRunbookSnapshotSnapshotVariablesSpacesParamsWithTimeout creates a new CreateRunbookSnapshotSnapshotVariablesSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRunbookSnapshotSnapshotVariablesSpacesParamsWithTimeout(timeout time.Duration) *CreateRunbookSnapshotSnapshotVariablesSpacesParams {
	var ()
	return &CreateRunbookSnapshotSnapshotVariablesSpacesParams{

		timeout: timeout,
	}
}

// NewCreateRunbookSnapshotSnapshotVariablesSpacesParamsWithContext creates a new CreateRunbookSnapshotSnapshotVariablesSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRunbookSnapshotSnapshotVariablesSpacesParamsWithContext(ctx context.Context) *CreateRunbookSnapshotSnapshotVariablesSpacesParams {
	var ()
	return &CreateRunbookSnapshotSnapshotVariablesSpacesParams{

		Context: ctx,
	}
}

// NewCreateRunbookSnapshotSnapshotVariablesSpacesParamsWithHTTPClient creates a new CreateRunbookSnapshotSnapshotVariablesSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRunbookSnapshotSnapshotVariablesSpacesParamsWithHTTPClient(client *http.Client) *CreateRunbookSnapshotSnapshotVariablesSpacesParams {
	var ()
	return &CreateRunbookSnapshotSnapshotVariablesSpacesParams{
		HTTPClient: client,
	}
}

/*CreateRunbookSnapshotSnapshotVariablesSpacesParams contains all the parameters to send to the API endpoint
for the create runbook snapshot snapshot variables spaces operation typically these are written to a http.Request
*/
type CreateRunbookSnapshotSnapshotVariablesSpacesParams struct {

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

// WithTimeout adds the timeout to the create runbook snapshot snapshot variables spaces params
func (o *CreateRunbookSnapshotSnapshotVariablesSpacesParams) WithTimeout(timeout time.Duration) *CreateRunbookSnapshotSnapshotVariablesSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create runbook snapshot snapshot variables spaces params
func (o *CreateRunbookSnapshotSnapshotVariablesSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create runbook snapshot snapshot variables spaces params
func (o *CreateRunbookSnapshotSnapshotVariablesSpacesParams) WithContext(ctx context.Context) *CreateRunbookSnapshotSnapshotVariablesSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create runbook snapshot snapshot variables spaces params
func (o *CreateRunbookSnapshotSnapshotVariablesSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create runbook snapshot snapshot variables spaces params
func (o *CreateRunbookSnapshotSnapshotVariablesSpacesParams) WithHTTPClient(client *http.Client) *CreateRunbookSnapshotSnapshotVariablesSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create runbook snapshot snapshot variables spaces params
func (o *CreateRunbookSnapshotSnapshotVariablesSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the create runbook snapshot snapshot variables spaces params
func (o *CreateRunbookSnapshotSnapshotVariablesSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateRunbookSnapshotSnapshotVariablesSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create runbook snapshot snapshot variables spaces params
func (o *CreateRunbookSnapshotSnapshotVariablesSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the create runbook snapshot snapshot variables spaces params
func (o *CreateRunbookSnapshotSnapshotVariablesSpacesParams) WithID(id string) *CreateRunbookSnapshotSnapshotVariablesSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create runbook snapshot snapshot variables spaces params
func (o *CreateRunbookSnapshotSnapshotVariablesSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRunbookSnapshotSnapshotVariablesSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
