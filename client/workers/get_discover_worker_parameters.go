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

// NewGetDiscoverWorkerParams creates a new GetDiscoverWorkerParams object
// with the default values initialized.
func NewGetDiscoverWorkerParams() *GetDiscoverWorkerParams {

	return &GetDiscoverWorkerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDiscoverWorkerParamsWithTimeout creates a new GetDiscoverWorkerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDiscoverWorkerParamsWithTimeout(timeout time.Duration) *GetDiscoverWorkerParams {

	return &GetDiscoverWorkerParams{

		timeout: timeout,
	}
}

// NewGetDiscoverWorkerParamsWithContext creates a new GetDiscoverWorkerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDiscoverWorkerParamsWithContext(ctx context.Context) *GetDiscoverWorkerParams {

	return &GetDiscoverWorkerParams{

		Context: ctx,
	}
}

// NewGetDiscoverWorkerParamsWithHTTPClient creates a new GetDiscoverWorkerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDiscoverWorkerParamsWithHTTPClient(client *http.Client) *GetDiscoverWorkerParams {

	return &GetDiscoverWorkerParams{
		HTTPClient: client,
	}
}

/*GetDiscoverWorkerParams contains all the parameters to send to the API endpoint
for the get discover worker operation typically these are written to a http.Request
*/
type GetDiscoverWorkerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get discover worker params
func (o *GetDiscoverWorkerParams) WithTimeout(timeout time.Duration) *GetDiscoverWorkerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get discover worker params
func (o *GetDiscoverWorkerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get discover worker params
func (o *GetDiscoverWorkerParams) WithContext(ctx context.Context) *GetDiscoverWorkerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get discover worker params
func (o *GetDiscoverWorkerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get discover worker params
func (o *GetDiscoverWorkerParams) WithHTTPClient(client *http.Client) *GetDiscoverWorkerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get discover worker params
func (o *GetDiscoverWorkerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDiscoverWorkerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
