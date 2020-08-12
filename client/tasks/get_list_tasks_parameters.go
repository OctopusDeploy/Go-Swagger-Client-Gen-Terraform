// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewGetListTasksParams creates a new GetListTasksParams object
// with the default values initialized.
func NewGetListTasksParams() *GetListTasksParams {

	return &GetListTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetListTasksParamsWithTimeout creates a new GetListTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetListTasksParamsWithTimeout(timeout time.Duration) *GetListTasksParams {

	return &GetListTasksParams{

		timeout: timeout,
	}
}

// NewGetListTasksParamsWithContext creates a new GetListTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetListTasksParamsWithContext(ctx context.Context) *GetListTasksParams {

	return &GetListTasksParams{

		Context: ctx,
	}
}

// NewGetListTasksParamsWithHTTPClient creates a new GetListTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetListTasksParamsWithHTTPClient(client *http.Client) *GetListTasksParams {

	return &GetListTasksParams{
		HTTPClient: client,
	}
}

/*GetListTasksParams contains all the parameters to send to the API endpoint
for the get list tasks operation typically these are written to a http.Request
*/
type GetListTasksParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get list tasks params
func (o *GetListTasksParams) WithTimeout(timeout time.Duration) *GetListTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get list tasks params
func (o *GetListTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get list tasks params
func (o *GetListTasksParams) WithContext(ctx context.Context) *GetListTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get list tasks params
func (o *GetListTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get list tasks params
func (o *GetListTasksParams) WithHTTPClient(client *http.Client) *GetListTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get list tasks params
func (o *GetListTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetListTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
