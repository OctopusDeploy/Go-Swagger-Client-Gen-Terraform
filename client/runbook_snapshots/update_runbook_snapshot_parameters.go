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

	"models"
)

// NewUpdateRunbookSnapshotParams creates a new UpdateRunbookSnapshotParams object
// with the default values initialized.
func NewUpdateRunbookSnapshotParams() *UpdateRunbookSnapshotParams {
	var ()
	return &UpdateRunbookSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRunbookSnapshotParamsWithTimeout creates a new UpdateRunbookSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRunbookSnapshotParamsWithTimeout(timeout time.Duration) *UpdateRunbookSnapshotParams {
	var ()
	return &UpdateRunbookSnapshotParams{

		timeout: timeout,
	}
}

// NewUpdateRunbookSnapshotParamsWithContext creates a new UpdateRunbookSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRunbookSnapshotParamsWithContext(ctx context.Context) *UpdateRunbookSnapshotParams {
	var ()
	return &UpdateRunbookSnapshotParams{

		Context: ctx,
	}
}

// NewUpdateRunbookSnapshotParamsWithHTTPClient creates a new UpdateRunbookSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRunbookSnapshotParamsWithHTTPClient(client *http.Client) *UpdateRunbookSnapshotParams {
	var ()
	return &UpdateRunbookSnapshotParams{
		HTTPClient: client,
	}
}

/*UpdateRunbookSnapshotParams contains all the parameters to send to the API endpoint
for the update runbook snapshot operation typically these are written to a http.Request
*/
type UpdateRunbookSnapshotParams struct {

	/*RunbookSnapshotResource
	  The RunbookSnapshotResource resource to create

	*/
	RunbookSnapshotResource *models.RunbookSnapshotResource
	/*ID
	  ID of the RunbookSnapshotResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update runbook snapshot params
func (o *UpdateRunbookSnapshotParams) WithTimeout(timeout time.Duration) *UpdateRunbookSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update runbook snapshot params
func (o *UpdateRunbookSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update runbook snapshot params
func (o *UpdateRunbookSnapshotParams) WithContext(ctx context.Context) *UpdateRunbookSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update runbook snapshot params
func (o *UpdateRunbookSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update runbook snapshot params
func (o *UpdateRunbookSnapshotParams) WithHTTPClient(client *http.Client) *UpdateRunbookSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update runbook snapshot params
func (o *UpdateRunbookSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRunbookSnapshotResource adds the runbookSnapshotResource to the update runbook snapshot params
func (o *UpdateRunbookSnapshotParams) WithRunbookSnapshotResource(runbookSnapshotResource *models.RunbookSnapshotResource) *UpdateRunbookSnapshotParams {
	o.SetRunbookSnapshotResource(runbookSnapshotResource)
	return o
}

// SetRunbookSnapshotResource adds the runbookSnapshotResource to the update runbook snapshot params
func (o *UpdateRunbookSnapshotParams) SetRunbookSnapshotResource(runbookSnapshotResource *models.RunbookSnapshotResource) {
	o.RunbookSnapshotResource = runbookSnapshotResource
}

// WithID adds the id to the update runbook snapshot params
func (o *UpdateRunbookSnapshotParams) WithID(id string) *UpdateRunbookSnapshotParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update runbook snapshot params
func (o *UpdateRunbookSnapshotParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRunbookSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RunbookSnapshotResource != nil {
		if err := r.SetBodyParam(o.RunbookSnapshotResource); err != nil {
			return err
		}
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
