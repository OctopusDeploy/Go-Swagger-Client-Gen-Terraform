// Code generated by go-swagger; DO NOT EDIT.

package runbook_runs

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

// NewDeleteRunbookRunSpacesParams creates a new DeleteRunbookRunSpacesParams object
// with the default values initialized.
func NewDeleteRunbookRunSpacesParams() *DeleteRunbookRunSpacesParams {
	var ()
	return &DeleteRunbookRunSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRunbookRunSpacesParamsWithTimeout creates a new DeleteRunbookRunSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRunbookRunSpacesParamsWithTimeout(timeout time.Duration) *DeleteRunbookRunSpacesParams {
	var ()
	return &DeleteRunbookRunSpacesParams{

		timeout: timeout,
	}
}

// NewDeleteRunbookRunSpacesParamsWithContext creates a new DeleteRunbookRunSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRunbookRunSpacesParamsWithContext(ctx context.Context) *DeleteRunbookRunSpacesParams {
	var ()
	return &DeleteRunbookRunSpacesParams{

		Context: ctx,
	}
}

// NewDeleteRunbookRunSpacesParamsWithHTTPClient creates a new DeleteRunbookRunSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRunbookRunSpacesParamsWithHTTPClient(client *http.Client) *DeleteRunbookRunSpacesParams {
	var ()
	return &DeleteRunbookRunSpacesParams{
		HTTPClient: client,
	}
}

/*DeleteRunbookRunSpacesParams contains all the parameters to send to the API endpoint
for the delete runbook run spaces operation typically these are written to a http.Request
*/
type DeleteRunbookRunSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the RunbookRunResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete runbook run spaces params
func (o *DeleteRunbookRunSpacesParams) WithTimeout(timeout time.Duration) *DeleteRunbookRunSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete runbook run spaces params
func (o *DeleteRunbookRunSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete runbook run spaces params
func (o *DeleteRunbookRunSpacesParams) WithContext(ctx context.Context) *DeleteRunbookRunSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete runbook run spaces params
func (o *DeleteRunbookRunSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete runbook run spaces params
func (o *DeleteRunbookRunSpacesParams) WithHTTPClient(client *http.Client) *DeleteRunbookRunSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete runbook run spaces params
func (o *DeleteRunbookRunSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the delete runbook run spaces params
func (o *DeleteRunbookRunSpacesParams) WithBaseSpaceID(baseSpaceID string) *DeleteRunbookRunSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the delete runbook run spaces params
func (o *DeleteRunbookRunSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the delete runbook run spaces params
func (o *DeleteRunbookRunSpacesParams) WithID(id string) *DeleteRunbookRunSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete runbook run spaces params
func (o *DeleteRunbookRunSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRunbookRunSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
