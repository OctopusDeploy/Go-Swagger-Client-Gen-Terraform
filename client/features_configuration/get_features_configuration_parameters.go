// Code generated by go-swagger; DO NOT EDIT.

package features_configuration

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

// NewGetFeaturesConfigurationParams creates a new GetFeaturesConfigurationParams object
// with the default values initialized.
func NewGetFeaturesConfigurationParams() *GetFeaturesConfigurationParams {

	return &GetFeaturesConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFeaturesConfigurationParamsWithTimeout creates a new GetFeaturesConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFeaturesConfigurationParamsWithTimeout(timeout time.Duration) *GetFeaturesConfigurationParams {

	return &GetFeaturesConfigurationParams{

		timeout: timeout,
	}
}

// NewGetFeaturesConfigurationParamsWithContext creates a new GetFeaturesConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFeaturesConfigurationParamsWithContext(ctx context.Context) *GetFeaturesConfigurationParams {

	return &GetFeaturesConfigurationParams{

		Context: ctx,
	}
}

// NewGetFeaturesConfigurationParamsWithHTTPClient creates a new GetFeaturesConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFeaturesConfigurationParamsWithHTTPClient(client *http.Client) *GetFeaturesConfigurationParams {

	return &GetFeaturesConfigurationParams{
		HTTPClient: client,
	}
}

/*GetFeaturesConfigurationParams contains all the parameters to send to the API endpoint
for the get features configuration operation typically these are written to a http.Request
*/
type GetFeaturesConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get features configuration params
func (o *GetFeaturesConfigurationParams) WithTimeout(timeout time.Duration) *GetFeaturesConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get features configuration params
func (o *GetFeaturesConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get features configuration params
func (o *GetFeaturesConfigurationParams) WithContext(ctx context.Context) *GetFeaturesConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get features configuration params
func (o *GetFeaturesConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get features configuration params
func (o *GetFeaturesConfigurationParams) WithHTTPClient(client *http.Client) *GetFeaturesConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get features configuration params
func (o *GetFeaturesConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFeaturesConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
