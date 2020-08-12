// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewGetListEventsParams creates a new GetListEventsParams object
// with the default values initialized.
func NewGetListEventsParams() *GetListEventsParams {

	return &GetListEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetListEventsParamsWithTimeout creates a new GetListEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetListEventsParamsWithTimeout(timeout time.Duration) *GetListEventsParams {

	return &GetListEventsParams{

		timeout: timeout,
	}
}

// NewGetListEventsParamsWithContext creates a new GetListEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetListEventsParamsWithContext(ctx context.Context) *GetListEventsParams {

	return &GetListEventsParams{

		Context: ctx,
	}
}

// NewGetListEventsParamsWithHTTPClient creates a new GetListEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetListEventsParamsWithHTTPClient(client *http.Client) *GetListEventsParams {

	return &GetListEventsParams{
		HTTPClient: client,
	}
}

/*GetListEventsParams contains all the parameters to send to the API endpoint
for the get list events operation typically these are written to a http.Request
*/
type GetListEventsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get list events params
func (o *GetListEventsParams) WithTimeout(timeout time.Duration) *GetListEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get list events params
func (o *GetListEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get list events params
func (o *GetListEventsParams) WithContext(ctx context.Context) *GetListEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get list events params
func (o *GetListEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get list events params
func (o *GetListEventsParams) WithHTTPClient(client *http.Client) *GetListEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get list events params
func (o *GetListEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetListEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}