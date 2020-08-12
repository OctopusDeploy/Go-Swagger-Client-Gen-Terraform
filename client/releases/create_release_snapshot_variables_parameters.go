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

// NewCreateReleaseSnapshotVariablesParams creates a new CreateReleaseSnapshotVariablesParams object
// with the default values initialized.
func NewCreateReleaseSnapshotVariablesParams() *CreateReleaseSnapshotVariablesParams {
	var ()
	return &CreateReleaseSnapshotVariablesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateReleaseSnapshotVariablesParamsWithTimeout creates a new CreateReleaseSnapshotVariablesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateReleaseSnapshotVariablesParamsWithTimeout(timeout time.Duration) *CreateReleaseSnapshotVariablesParams {
	var ()
	return &CreateReleaseSnapshotVariablesParams{

		timeout: timeout,
	}
}

// NewCreateReleaseSnapshotVariablesParamsWithContext creates a new CreateReleaseSnapshotVariablesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateReleaseSnapshotVariablesParamsWithContext(ctx context.Context) *CreateReleaseSnapshotVariablesParams {
	var ()
	return &CreateReleaseSnapshotVariablesParams{

		Context: ctx,
	}
}

// NewCreateReleaseSnapshotVariablesParamsWithHTTPClient creates a new CreateReleaseSnapshotVariablesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateReleaseSnapshotVariablesParamsWithHTTPClient(client *http.Client) *CreateReleaseSnapshotVariablesParams {
	var ()
	return &CreateReleaseSnapshotVariablesParams{
		HTTPClient: client,
	}
}

/*CreateReleaseSnapshotVariablesParams contains all the parameters to send to the API endpoint
for the create release snapshot variables operation typically these are written to a http.Request
*/
type CreateReleaseSnapshotVariablesParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create release snapshot variables params
func (o *CreateReleaseSnapshotVariablesParams) WithTimeout(timeout time.Duration) *CreateReleaseSnapshotVariablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create release snapshot variables params
func (o *CreateReleaseSnapshotVariablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create release snapshot variables params
func (o *CreateReleaseSnapshotVariablesParams) WithContext(ctx context.Context) *CreateReleaseSnapshotVariablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create release snapshot variables params
func (o *CreateReleaseSnapshotVariablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create release snapshot variables params
func (o *CreateReleaseSnapshotVariablesParams) WithHTTPClient(client *http.Client) *CreateReleaseSnapshotVariablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create release snapshot variables params
func (o *CreateReleaseSnapshotVariablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create release snapshot variables params
func (o *CreateReleaseSnapshotVariablesParams) WithID(id string) *CreateReleaseSnapshotVariablesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create release snapshot variables params
func (o *CreateReleaseSnapshotVariablesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateReleaseSnapshotVariablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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