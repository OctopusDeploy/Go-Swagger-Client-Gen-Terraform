// Code generated by go-swagger; DO NOT EDIT.

package tenant_variables

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

// NewGetTenantVariablesAllParams creates a new GetTenantVariablesAllParams object
// with the default values initialized.
func NewGetTenantVariablesAllParams() *GetTenantVariablesAllParams {

	return &GetTenantVariablesAllParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantVariablesAllParamsWithTimeout creates a new GetTenantVariablesAllParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenantVariablesAllParamsWithTimeout(timeout time.Duration) *GetTenantVariablesAllParams {

	return &GetTenantVariablesAllParams{

		timeout: timeout,
	}
}

// NewGetTenantVariablesAllParamsWithContext creates a new GetTenantVariablesAllParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenantVariablesAllParamsWithContext(ctx context.Context) *GetTenantVariablesAllParams {

	return &GetTenantVariablesAllParams{

		Context: ctx,
	}
}

// NewGetTenantVariablesAllParamsWithHTTPClient creates a new GetTenantVariablesAllParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenantVariablesAllParamsWithHTTPClient(client *http.Client) *GetTenantVariablesAllParams {

	return &GetTenantVariablesAllParams{
		HTTPClient: client,
	}
}

/*GetTenantVariablesAllParams contains all the parameters to send to the API endpoint
for the get tenant variables all operation typically these are written to a http.Request
*/
type GetTenantVariablesAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tenant variables all params
func (o *GetTenantVariablesAllParams) WithTimeout(timeout time.Duration) *GetTenantVariablesAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenant variables all params
func (o *GetTenantVariablesAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenant variables all params
func (o *GetTenantVariablesAllParams) WithContext(ctx context.Context) *GetTenantVariablesAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenant variables all params
func (o *GetTenantVariablesAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenant variables all params
func (o *GetTenantVariablesAllParams) WithHTTPClient(client *http.Client) *GetTenantVariablesAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenant variables all params
func (o *GetTenantVariablesAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantVariablesAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
