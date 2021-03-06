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

// NewGetTenantVariableSetTestParams creates a new GetTenantVariableSetTestParams object
// with the default values initialized.
func NewGetTenantVariableSetTestParams() *GetTenantVariableSetTestParams {

	return &GetTenantVariableSetTestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantVariableSetTestParamsWithTimeout creates a new GetTenantVariableSetTestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenantVariableSetTestParamsWithTimeout(timeout time.Duration) *GetTenantVariableSetTestParams {

	return &GetTenantVariableSetTestParams{

		timeout: timeout,
	}
}

// NewGetTenantVariableSetTestParamsWithContext creates a new GetTenantVariableSetTestParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenantVariableSetTestParamsWithContext(ctx context.Context) *GetTenantVariableSetTestParams {

	return &GetTenantVariableSetTestParams{

		Context: ctx,
	}
}

// NewGetTenantVariableSetTestParamsWithHTTPClient creates a new GetTenantVariableSetTestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenantVariableSetTestParamsWithHTTPClient(client *http.Client) *GetTenantVariableSetTestParams {

	return &GetTenantVariableSetTestParams{
		HTTPClient: client,
	}
}

/*GetTenantVariableSetTestParams contains all the parameters to send to the API endpoint
for the get tenant variable set test operation typically these are written to a http.Request
*/
type GetTenantVariableSetTestParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tenant variable set test params
func (o *GetTenantVariableSetTestParams) WithTimeout(timeout time.Duration) *GetTenantVariableSetTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenant variable set test params
func (o *GetTenantVariableSetTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenant variable set test params
func (o *GetTenantVariableSetTestParams) WithContext(ctx context.Context) *GetTenantVariableSetTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenant variable set test params
func (o *GetTenantVariableSetTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenant variable set test params
func (o *GetTenantVariableSetTestParams) WithHTTPClient(client *http.Client) *GetTenantVariableSetTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenant variable set test params
func (o *GetTenantVariableSetTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantVariableSetTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
