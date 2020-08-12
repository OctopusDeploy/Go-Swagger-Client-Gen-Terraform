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

// NewGetVariableNamesListParams creates a new GetVariableNamesListParams object
// with the default values initialized.
func NewGetVariableNamesListParams() *GetVariableNamesListParams {

	return &GetVariableNamesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVariableNamesListParamsWithTimeout creates a new GetVariableNamesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVariableNamesListParamsWithTimeout(timeout time.Duration) *GetVariableNamesListParams {

	return &GetVariableNamesListParams{

		timeout: timeout,
	}
}

// NewGetVariableNamesListParamsWithContext creates a new GetVariableNamesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVariableNamesListParamsWithContext(ctx context.Context) *GetVariableNamesListParams {

	return &GetVariableNamesListParams{

		Context: ctx,
	}
}

// NewGetVariableNamesListParamsWithHTTPClient creates a new GetVariableNamesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVariableNamesListParamsWithHTTPClient(client *http.Client) *GetVariableNamesListParams {

	return &GetVariableNamesListParams{
		HTTPClient: client,
	}
}

/*GetVariableNamesListParams contains all the parameters to send to the API endpoint
for the get variable names list operation typically these are written to a http.Request
*/
type GetVariableNamesListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get variable names list params
func (o *GetVariableNamesListParams) WithTimeout(timeout time.Duration) *GetVariableNamesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get variable names list params
func (o *GetVariableNamesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get variable names list params
func (o *GetVariableNamesListParams) WithContext(ctx context.Context) *GetVariableNamesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get variable names list params
func (o *GetVariableNamesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get variable names list params
func (o *GetVariableNamesListParams) WithHTTPClient(client *http.Client) *GetVariableNamesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get variable names list params
func (o *GetVariableNamesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVariableNamesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
