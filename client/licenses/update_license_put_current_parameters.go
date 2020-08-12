// Code generated by go-swagger; DO NOT EDIT.

package licenses

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

// NewUpdateLicensePutCurrentParams creates a new UpdateLicensePutCurrentParams object
// with the default values initialized.
func NewUpdateLicensePutCurrentParams() *UpdateLicensePutCurrentParams {

	return &UpdateLicensePutCurrentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLicensePutCurrentParamsWithTimeout creates a new UpdateLicensePutCurrentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateLicensePutCurrentParamsWithTimeout(timeout time.Duration) *UpdateLicensePutCurrentParams {

	return &UpdateLicensePutCurrentParams{

		timeout: timeout,
	}
}

// NewUpdateLicensePutCurrentParamsWithContext creates a new UpdateLicensePutCurrentParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateLicensePutCurrentParamsWithContext(ctx context.Context) *UpdateLicensePutCurrentParams {

	return &UpdateLicensePutCurrentParams{

		Context: ctx,
	}
}

// NewUpdateLicensePutCurrentParamsWithHTTPClient creates a new UpdateLicensePutCurrentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateLicensePutCurrentParamsWithHTTPClient(client *http.Client) *UpdateLicensePutCurrentParams {

	return &UpdateLicensePutCurrentParams{
		HTTPClient: client,
	}
}

/*UpdateLicensePutCurrentParams contains all the parameters to send to the API endpoint
for the update license put current operation typically these are written to a http.Request
*/
type UpdateLicensePutCurrentParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update license put current params
func (o *UpdateLicensePutCurrentParams) WithTimeout(timeout time.Duration) *UpdateLicensePutCurrentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update license put current params
func (o *UpdateLicensePutCurrentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update license put current params
func (o *UpdateLicensePutCurrentParams) WithContext(ctx context.Context) *UpdateLicensePutCurrentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update license put current params
func (o *UpdateLicensePutCurrentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update license put current params
func (o *UpdateLicensePutCurrentParams) WithHTTPClient(client *http.Client) *UpdateLicensePutCurrentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update license put current params
func (o *UpdateLicensePutCurrentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLicensePutCurrentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
