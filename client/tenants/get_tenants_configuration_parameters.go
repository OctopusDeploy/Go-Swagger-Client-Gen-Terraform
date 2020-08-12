// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewGetTenantsConfigurationParams creates a new GetTenantsConfigurationParams object
// with the default values initialized.
func NewGetTenantsConfigurationParams() *GetTenantsConfigurationParams {

	return &GetTenantsConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantsConfigurationParamsWithTimeout creates a new GetTenantsConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenantsConfigurationParamsWithTimeout(timeout time.Duration) *GetTenantsConfigurationParams {

	return &GetTenantsConfigurationParams{

		timeout: timeout,
	}
}

// NewGetTenantsConfigurationParamsWithContext creates a new GetTenantsConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenantsConfigurationParamsWithContext(ctx context.Context) *GetTenantsConfigurationParams {

	return &GetTenantsConfigurationParams{

		Context: ctx,
	}
}

// NewGetTenantsConfigurationParamsWithHTTPClient creates a new GetTenantsConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenantsConfigurationParamsWithHTTPClient(client *http.Client) *GetTenantsConfigurationParams {

	return &GetTenantsConfigurationParams{
		HTTPClient: client,
	}
}

/*GetTenantsConfigurationParams contains all the parameters to send to the API endpoint
for the get tenants configuration operation typically these are written to a http.Request
*/
type GetTenantsConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tenants configuration params
func (o *GetTenantsConfigurationParams) WithTimeout(timeout time.Duration) *GetTenantsConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenants configuration params
func (o *GetTenantsConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenants configuration params
func (o *GetTenantsConfigurationParams) WithContext(ctx context.Context) *GetTenantsConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenants configuration params
func (o *GetTenantsConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenants configuration params
func (o *GetTenantsConfigurationParams) WithHTTPClient(client *http.Client) *GetTenantsConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenants configuration params
func (o *GetTenantsConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantsConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}