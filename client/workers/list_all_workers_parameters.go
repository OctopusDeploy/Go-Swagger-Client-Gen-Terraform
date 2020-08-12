// Code generated by go-swagger; DO NOT EDIT.

package workers

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

// NewListAllWorkersParams creates a new ListAllWorkersParams object
// with the default values initialized.
func NewListAllWorkersParams() *ListAllWorkersParams {

	return &ListAllWorkersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllWorkersParamsWithTimeout creates a new ListAllWorkersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllWorkersParamsWithTimeout(timeout time.Duration) *ListAllWorkersParams {

	return &ListAllWorkersParams{

		timeout: timeout,
	}
}

// NewListAllWorkersParamsWithContext creates a new ListAllWorkersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllWorkersParamsWithContext(ctx context.Context) *ListAllWorkersParams {

	return &ListAllWorkersParams{

		Context: ctx,
	}
}

// NewListAllWorkersParamsWithHTTPClient creates a new ListAllWorkersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllWorkersParamsWithHTTPClient(client *http.Client) *ListAllWorkersParams {

	return &ListAllWorkersParams{
		HTTPClient: client,
	}
}

/*ListAllWorkersParams contains all the parameters to send to the API endpoint
for the list all workers operation typically these are written to a http.Request
*/
type ListAllWorkersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all workers params
func (o *ListAllWorkersParams) WithTimeout(timeout time.Duration) *ListAllWorkersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all workers params
func (o *ListAllWorkersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all workers params
func (o *ListAllWorkersParams) WithContext(ctx context.Context) *ListAllWorkersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all workers params
func (o *ListAllWorkersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all workers params
func (o *ListAllWorkersParams) WithHTTPClient(client *http.Client) *ListAllWorkersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all workers params
func (o *ListAllWorkersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllWorkersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}