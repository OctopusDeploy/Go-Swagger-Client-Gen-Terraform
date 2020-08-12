// Code generated by go-swagger; DO NOT EDIT.

package dashboard_configurations

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

// NewGetDashboardConfigurationParams creates a new GetDashboardConfigurationParams object
// with the default values initialized.
func NewGetDashboardConfigurationParams() *GetDashboardConfigurationParams {

	return &GetDashboardConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardConfigurationParamsWithTimeout creates a new GetDashboardConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDashboardConfigurationParamsWithTimeout(timeout time.Duration) *GetDashboardConfigurationParams {

	return &GetDashboardConfigurationParams{

		timeout: timeout,
	}
}

// NewGetDashboardConfigurationParamsWithContext creates a new GetDashboardConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDashboardConfigurationParamsWithContext(ctx context.Context) *GetDashboardConfigurationParams {

	return &GetDashboardConfigurationParams{

		Context: ctx,
	}
}

// NewGetDashboardConfigurationParamsWithHTTPClient creates a new GetDashboardConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDashboardConfigurationParamsWithHTTPClient(client *http.Client) *GetDashboardConfigurationParams {

	return &GetDashboardConfigurationParams{
		HTTPClient: client,
	}
}

/*GetDashboardConfigurationParams contains all the parameters to send to the API endpoint
for the get dashboard configuration operation typically these are written to a http.Request
*/
type GetDashboardConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get dashboard configuration params
func (o *GetDashboardConfigurationParams) WithTimeout(timeout time.Duration) *GetDashboardConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard configuration params
func (o *GetDashboardConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard configuration params
func (o *GetDashboardConfigurationParams) WithContext(ctx context.Context) *GetDashboardConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard configuration params
func (o *GetDashboardConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard configuration params
func (o *GetDashboardConfigurationParams) WithHTTPClient(client *http.Client) *GetDashboardConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard configuration params
func (o *GetDashboardConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
