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

	"models"
)

// NewCreateTaskSpacesParams creates a new CreateTaskSpacesParams object
// with the default values initialized.
func NewCreateTaskSpacesParams() *CreateTaskSpacesParams {
	var ()
	return &CreateTaskSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTaskSpacesParamsWithTimeout creates a new CreateTaskSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTaskSpacesParamsWithTimeout(timeout time.Duration) *CreateTaskSpacesParams {
	var ()
	return &CreateTaskSpacesParams{

		timeout: timeout,
	}
}

// NewCreateTaskSpacesParamsWithContext creates a new CreateTaskSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTaskSpacesParamsWithContext(ctx context.Context) *CreateTaskSpacesParams {
	var ()
	return &CreateTaskSpacesParams{

		Context: ctx,
	}
}

// NewCreateTaskSpacesParamsWithHTTPClient creates a new CreateTaskSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTaskSpacesParamsWithHTTPClient(client *http.Client) *CreateTaskSpacesParams {
	var ()
	return &CreateTaskSpacesParams{
		HTTPClient: client,
	}
}

/*CreateTaskSpacesParams contains all the parameters to send to the API endpoint
for the create task spaces operation typically these are written to a http.Request
*/
type CreateTaskSpacesParams struct {

	/*TaskResource
	  The TaskResource resource to create

	*/
	TaskResource *models.TaskResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create task spaces params
func (o *CreateTaskSpacesParams) WithTimeout(timeout time.Duration) *CreateTaskSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create task spaces params
func (o *CreateTaskSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create task spaces params
func (o *CreateTaskSpacesParams) WithContext(ctx context.Context) *CreateTaskSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create task spaces params
func (o *CreateTaskSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create task spaces params
func (o *CreateTaskSpacesParams) WithHTTPClient(client *http.Client) *CreateTaskSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create task spaces params
func (o *CreateTaskSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskResource adds the taskResource to the create task spaces params
func (o *CreateTaskSpacesParams) WithTaskResource(taskResource *models.TaskResource) *CreateTaskSpacesParams {
	o.SetTaskResource(taskResource)
	return o
}

// SetTaskResource adds the taskResource to the create task spaces params
func (o *CreateTaskSpacesParams) SetTaskResource(taskResource *models.TaskResource) {
	o.TaskResource = taskResource
}

// WithBaseSpaceID adds the baseSpaceID to the create task spaces params
func (o *CreateTaskSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateTaskSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create task spaces params
func (o *CreateTaskSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTaskSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TaskResource != nil {
		if err := r.SetBodyParam(o.TaskResource); err != nil {
			return err
		}
	}

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}