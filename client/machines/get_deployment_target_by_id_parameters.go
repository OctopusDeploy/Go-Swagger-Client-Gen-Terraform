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

// NewGetDeploymentTargetByIDParams creates a new GetDeploymentTargetByIDParams object
// with the default values initialized.
func NewGetDeploymentTargetByIDParams() *GetDeploymentTargetByIDParams {
	var ()
	return &GetDeploymentTargetByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentTargetByIDParamsWithTimeout creates a new GetDeploymentTargetByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentTargetByIDParamsWithTimeout(timeout time.Duration) *GetDeploymentTargetByIDParams {
	var ()
	return &GetDeploymentTargetByIDParams{

		timeout: timeout,
	}
}

// NewGetDeploymentTargetByIDParamsWithContext creates a new GetDeploymentTargetByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentTargetByIDParamsWithContext(ctx context.Context) *GetDeploymentTargetByIDParams {
	var ()
	return &GetDeploymentTargetByIDParams{

		Context: ctx,
	}
}

// NewGetDeploymentTargetByIDParamsWithHTTPClient creates a new GetDeploymentTargetByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentTargetByIDParamsWithHTTPClient(client *http.Client) *GetDeploymentTargetByIDParams {
	var ()
	return &GetDeploymentTargetByIDParams{
		HTTPClient: client,
	}
}

/*GetDeploymentTargetByIDParams contains all the parameters to send to the API endpoint
for the get deployment target by Id operation typically these are written to a http.Request
*/
type GetDeploymentTargetByIDParams struct {

	/*ID
	  ID of the DeploymentTargetResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment target by Id params
func (o *GetDeploymentTargetByIDParams) WithTimeout(timeout time.Duration) *GetDeploymentTargetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment target by Id params
func (o *GetDeploymentTargetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment target by Id params
func (o *GetDeploymentTargetByIDParams) WithContext(ctx context.Context) *GetDeploymentTargetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment target by Id params
func (o *GetDeploymentTargetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment target by Id params
func (o *GetDeploymentTargetByIDParams) WithHTTPClient(client *http.Client) *GetDeploymentTargetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment target by Id params
func (o *GetDeploymentTargetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get deployment target by Id params
func (o *GetDeploymentTargetByIDParams) WithID(id string) *GetDeploymentTargetByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get deployment target by Id params
func (o *GetDeploymentTargetByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentTargetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
