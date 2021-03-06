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

// NewUpdateFeaturesConfigurationParams creates a new UpdateFeaturesConfigurationParams object
// with the default values initialized.
func NewUpdateFeaturesConfigurationParams() *UpdateFeaturesConfigurationParams {

	return &UpdateFeaturesConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFeaturesConfigurationParamsWithTimeout creates a new UpdateFeaturesConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFeaturesConfigurationParamsWithTimeout(timeout time.Duration) *UpdateFeaturesConfigurationParams {

	return &UpdateFeaturesConfigurationParams{

		timeout: timeout,
	}
}

// NewUpdateFeaturesConfigurationParamsWithContext creates a new UpdateFeaturesConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFeaturesConfigurationParamsWithContext(ctx context.Context) *UpdateFeaturesConfigurationParams {

	return &UpdateFeaturesConfigurationParams{

		Context: ctx,
	}
}

// NewUpdateFeaturesConfigurationParamsWithHTTPClient creates a new UpdateFeaturesConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFeaturesConfigurationParamsWithHTTPClient(client *http.Client) *UpdateFeaturesConfigurationParams {

	return &UpdateFeaturesConfigurationParams{
		HTTPClient: client,
	}
}

/*UpdateFeaturesConfigurationParams contains all the parameters to send to the API endpoint
for the update features configuration operation typically these are written to a http.Request
*/
type UpdateFeaturesConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update features configuration params
func (o *UpdateFeaturesConfigurationParams) WithTimeout(timeout time.Duration) *UpdateFeaturesConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update features configuration params
func (o *UpdateFeaturesConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update features configuration params
func (o *UpdateFeaturesConfigurationParams) WithContext(ctx context.Context) *UpdateFeaturesConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update features configuration params
func (o *UpdateFeaturesConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update features configuration params
func (o *UpdateFeaturesConfigurationParams) WithHTTPClient(client *http.Client) *UpdateFeaturesConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update features configuration params
func (o *UpdateFeaturesConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFeaturesConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
