// Code generated by go-swagger; DO NOT EDIT.

package configuration

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

// NewUpdateConfigurationItemParams creates a new UpdateConfigurationItemParams object
// with the default values initialized.
func NewUpdateConfigurationItemParams() *UpdateConfigurationItemParams {
	var ()
	return &UpdateConfigurationItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConfigurationItemParamsWithTimeout creates a new UpdateConfigurationItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateConfigurationItemParamsWithTimeout(timeout time.Duration) *UpdateConfigurationItemParams {
	var ()
	return &UpdateConfigurationItemParams{

		timeout: timeout,
	}
}

// NewUpdateConfigurationItemParamsWithContext creates a new UpdateConfigurationItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateConfigurationItemParamsWithContext(ctx context.Context) *UpdateConfigurationItemParams {
	var ()
	return &UpdateConfigurationItemParams{

		Context: ctx,
	}
}

// NewUpdateConfigurationItemParamsWithHTTPClient creates a new UpdateConfigurationItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateConfigurationItemParamsWithHTTPClient(client *http.Client) *UpdateConfigurationItemParams {
	var ()
	return &UpdateConfigurationItemParams{
		HTTPClient: client,
	}
}

/*UpdateConfigurationItemParams contains all the parameters to send to the API endpoint
for the update configuration item operation typically these are written to a http.Request
*/
type UpdateConfigurationItemParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update configuration item params
func (o *UpdateConfigurationItemParams) WithTimeout(timeout time.Duration) *UpdateConfigurationItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update configuration item params
func (o *UpdateConfigurationItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update configuration item params
func (o *UpdateConfigurationItemParams) WithContext(ctx context.Context) *UpdateConfigurationItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update configuration item params
func (o *UpdateConfigurationItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update configuration item params
func (o *UpdateConfigurationItemParams) WithHTTPClient(client *http.Client) *UpdateConfigurationItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update configuration item params
func (o *UpdateConfigurationItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update configuration item params
func (o *UpdateConfigurationItemParams) WithID(id string) *UpdateConfigurationItemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update configuration item params
func (o *UpdateConfigurationItemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConfigurationItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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