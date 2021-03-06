// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetUserGetCurrentParams creates a new GetUserGetCurrentParams object
// with the default values initialized.
func NewGetUserGetCurrentParams() *GetUserGetCurrentParams {

	return &GetUserGetCurrentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserGetCurrentParamsWithTimeout creates a new GetUserGetCurrentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserGetCurrentParamsWithTimeout(timeout time.Duration) *GetUserGetCurrentParams {

	return &GetUserGetCurrentParams{

		timeout: timeout,
	}
}

// NewGetUserGetCurrentParamsWithContext creates a new GetUserGetCurrentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserGetCurrentParamsWithContext(ctx context.Context) *GetUserGetCurrentParams {

	return &GetUserGetCurrentParams{

		Context: ctx,
	}
}

// NewGetUserGetCurrentParamsWithHTTPClient creates a new GetUserGetCurrentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserGetCurrentParamsWithHTTPClient(client *http.Client) *GetUserGetCurrentParams {

	return &GetUserGetCurrentParams{
		HTTPClient: client,
	}
}

/*GetUserGetCurrentParams contains all the parameters to send to the API endpoint
for the get user get current operation typically these are written to a http.Request
*/
type GetUserGetCurrentParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user get current params
func (o *GetUserGetCurrentParams) WithTimeout(timeout time.Duration) *GetUserGetCurrentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user get current params
func (o *GetUserGetCurrentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user get current params
func (o *GetUserGetCurrentParams) WithContext(ctx context.Context) *GetUserGetCurrentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user get current params
func (o *GetUserGetCurrentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user get current params
func (o *GetUserGetCurrentParams) WithHTTPClient(client *http.Client) *GetUserGetCurrentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user get current params
func (o *GetUserGetCurrentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserGetCurrentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
