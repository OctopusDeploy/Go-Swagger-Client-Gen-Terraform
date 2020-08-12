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

// NewGetRunbookRunPreviewForRunbookSpaces1Params creates a new GetRunbookRunPreviewForRunbookSpaces1Params object
// with the default values initialized.
func NewGetRunbookRunPreviewForRunbookSpaces1Params() *GetRunbookRunPreviewForRunbookSpaces1Params {
	var ()
	return &GetRunbookRunPreviewForRunbookSpaces1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunbookRunPreviewForRunbookSpaces1ParamsWithTimeout creates a new GetRunbookRunPreviewForRunbookSpaces1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunbookRunPreviewForRunbookSpaces1ParamsWithTimeout(timeout time.Duration) *GetRunbookRunPreviewForRunbookSpaces1Params {
	var ()
	return &GetRunbookRunPreviewForRunbookSpaces1Params{

		timeout: timeout,
	}
}

// NewGetRunbookRunPreviewForRunbookSpaces1ParamsWithContext creates a new GetRunbookRunPreviewForRunbookSpaces1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunbookRunPreviewForRunbookSpaces1ParamsWithContext(ctx context.Context) *GetRunbookRunPreviewForRunbookSpaces1Params {
	var ()
	return &GetRunbookRunPreviewForRunbookSpaces1Params{

		Context: ctx,
	}
}

// NewGetRunbookRunPreviewForRunbookSpaces1ParamsWithHTTPClient creates a new GetRunbookRunPreviewForRunbookSpaces1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunbookRunPreviewForRunbookSpaces1ParamsWithHTTPClient(client *http.Client) *GetRunbookRunPreviewForRunbookSpaces1Params {
	var ()
	return &GetRunbookRunPreviewForRunbookSpaces1Params{
		HTTPClient: client,
	}
}

/*GetRunbookRunPreviewForRunbookSpaces1Params contains all the parameters to send to the API endpoint
for the get runbook run preview for runbook spaces 1 operation typically these are written to a http.Request
*/
type GetRunbookRunPreviewForRunbookSpaces1Params struct {

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
	/*Tenant
	  ID of the tenant

	*/
	Tenant string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) WithTimeout(timeout time.Duration) *GetRunbookRunPreviewForRunbookSpaces1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) WithContext(ctx context.Context) *GetRunbookRunPreviewForRunbookSpaces1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) WithHTTPClient(client *http.Client) *GetRunbookRunPreviewForRunbookSpaces1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) WithBaseSpaceID(baseSpaceID string) *GetRunbookRunPreviewForRunbookSpaces1Params {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithEnvironment adds the environment to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) WithEnvironment(environment string) *GetRunbookRunPreviewForRunbookSpaces1Params {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) SetEnvironment(environment string) {
	o.Environment = environment
}

// WithID adds the id to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) WithID(id string) *GetRunbookRunPreviewForRunbookSpaces1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) SetID(id string) {
	o.ID = id
}

// WithTenant adds the tenant to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) WithTenant(tenant string) *GetRunbookRunPreviewForRunbookSpaces1Params {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the get runbook run preview for runbook spaces 1 params
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) SetTenant(tenant string) {
	o.Tenant = tenant
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunbookRunPreviewForRunbookSpaces1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tenant
	if err := r.SetPathParam("tenant", o.Tenant); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
