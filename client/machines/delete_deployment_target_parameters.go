// Code generated by go-swagger; DO NOT EDIT.

package machines

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

// NewDeleteDeploymentTargetParams creates a new DeleteDeploymentTargetParams object
// with the default values initialized.
func NewDeleteDeploymentTargetParams() *DeleteDeploymentTargetParams {
	var ()
	return &DeleteDeploymentTargetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeploymentTargetParamsWithTimeout creates a new DeleteDeploymentTargetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeploymentTargetParamsWithTimeout(timeout time.Duration) *DeleteDeploymentTargetParams {
	var ()
	return &DeleteDeploymentTargetParams{

		timeout: timeout,
	}
}

// NewDeleteDeploymentTargetParamsWithContext creates a new DeleteDeploymentTargetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeploymentTargetParamsWithContext(ctx context.Context) *DeleteDeploymentTargetParams {
	var ()
	return &DeleteDeploymentTargetParams{

		Context: ctx,
	}
}

// NewDeleteDeploymentTargetParamsWithHTTPClient creates a new DeleteDeploymentTargetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeploymentTargetParamsWithHTTPClient(client *http.Client) *DeleteDeploymentTargetParams {
	var ()
	return &DeleteDeploymentTargetParams{
		HTTPClient: client,
	}
}

/*DeleteDeploymentTargetParams contains all the parameters to send to the API endpoint
for the delete deployment target operation typically these are written to a http.Request
*/
type DeleteDeploymentTargetParams struct {

	/*ID
	  ID of the DeploymentTargetResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete deployment target params
func (o *DeleteDeploymentTargetParams) WithTimeout(timeout time.Duration) *DeleteDeploymentTargetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete deployment target params
func (o *DeleteDeploymentTargetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete deployment target params
func (o *DeleteDeploymentTargetParams) WithContext(ctx context.Context) *DeleteDeploymentTargetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete deployment target params
func (o *DeleteDeploymentTargetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete deployment target params
func (o *DeleteDeploymentTargetParams) WithHTTPClient(client *http.Client) *DeleteDeploymentTargetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete deployment target params
func (o *DeleteDeploymentTargetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete deployment target params
func (o *DeleteDeploymentTargetParams) WithID(id string) *DeleteDeploymentTargetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete deployment target params
func (o *DeleteDeploymentTargetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeploymentTargetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
