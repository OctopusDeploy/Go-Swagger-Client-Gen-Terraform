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

	"models"
)

// NewCreateDeploymentTargetParams creates a new CreateDeploymentTargetParams object
// with the default values initialized.
func NewCreateDeploymentTargetParams() *CreateDeploymentTargetParams {
	var ()
	return &CreateDeploymentTargetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeploymentTargetParamsWithTimeout creates a new CreateDeploymentTargetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDeploymentTargetParamsWithTimeout(timeout time.Duration) *CreateDeploymentTargetParams {
	var ()
	return &CreateDeploymentTargetParams{

		timeout: timeout,
	}
}

// NewCreateDeploymentTargetParamsWithContext creates a new CreateDeploymentTargetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDeploymentTargetParamsWithContext(ctx context.Context) *CreateDeploymentTargetParams {
	var ()
	return &CreateDeploymentTargetParams{

		Context: ctx,
	}
}

// NewCreateDeploymentTargetParamsWithHTTPClient creates a new CreateDeploymentTargetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDeploymentTargetParamsWithHTTPClient(client *http.Client) *CreateDeploymentTargetParams {
	var ()
	return &CreateDeploymentTargetParams{
		HTTPClient: client,
	}
}

/*CreateDeploymentTargetParams contains all the parameters to send to the API endpoint
for the create deployment target operation typically these are written to a http.Request
*/
type CreateDeploymentTargetParams struct {

	/*DeploymentTargetResource
	  The DeploymentTargetResource resource to create

	*/
	DeploymentTargetResource *models.DeploymentTargetResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create deployment target params
func (o *CreateDeploymentTargetParams) WithTimeout(timeout time.Duration) *CreateDeploymentTargetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create deployment target params
func (o *CreateDeploymentTargetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create deployment target params
func (o *CreateDeploymentTargetParams) WithContext(ctx context.Context) *CreateDeploymentTargetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create deployment target params
func (o *CreateDeploymentTargetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create deployment target params
func (o *CreateDeploymentTargetParams) WithHTTPClient(client *http.Client) *CreateDeploymentTargetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create deployment target params
func (o *CreateDeploymentTargetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentTargetResource adds the deploymentTargetResource to the create deployment target params
func (o *CreateDeploymentTargetParams) WithDeploymentTargetResource(deploymentTargetResource *models.DeploymentTargetResource) *CreateDeploymentTargetParams {
	o.SetDeploymentTargetResource(deploymentTargetResource)
	return o
}

// SetDeploymentTargetResource adds the deploymentTargetResource to the create deployment target params
func (o *CreateDeploymentTargetParams) SetDeploymentTargetResource(deploymentTargetResource *models.DeploymentTargetResource) {
	o.DeploymentTargetResource = deploymentTargetResource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeploymentTargetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeploymentTargetResource != nil {
		if err := r.SetBodyParam(o.DeploymentTargetResource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}