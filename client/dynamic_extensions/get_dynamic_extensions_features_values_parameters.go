// Code generated by go-swagger; DO NOT EDIT.

package dynamic_extensions

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

// NewGetDynamicExtensionsFeaturesValuesParams creates a new GetDynamicExtensionsFeaturesValuesParams object
// with the default values initialized.
func NewGetDynamicExtensionsFeaturesValuesParams() *GetDynamicExtensionsFeaturesValuesParams {

	return &GetDynamicExtensionsFeaturesValuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDynamicExtensionsFeaturesValuesParamsWithTimeout creates a new GetDynamicExtensionsFeaturesValuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDynamicExtensionsFeaturesValuesParamsWithTimeout(timeout time.Duration) *GetDynamicExtensionsFeaturesValuesParams {

	return &GetDynamicExtensionsFeaturesValuesParams{

		timeout: timeout,
	}
}

// NewGetDynamicExtensionsFeaturesValuesParamsWithContext creates a new GetDynamicExtensionsFeaturesValuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDynamicExtensionsFeaturesValuesParamsWithContext(ctx context.Context) *GetDynamicExtensionsFeaturesValuesParams {

	return &GetDynamicExtensionsFeaturesValuesParams{

		Context: ctx,
	}
}

// NewGetDynamicExtensionsFeaturesValuesParamsWithHTTPClient creates a new GetDynamicExtensionsFeaturesValuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDynamicExtensionsFeaturesValuesParamsWithHTTPClient(client *http.Client) *GetDynamicExtensionsFeaturesValuesParams {

	return &GetDynamicExtensionsFeaturesValuesParams{
		HTTPClient: client,
	}
}

/*GetDynamicExtensionsFeaturesValuesParams contains all the parameters to send to the API endpoint
for the get dynamic extensions features values operation typically these are written to a http.Request
*/
type GetDynamicExtensionsFeaturesValuesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get dynamic extensions features values params
func (o *GetDynamicExtensionsFeaturesValuesParams) WithTimeout(timeout time.Duration) *GetDynamicExtensionsFeaturesValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dynamic extensions features values params
func (o *GetDynamicExtensionsFeaturesValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dynamic extensions features values params
func (o *GetDynamicExtensionsFeaturesValuesParams) WithContext(ctx context.Context) *GetDynamicExtensionsFeaturesValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dynamic extensions features values params
func (o *GetDynamicExtensionsFeaturesValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dynamic extensions features values params
func (o *GetDynamicExtensionsFeaturesValuesParams) WithHTTPClient(client *http.Client) *GetDynamicExtensionsFeaturesValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dynamic extensions features values params
func (o *GetDynamicExtensionsFeaturesValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDynamicExtensionsFeaturesValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
