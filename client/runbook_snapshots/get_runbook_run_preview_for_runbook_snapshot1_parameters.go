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

// NewGetRunbookRunPreviewForRunbookSnapshot1Params creates a new GetRunbookRunPreviewForRunbookSnapshot1Params object
// with the default values initialized.
func NewGetRunbookRunPreviewForRunbookSnapshot1Params() *GetRunbookRunPreviewForRunbookSnapshot1Params {
	var ()
	return &GetRunbookRunPreviewForRunbookSnapshot1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunbookRunPreviewForRunbookSnapshot1ParamsWithTimeout creates a new GetRunbookRunPreviewForRunbookSnapshot1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunbookRunPreviewForRunbookSnapshot1ParamsWithTimeout(timeout time.Duration) *GetRunbookRunPreviewForRunbookSnapshot1Params {
	var ()
	return &GetRunbookRunPreviewForRunbookSnapshot1Params{

		timeout: timeout,
	}
}

// NewGetRunbookRunPreviewForRunbookSnapshot1ParamsWithContext creates a new GetRunbookRunPreviewForRunbookSnapshot1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunbookRunPreviewForRunbookSnapshot1ParamsWithContext(ctx context.Context) *GetRunbookRunPreviewForRunbookSnapshot1Params {
	var ()
	return &GetRunbookRunPreviewForRunbookSnapshot1Params{

		Context: ctx,
	}
}

// NewGetRunbookRunPreviewForRunbookSnapshot1ParamsWithHTTPClient creates a new GetRunbookRunPreviewForRunbookSnapshot1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunbookRunPreviewForRunbookSnapshot1ParamsWithHTTPClient(client *http.Client) *GetRunbookRunPreviewForRunbookSnapshot1Params {
	var ()
	return &GetRunbookRunPreviewForRunbookSnapshot1Params{
		HTTPClient: client,
	}
}

/*GetRunbookRunPreviewForRunbookSnapshot1Params contains all the parameters to send to the API endpoint
for the get runbook run preview for runbook snapshot 1 operation typically these are written to a http.Request
*/
type GetRunbookRunPreviewForRunbookSnapshot1Params struct {

	/*Environment
	  ID of the environment

	*/
	Environment string
	/*ID
	  ID of the resource

	*/
	ID string
	/*Tenant
	  ID of the tenant

	*/
	Tenant string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) WithTimeout(timeout time.Duration) *GetRunbookRunPreviewForRunbookSnapshot1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) WithContext(ctx context.Context) *GetRunbookRunPreviewForRunbookSnapshot1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) WithHTTPClient(client *http.Client) *GetRunbookRunPreviewForRunbookSnapshot1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) WithEnvironment(environment string) *GetRunbookRunPreviewForRunbookSnapshot1Params {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) SetEnvironment(environment string) {
	o.Environment = environment
}

// WithID adds the id to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) WithID(id string) *GetRunbookRunPreviewForRunbookSnapshot1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) SetID(id string) {
	o.ID = id
}

// WithTenant adds the tenant to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) WithTenant(tenant string) *GetRunbookRunPreviewForRunbookSnapshot1Params {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the get runbook run preview for runbook snapshot 1 params
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) SetTenant(tenant string) {
	o.Tenant = tenant
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunbookRunPreviewForRunbookSnapshot1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environment
	if err := r.SetPathParam("environment", o.Environment); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param tenant
	if err := r.SetPathParam("tenant", o.Tenant); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
