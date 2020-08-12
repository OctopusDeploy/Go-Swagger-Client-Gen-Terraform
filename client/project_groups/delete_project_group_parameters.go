// Code generated by go-swagger; DO NOT EDIT.

package project_groups

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

// NewDeleteProjectGroupParams creates a new DeleteProjectGroupParams object
// with the default values initialized.
func NewDeleteProjectGroupParams() *DeleteProjectGroupParams {
	var ()
	return &DeleteProjectGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectGroupParamsWithTimeout creates a new DeleteProjectGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProjectGroupParamsWithTimeout(timeout time.Duration) *DeleteProjectGroupParams {
	var ()
	return &DeleteProjectGroupParams{

		timeout: timeout,
	}
}

// NewDeleteProjectGroupParamsWithContext creates a new DeleteProjectGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProjectGroupParamsWithContext(ctx context.Context) *DeleteProjectGroupParams {
	var ()
	return &DeleteProjectGroupParams{

		Context: ctx,
	}
}

// NewDeleteProjectGroupParamsWithHTTPClient creates a new DeleteProjectGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProjectGroupParamsWithHTTPClient(client *http.Client) *DeleteProjectGroupParams {
	var ()
	return &DeleteProjectGroupParams{
		HTTPClient: client,
	}
}

/*DeleteProjectGroupParams contains all the parameters to send to the API endpoint
for the delete project group operation typically these are written to a http.Request
*/
type DeleteProjectGroupParams struct {

	/*ID
	  ID of the ProjectGroupResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete project group params
func (o *DeleteProjectGroupParams) WithTimeout(timeout time.Duration) *DeleteProjectGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project group params
func (o *DeleteProjectGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project group params
func (o *DeleteProjectGroupParams) WithContext(ctx context.Context) *DeleteProjectGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project group params
func (o *DeleteProjectGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project group params
func (o *DeleteProjectGroupParams) WithHTTPClient(client *http.Client) *DeleteProjectGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project group params
func (o *DeleteProjectGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete project group params
func (o *DeleteProjectGroupParams) WithID(id string) *DeleteProjectGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete project group params
func (o *DeleteProjectGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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