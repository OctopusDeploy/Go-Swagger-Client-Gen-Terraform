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

// NewGetDeploymentPreviewParams creates a new GetDeploymentPreviewParams object
// with the default values initialized.
func NewGetDeploymentPreviewParams() *GetDeploymentPreviewParams {
	var ()
	return &GetDeploymentPreviewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentPreviewParamsWithTimeout creates a new GetDeploymentPreviewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentPreviewParamsWithTimeout(timeout time.Duration) *GetDeploymentPreviewParams {
	var ()
	return &GetDeploymentPreviewParams{

		timeout: timeout,
	}
}

// NewGetDeploymentPreviewParamsWithContext creates a new GetDeploymentPreviewParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentPreviewParamsWithContext(ctx context.Context) *GetDeploymentPreviewParams {
	var ()
	return &GetDeploymentPreviewParams{

		Context: ctx,
	}
}

// NewGetDeploymentPreviewParamsWithHTTPClient creates a new GetDeploymentPreviewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentPreviewParamsWithHTTPClient(client *http.Client) *GetDeploymentPreviewParams {
	var ()
	return &GetDeploymentPreviewParams{
		HTTPClient: client,
	}
}

/*GetDeploymentPreviewParams contains all the parameters to send to the API endpoint
for the get deployment preview operation typically these are written to a http.Request
*/
type GetDeploymentPreviewParams struct {

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

// WithTimeout adds the timeout to the get deployment preview params
func (o *GetDeploymentPreviewParams) WithTimeout(timeout time.Duration) *GetDeploymentPreviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment preview params
func (o *GetDeploymentPreviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment preview params
func (o *GetDeploymentPreviewParams) WithContext(ctx context.Context) *GetDeploymentPreviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment preview params
func (o *GetDeploymentPreviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment preview params
func (o *GetDeploymentPreviewParams) WithHTTPClient(client *http.Client) *GetDeploymentPreviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment preview params
func (o *GetDeploymentPreviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the get deployment preview params
func (o *GetDeploymentPreviewParams) WithEnvironment(environment string) *GetDeploymentPreviewParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the get deployment preview params
func (o *GetDeploymentPreviewParams) SetEnvironment(environment string) {
	o.Environment = environment
}

// WithID adds the id to the get deployment preview params
func (o *GetDeploymentPreviewParams) WithID(id string) *GetDeploymentPreviewParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get deployment preview params
func (o *GetDeploymentPreviewParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentPreviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
