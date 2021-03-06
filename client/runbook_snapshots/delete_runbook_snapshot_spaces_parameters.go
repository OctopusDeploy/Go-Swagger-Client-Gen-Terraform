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

// NewDeleteRunbookSnapshotSpacesParams creates a new DeleteRunbookSnapshotSpacesParams object
// with the default values initialized.
func NewDeleteRunbookSnapshotSpacesParams() *DeleteRunbookSnapshotSpacesParams {
	var ()
	return &DeleteRunbookSnapshotSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRunbookSnapshotSpacesParamsWithTimeout creates a new DeleteRunbookSnapshotSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRunbookSnapshotSpacesParamsWithTimeout(timeout time.Duration) *DeleteRunbookSnapshotSpacesParams {
	var ()
	return &DeleteRunbookSnapshotSpacesParams{

		timeout: timeout,
	}
}

// NewDeleteRunbookSnapshotSpacesParamsWithContext creates a new DeleteRunbookSnapshotSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRunbookSnapshotSpacesParamsWithContext(ctx context.Context) *DeleteRunbookSnapshotSpacesParams {
	var ()
	return &DeleteRunbookSnapshotSpacesParams{

		Context: ctx,
	}
}

// NewDeleteRunbookSnapshotSpacesParamsWithHTTPClient creates a new DeleteRunbookSnapshotSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRunbookSnapshotSpacesParamsWithHTTPClient(client *http.Client) *DeleteRunbookSnapshotSpacesParams {
	var ()
	return &DeleteRunbookSnapshotSpacesParams{
		HTTPClient: client,
	}
}

/*DeleteRunbookSnapshotSpacesParams contains all the parameters to send to the API endpoint
for the delete runbook snapshot spaces operation typically these are written to a http.Request
*/
type DeleteRunbookSnapshotSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the RunbookSnapshotResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete runbook snapshot spaces params
func (o *DeleteRunbookSnapshotSpacesParams) WithTimeout(timeout time.Duration) *DeleteRunbookSnapshotSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete runbook snapshot spaces params
func (o *DeleteRunbookSnapshotSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete runbook snapshot spaces params
func (o *DeleteRunbookSnapshotSpacesParams) WithContext(ctx context.Context) *DeleteRunbookSnapshotSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete runbook snapshot spaces params
func (o *DeleteRunbookSnapshotSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete runbook snapshot spaces params
func (o *DeleteRunbookSnapshotSpacesParams) WithHTTPClient(client *http.Client) *DeleteRunbookSnapshotSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete runbook snapshot spaces params
func (o *DeleteRunbookSnapshotSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the delete runbook snapshot spaces params
func (o *DeleteRunbookSnapshotSpacesParams) WithBaseSpaceID(baseSpaceID string) *DeleteRunbookSnapshotSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the delete runbook snapshot spaces params
func (o *DeleteRunbookSnapshotSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the delete runbook snapshot spaces params
func (o *DeleteRunbookSnapshotSpacesParams) WithID(id string) *DeleteRunbookSnapshotSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete runbook snapshot spaces params
func (o *DeleteRunbookSnapshotSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRunbookSnapshotSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
