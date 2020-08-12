// Code generated by go-swagger; DO NOT EDIT.

package server_status

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

// NewCreateSystemGarbageCollectParams creates a new CreateSystemGarbageCollectParams object
// with the default values initialized.
func NewCreateSystemGarbageCollectParams() *CreateSystemGarbageCollectParams {

	return &CreateSystemGarbageCollectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSystemGarbageCollectParamsWithTimeout creates a new CreateSystemGarbageCollectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSystemGarbageCollectParamsWithTimeout(timeout time.Duration) *CreateSystemGarbageCollectParams {

	return &CreateSystemGarbageCollectParams{

		timeout: timeout,
	}
}

// NewCreateSystemGarbageCollectParamsWithContext creates a new CreateSystemGarbageCollectParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSystemGarbageCollectParamsWithContext(ctx context.Context) *CreateSystemGarbageCollectParams {

	return &CreateSystemGarbageCollectParams{

		Context: ctx,
	}
}

// NewCreateSystemGarbageCollectParamsWithHTTPClient creates a new CreateSystemGarbageCollectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSystemGarbageCollectParamsWithHTTPClient(client *http.Client) *CreateSystemGarbageCollectParams {

	return &CreateSystemGarbageCollectParams{
		HTTPClient: client,
	}
}

/*CreateSystemGarbageCollectParams contains all the parameters to send to the API endpoint
for the create system garbage collect operation typically these are written to a http.Request
*/
type CreateSystemGarbageCollectParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create system garbage collect params
func (o *CreateSystemGarbageCollectParams) WithTimeout(timeout time.Duration) *CreateSystemGarbageCollectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create system garbage collect params
func (o *CreateSystemGarbageCollectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create system garbage collect params
func (o *CreateSystemGarbageCollectParams) WithContext(ctx context.Context) *CreateSystemGarbageCollectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create system garbage collect params
func (o *CreateSystemGarbageCollectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create system garbage collect params
func (o *CreateSystemGarbageCollectParams) WithHTTPClient(client *http.Client) *CreateSystemGarbageCollectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create system garbage collect params
func (o *CreateSystemGarbageCollectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSystemGarbageCollectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
