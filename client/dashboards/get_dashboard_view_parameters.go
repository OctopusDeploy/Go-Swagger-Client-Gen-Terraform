// Code generated by go-swagger; DO NOT EDIT.

package dashboards

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

// NewGetDashboardViewParams creates a new GetDashboardViewParams object
// with the default values initialized.
func NewGetDashboardViewParams() *GetDashboardViewParams {

	return &GetDashboardViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardViewParamsWithTimeout creates a new GetDashboardViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDashboardViewParamsWithTimeout(timeout time.Duration) *GetDashboardViewParams {

	return &GetDashboardViewParams{

		timeout: timeout,
	}
}

// NewGetDashboardViewParamsWithContext creates a new GetDashboardViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDashboardViewParamsWithContext(ctx context.Context) *GetDashboardViewParams {

	return &GetDashboardViewParams{

		Context: ctx,
	}
}

// NewGetDashboardViewParamsWithHTTPClient creates a new GetDashboardViewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDashboardViewParamsWithHTTPClient(client *http.Client) *GetDashboardViewParams {

	return &GetDashboardViewParams{
		HTTPClient: client,
	}
}

/*GetDashboardViewParams contains all the parameters to send to the API endpoint
for the get dashboard view operation typically these are written to a http.Request
*/
type GetDashboardViewParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get dashboard view params
func (o *GetDashboardViewParams) WithTimeout(timeout time.Duration) *GetDashboardViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard view params
func (o *GetDashboardViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard view params
func (o *GetDashboardViewParams) WithContext(ctx context.Context) *GetDashboardViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard view params
func (o *GetDashboardViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard view params
func (o *GetDashboardViewParams) WithHTTPClient(client *http.Client) *GetDashboardViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard view params
func (o *GetDashboardViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
