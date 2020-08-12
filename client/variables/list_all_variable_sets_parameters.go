// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewListAllVariableSetsParams creates a new ListAllVariableSetsParams object
// with the default values initialized.
func NewListAllVariableSetsParams() *ListAllVariableSetsParams {

	return &ListAllVariableSetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllVariableSetsParamsWithTimeout creates a new ListAllVariableSetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllVariableSetsParamsWithTimeout(timeout time.Duration) *ListAllVariableSetsParams {

	return &ListAllVariableSetsParams{

		timeout: timeout,
	}
}

// NewListAllVariableSetsParamsWithContext creates a new ListAllVariableSetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllVariableSetsParamsWithContext(ctx context.Context) *ListAllVariableSetsParams {

	return &ListAllVariableSetsParams{

		Context: ctx,
	}
}

// NewListAllVariableSetsParamsWithHTTPClient creates a new ListAllVariableSetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllVariableSetsParamsWithHTTPClient(client *http.Client) *ListAllVariableSetsParams {

	return &ListAllVariableSetsParams{
		HTTPClient: client,
	}
}

/*ListAllVariableSetsParams contains all the parameters to send to the API endpoint
for the list all variable sets operation typically these are written to a http.Request
*/
type ListAllVariableSetsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all variable sets params
func (o *ListAllVariableSetsParams) WithTimeout(timeout time.Duration) *ListAllVariableSetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all variable sets params
func (o *ListAllVariableSetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all variable sets params
func (o *ListAllVariableSetsParams) WithContext(ctx context.Context) *ListAllVariableSetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all variable sets params
func (o *ListAllVariableSetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all variable sets params
func (o *ListAllVariableSetsParams) WithHTTPClient(client *http.Client) *ListAllVariableSetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all variable sets params
func (o *ListAllVariableSetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllVariableSetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}