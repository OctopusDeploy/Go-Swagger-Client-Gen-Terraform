// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

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

// NewListAllLifecyclesParams creates a new ListAllLifecyclesParams object
// with the default values initialized.
func NewListAllLifecyclesParams() *ListAllLifecyclesParams {

	return &ListAllLifecyclesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllLifecyclesParamsWithTimeout creates a new ListAllLifecyclesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllLifecyclesParamsWithTimeout(timeout time.Duration) *ListAllLifecyclesParams {

	return &ListAllLifecyclesParams{

		timeout: timeout,
	}
}

// NewListAllLifecyclesParamsWithContext creates a new ListAllLifecyclesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllLifecyclesParamsWithContext(ctx context.Context) *ListAllLifecyclesParams {

	return &ListAllLifecyclesParams{

		Context: ctx,
	}
}

// NewListAllLifecyclesParamsWithHTTPClient creates a new ListAllLifecyclesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllLifecyclesParamsWithHTTPClient(client *http.Client) *ListAllLifecyclesParams {

	return &ListAllLifecyclesParams{
		HTTPClient: client,
	}
}

/*ListAllLifecyclesParams contains all the parameters to send to the API endpoint
for the list all lifecycles operation typically these are written to a http.Request
*/
type ListAllLifecyclesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all lifecycles params
func (o *ListAllLifecyclesParams) WithTimeout(timeout time.Duration) *ListAllLifecyclesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all lifecycles params
func (o *ListAllLifecyclesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all lifecycles params
func (o *ListAllLifecyclesParams) WithContext(ctx context.Context) *ListAllLifecyclesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all lifecycles params
func (o *ListAllLifecyclesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all lifecycles params
func (o *ListAllLifecyclesParams) WithHTTPClient(client *http.Client) *ListAllLifecyclesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all lifecycles params
func (o *ListAllLifecyclesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllLifecyclesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}