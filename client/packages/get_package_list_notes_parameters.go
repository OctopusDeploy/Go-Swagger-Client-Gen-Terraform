// Code generated by go-swagger; DO NOT EDIT.

package packages

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

// NewGetPackageListNotesParams creates a new GetPackageListNotesParams object
// with the default values initialized.
func NewGetPackageListNotesParams() *GetPackageListNotesParams {

	return &GetPackageListNotesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackageListNotesParamsWithTimeout creates a new GetPackageListNotesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPackageListNotesParamsWithTimeout(timeout time.Duration) *GetPackageListNotesParams {

	return &GetPackageListNotesParams{

		timeout: timeout,
	}
}

// NewGetPackageListNotesParamsWithContext creates a new GetPackageListNotesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPackageListNotesParamsWithContext(ctx context.Context) *GetPackageListNotesParams {

	return &GetPackageListNotesParams{

		Context: ctx,
	}
}

// NewGetPackageListNotesParamsWithHTTPClient creates a new GetPackageListNotesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPackageListNotesParamsWithHTTPClient(client *http.Client) *GetPackageListNotesParams {

	return &GetPackageListNotesParams{
		HTTPClient: client,
	}
}

/*GetPackageListNotesParams contains all the parameters to send to the API endpoint
for the get package list notes operation typically these are written to a http.Request
*/
type GetPackageListNotesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get package list notes params
func (o *GetPackageListNotesParams) WithTimeout(timeout time.Duration) *GetPackageListNotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get package list notes params
func (o *GetPackageListNotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get package list notes params
func (o *GetPackageListNotesParams) WithContext(ctx context.Context) *GetPackageListNotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get package list notes params
func (o *GetPackageListNotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get package list notes params
func (o *GetPackageListNotesParams) WithHTTPClient(client *http.Client) *GetPackageListNotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get package list notes params
func (o *GetPackageListNotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackageListNotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
