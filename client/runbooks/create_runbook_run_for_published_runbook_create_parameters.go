// Code generated by go-swagger; DO NOT EDIT.

package runbooks

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

// NewCreateRunbookRunForPublishedRunbookCreateParams creates a new CreateRunbookRunForPublishedRunbookCreateParams object
// with the default values initialized.
func NewCreateRunbookRunForPublishedRunbookCreateParams() *CreateRunbookRunForPublishedRunbookCreateParams {
	var ()
	return &CreateRunbookRunForPublishedRunbookCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRunbookRunForPublishedRunbookCreateParamsWithTimeout creates a new CreateRunbookRunForPublishedRunbookCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRunbookRunForPublishedRunbookCreateParamsWithTimeout(timeout time.Duration) *CreateRunbookRunForPublishedRunbookCreateParams {
	var ()
	return &CreateRunbookRunForPublishedRunbookCreateParams{

		timeout: timeout,
	}
}

// NewCreateRunbookRunForPublishedRunbookCreateParamsWithContext creates a new CreateRunbookRunForPublishedRunbookCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRunbookRunForPublishedRunbookCreateParamsWithContext(ctx context.Context) *CreateRunbookRunForPublishedRunbookCreateParams {
	var ()
	return &CreateRunbookRunForPublishedRunbookCreateParams{

		Context: ctx,
	}
}

// NewCreateRunbookRunForPublishedRunbookCreateParamsWithHTTPClient creates a new CreateRunbookRunForPublishedRunbookCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRunbookRunForPublishedRunbookCreateParamsWithHTTPClient(client *http.Client) *CreateRunbookRunForPublishedRunbookCreateParams {
	var ()
	return &CreateRunbookRunForPublishedRunbookCreateParams{
		HTTPClient: client,
	}
}

/*CreateRunbookRunForPublishedRunbookCreateParams contains all the parameters to send to the API endpoint
for the create runbook run for published runbook create operation typically these are written to a http.Request
*/
type CreateRunbookRunForPublishedRunbookCreateParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create runbook run for published runbook create params
func (o *CreateRunbookRunForPublishedRunbookCreateParams) WithTimeout(timeout time.Duration) *CreateRunbookRunForPublishedRunbookCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create runbook run for published runbook create params
func (o *CreateRunbookRunForPublishedRunbookCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create runbook run for published runbook create params
func (o *CreateRunbookRunForPublishedRunbookCreateParams) WithContext(ctx context.Context) *CreateRunbookRunForPublishedRunbookCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create runbook run for published runbook create params
func (o *CreateRunbookRunForPublishedRunbookCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create runbook run for published runbook create params
func (o *CreateRunbookRunForPublishedRunbookCreateParams) WithHTTPClient(client *http.Client) *CreateRunbookRunForPublishedRunbookCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create runbook run for published runbook create params
func (o *CreateRunbookRunForPublishedRunbookCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create runbook run for published runbook create params
func (o *CreateRunbookRunForPublishedRunbookCreateParams) WithID(id string) *CreateRunbookRunForPublishedRunbookCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create runbook run for published runbook create params
func (o *CreateRunbookRunForPublishedRunbookCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRunbookRunForPublishedRunbookCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
