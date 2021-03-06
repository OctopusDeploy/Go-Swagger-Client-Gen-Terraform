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

// NewCreateTaskRerunSpacesParams creates a new CreateTaskRerunSpacesParams object
// with the default values initialized.
func NewCreateTaskRerunSpacesParams() *CreateTaskRerunSpacesParams {
	var ()
	return &CreateTaskRerunSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTaskRerunSpacesParamsWithTimeout creates a new CreateTaskRerunSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTaskRerunSpacesParamsWithTimeout(timeout time.Duration) *CreateTaskRerunSpacesParams {
	var ()
	return &CreateTaskRerunSpacesParams{

		timeout: timeout,
	}
}

// NewCreateTaskRerunSpacesParamsWithContext creates a new CreateTaskRerunSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTaskRerunSpacesParamsWithContext(ctx context.Context) *CreateTaskRerunSpacesParams {
	var ()
	return &CreateTaskRerunSpacesParams{

		Context: ctx,
	}
}

// NewCreateTaskRerunSpacesParamsWithHTTPClient creates a new CreateTaskRerunSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTaskRerunSpacesParamsWithHTTPClient(client *http.Client) *CreateTaskRerunSpacesParams {
	var ()
	return &CreateTaskRerunSpacesParams{
		HTTPClient: client,
	}
}

/*CreateTaskRerunSpacesParams contains all the parameters to send to the API endpoint
for the create task rerun spaces operation typically these are written to a http.Request
*/
type CreateTaskRerunSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create task rerun spaces params
func (o *CreateTaskRerunSpacesParams) WithTimeout(timeout time.Duration) *CreateTaskRerunSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create task rerun spaces params
func (o *CreateTaskRerunSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create task rerun spaces params
func (o *CreateTaskRerunSpacesParams) WithContext(ctx context.Context) *CreateTaskRerunSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create task rerun spaces params
func (o *CreateTaskRerunSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create task rerun spaces params
func (o *CreateTaskRerunSpacesParams) WithHTTPClient(client *http.Client) *CreateTaskRerunSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create task rerun spaces params
func (o *CreateTaskRerunSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the create task rerun spaces params
func (o *CreateTaskRerunSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateTaskRerunSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create task rerun spaces params
func (o *CreateTaskRerunSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the create task rerun spaces params
func (o *CreateTaskRerunSpacesParams) WithID(id string) *CreateTaskRerunSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create task rerun spaces params
func (o *CreateTaskRerunSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTaskRerunSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
