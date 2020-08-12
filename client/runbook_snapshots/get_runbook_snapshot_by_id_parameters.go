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

// NewGetRunbookSnapshotByIDParams creates a new GetRunbookSnapshotByIDParams object
// with the default values initialized.
func NewGetRunbookSnapshotByIDParams() *GetRunbookSnapshotByIDParams {
	var ()
	return &GetRunbookSnapshotByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunbookSnapshotByIDParamsWithTimeout creates a new GetRunbookSnapshotByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunbookSnapshotByIDParamsWithTimeout(timeout time.Duration) *GetRunbookSnapshotByIDParams {
	var ()
	return &GetRunbookSnapshotByIDParams{

		timeout: timeout,
	}
}

// NewGetRunbookSnapshotByIDParamsWithContext creates a new GetRunbookSnapshotByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunbookSnapshotByIDParamsWithContext(ctx context.Context) *GetRunbookSnapshotByIDParams {
	var ()
	return &GetRunbookSnapshotByIDParams{

		Context: ctx,
	}
}

// NewGetRunbookSnapshotByIDParamsWithHTTPClient creates a new GetRunbookSnapshotByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunbookSnapshotByIDParamsWithHTTPClient(client *http.Client) *GetRunbookSnapshotByIDParams {
	var ()
	return &GetRunbookSnapshotByIDParams{
		HTTPClient: client,
	}
}

/*GetRunbookSnapshotByIDParams contains all the parameters to send to the API endpoint
for the get runbook snapshot by Id operation typically these are written to a http.Request
*/
type GetRunbookSnapshotByIDParams struct {

	/*ID
	  ID of the RunbookSnapshotResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runbook snapshot by Id params
func (o *GetRunbookSnapshotByIDParams) WithTimeout(timeout time.Duration) *GetRunbookSnapshotByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runbook snapshot by Id params
func (o *GetRunbookSnapshotByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runbook snapshot by Id params
func (o *GetRunbookSnapshotByIDParams) WithContext(ctx context.Context) *GetRunbookSnapshotByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runbook snapshot by Id params
func (o *GetRunbookSnapshotByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runbook snapshot by Id params
func (o *GetRunbookSnapshotByIDParams) WithHTTPClient(client *http.Client) *GetRunbookSnapshotByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runbook snapshot by Id params
func (o *GetRunbookSnapshotByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get runbook snapshot by Id params
func (o *GetRunbookSnapshotByIDParams) WithID(id string) *GetRunbookSnapshotByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get runbook snapshot by Id params
func (o *GetRunbookSnapshotByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunbookSnapshotByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}