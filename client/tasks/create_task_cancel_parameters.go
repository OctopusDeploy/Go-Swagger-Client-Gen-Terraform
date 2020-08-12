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

// NewCreateTaskCancelParams creates a new CreateTaskCancelParams object
// with the default values initialized.
func NewCreateTaskCancelParams() *CreateTaskCancelParams {
	var ()
	return &CreateTaskCancelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTaskCancelParamsWithTimeout creates a new CreateTaskCancelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTaskCancelParamsWithTimeout(timeout time.Duration) *CreateTaskCancelParams {
	var ()
	return &CreateTaskCancelParams{

		timeout: timeout,
	}
}

// NewCreateTaskCancelParamsWithContext creates a new CreateTaskCancelParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTaskCancelParamsWithContext(ctx context.Context) *CreateTaskCancelParams {
	var ()
	return &CreateTaskCancelParams{

		Context: ctx,
	}
}

// NewCreateTaskCancelParamsWithHTTPClient creates a new CreateTaskCancelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTaskCancelParamsWithHTTPClient(client *http.Client) *CreateTaskCancelParams {
	var ()
	return &CreateTaskCancelParams{
		HTTPClient: client,
	}
}

/*CreateTaskCancelParams contains all the parameters to send to the API endpoint
for the create task cancel operation typically these are written to a http.Request
*/
type CreateTaskCancelParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create task cancel params
func (o *CreateTaskCancelParams) WithTimeout(timeout time.Duration) *CreateTaskCancelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create task cancel params
func (o *CreateTaskCancelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create task cancel params
func (o *CreateTaskCancelParams) WithContext(ctx context.Context) *CreateTaskCancelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create task cancel params
func (o *CreateTaskCancelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create task cancel params
func (o *CreateTaskCancelParams) WithHTTPClient(client *http.Client) *CreateTaskCancelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create task cancel params
func (o *CreateTaskCancelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create task cancel params
func (o *CreateTaskCancelParams) WithID(id string) *CreateTaskCancelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create task cancel params
func (o *CreateTaskCancelParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTaskCancelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}