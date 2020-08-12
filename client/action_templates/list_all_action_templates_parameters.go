// Code generated by go-swagger; DO NOT EDIT.

package action_templates

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

// NewListAllActionTemplatesParams creates a new ListAllActionTemplatesParams object
// with the default values initialized.
func NewListAllActionTemplatesParams() *ListAllActionTemplatesParams {

	return &ListAllActionTemplatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllActionTemplatesParamsWithTimeout creates a new ListAllActionTemplatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllActionTemplatesParamsWithTimeout(timeout time.Duration) *ListAllActionTemplatesParams {

	return &ListAllActionTemplatesParams{

		timeout: timeout,
	}
}

// NewListAllActionTemplatesParamsWithContext creates a new ListAllActionTemplatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllActionTemplatesParamsWithContext(ctx context.Context) *ListAllActionTemplatesParams {

	return &ListAllActionTemplatesParams{

		Context: ctx,
	}
}

// NewListAllActionTemplatesParamsWithHTTPClient creates a new ListAllActionTemplatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllActionTemplatesParamsWithHTTPClient(client *http.Client) *ListAllActionTemplatesParams {

	return &ListAllActionTemplatesParams{
		HTTPClient: client,
	}
}

/*ListAllActionTemplatesParams contains all the parameters to send to the API endpoint
for the list all action templates operation typically these are written to a http.Request
*/
type ListAllActionTemplatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all action templates params
func (o *ListAllActionTemplatesParams) WithTimeout(timeout time.Duration) *ListAllActionTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all action templates params
func (o *ListAllActionTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all action templates params
func (o *ListAllActionTemplatesParams) WithContext(ctx context.Context) *ListAllActionTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all action templates params
func (o *ListAllActionTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all action templates params
func (o *ListAllActionTemplatesParams) WithHTTPClient(client *http.Client) *ListAllActionTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all action templates params
func (o *ListAllActionTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllActionTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
