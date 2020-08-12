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

// NewGetRunbookRunPreviewForRunbookSnapshotSpacesParams creates a new GetRunbookRunPreviewForRunbookSnapshotSpacesParams object
// with the default values initialized.
func NewGetRunbookRunPreviewForRunbookSnapshotSpacesParams() *GetRunbookRunPreviewForRunbookSnapshotSpacesParams {
	var ()
	return &GetRunbookRunPreviewForRunbookSnapshotSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunbookRunPreviewForRunbookSnapshotSpacesParamsWithTimeout creates a new GetRunbookRunPreviewForRunbookSnapshotSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunbookRunPreviewForRunbookSnapshotSpacesParamsWithTimeout(timeout time.Duration) *GetRunbookRunPreviewForRunbookSnapshotSpacesParams {
	var ()
	return &GetRunbookRunPreviewForRunbookSnapshotSpacesParams{

		timeout: timeout,
	}
}

// NewGetRunbookRunPreviewForRunbookSnapshotSpacesParamsWithContext creates a new GetRunbookRunPreviewForRunbookSnapshotSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunbookRunPreviewForRunbookSnapshotSpacesParamsWithContext(ctx context.Context) *GetRunbookRunPreviewForRunbookSnapshotSpacesParams {
	var ()
	return &GetRunbookRunPreviewForRunbookSnapshotSpacesParams{

		Context: ctx,
	}
}

// NewGetRunbookRunPreviewForRunbookSnapshotSpacesParamsWithHTTPClient creates a new GetRunbookRunPreviewForRunbookSnapshotSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunbookRunPreviewForRunbookSnapshotSpacesParamsWithHTTPClient(client *http.Client) *GetRunbookRunPreviewForRunbookSnapshotSpacesParams {
	var ()
	return &GetRunbookRunPreviewForRunbookSnapshotSpacesParams{
		HTTPClient: client,
	}
}

/*GetRunbookRunPreviewForRunbookSnapshotSpacesParams contains all the parameters to send to the API endpoint
for the get runbook run preview for runbook snapshot spaces operation typically these are written to a http.Request
*/
type GetRunbookRunPreviewForRunbookSnapshotSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*Environment
	  ID of the environment

	*/
	Environment string
	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) WithTimeout(timeout time.Duration) *GetRunbookRunPreviewForRunbookSnapshotSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) WithContext(ctx context.Context) *GetRunbookRunPreviewForRunbookSnapshotSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) WithHTTPClient(client *http.Client) *GetRunbookRunPreviewForRunbookSnapshotSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetRunbookRunPreviewForRunbookSnapshotSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithEnvironment adds the environment to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) WithEnvironment(environment string) *GetRunbookRunPreviewForRunbookSnapshotSpacesParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) SetEnvironment(environment string) {
	o.Environment = environment
}

// WithID adds the id to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) WithID(id string) *GetRunbookRunPreviewForRunbookSnapshotSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get runbook run preview for runbook snapshot spaces params
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	// path param environment
	if err := r.SetPathParam("environment", o.Environment); err != nil {
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
