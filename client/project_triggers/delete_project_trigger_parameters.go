// Code generated by go-swagger; DO NOT EDIT.

package project_triggers

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

// NewDeleteProjectTriggerParams creates a new DeleteProjectTriggerParams object
// with the default values initialized.
func NewDeleteProjectTriggerParams() *DeleteProjectTriggerParams {
	var ()
	return &DeleteProjectTriggerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectTriggerParamsWithTimeout creates a new DeleteProjectTriggerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProjectTriggerParamsWithTimeout(timeout time.Duration) *DeleteProjectTriggerParams {
	var ()
	return &DeleteProjectTriggerParams{

		timeout: timeout,
	}
}

// NewDeleteProjectTriggerParamsWithContext creates a new DeleteProjectTriggerParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProjectTriggerParamsWithContext(ctx context.Context) *DeleteProjectTriggerParams {
	var ()
	return &DeleteProjectTriggerParams{

		Context: ctx,
	}
}

// NewDeleteProjectTriggerParamsWithHTTPClient creates a new DeleteProjectTriggerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProjectTriggerParamsWithHTTPClient(client *http.Client) *DeleteProjectTriggerParams {
	var ()
	return &DeleteProjectTriggerParams{
		HTTPClient: client,
	}
}

/*DeleteProjectTriggerParams contains all the parameters to send to the API endpoint
for the delete project trigger operation typically these are written to a http.Request
*/
type DeleteProjectTriggerParams struct {

	/*ID
	  ID of the ProjectTriggerResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete project trigger params
func (o *DeleteProjectTriggerParams) WithTimeout(timeout time.Duration) *DeleteProjectTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project trigger params
func (o *DeleteProjectTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project trigger params
func (o *DeleteProjectTriggerParams) WithContext(ctx context.Context) *DeleteProjectTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project trigger params
func (o *DeleteProjectTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project trigger params
func (o *DeleteProjectTriggerParams) WithHTTPClient(client *http.Client) *DeleteProjectTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project trigger params
func (o *DeleteProjectTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete project trigger params
func (o *DeleteProjectTriggerParams) WithID(id string) *DeleteProjectTriggerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete project trigger params
func (o *DeleteProjectTriggerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
