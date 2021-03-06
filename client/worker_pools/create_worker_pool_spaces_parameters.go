// Code generated by go-swagger; DO NOT EDIT.

package worker_pools

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

// NewCreateWorkerPoolSpacesParams creates a new CreateWorkerPoolSpacesParams object
// with the default values initialized.
func NewCreateWorkerPoolSpacesParams() *CreateWorkerPoolSpacesParams {
	var ()
	return &CreateWorkerPoolSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWorkerPoolSpacesParamsWithTimeout creates a new CreateWorkerPoolSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateWorkerPoolSpacesParamsWithTimeout(timeout time.Duration) *CreateWorkerPoolSpacesParams {
	var ()
	return &CreateWorkerPoolSpacesParams{

		timeout: timeout,
	}
}

// NewCreateWorkerPoolSpacesParamsWithContext creates a new CreateWorkerPoolSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateWorkerPoolSpacesParamsWithContext(ctx context.Context) *CreateWorkerPoolSpacesParams {
	var ()
	return &CreateWorkerPoolSpacesParams{

		Context: ctx,
	}
}

// NewCreateWorkerPoolSpacesParamsWithHTTPClient creates a new CreateWorkerPoolSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateWorkerPoolSpacesParamsWithHTTPClient(client *http.Client) *CreateWorkerPoolSpacesParams {
	var ()
	return &CreateWorkerPoolSpacesParams{
		HTTPClient: client,
	}
}

/*CreateWorkerPoolSpacesParams contains all the parameters to send to the API endpoint
for the create worker pool spaces operation typically these are written to a http.Request
*/
type CreateWorkerPoolSpacesParams struct {

	/*WorkerPoolResource
	  The WorkerPoolResource resource to create

	*/
	WorkerPoolResource *models.WorkerPoolResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create worker pool spaces params
func (o *CreateWorkerPoolSpacesParams) WithTimeout(timeout time.Duration) *CreateWorkerPoolSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create worker pool spaces params
func (o *CreateWorkerPoolSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create worker pool spaces params
func (o *CreateWorkerPoolSpacesParams) WithContext(ctx context.Context) *CreateWorkerPoolSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create worker pool spaces params
func (o *CreateWorkerPoolSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create worker pool spaces params
func (o *CreateWorkerPoolSpacesParams) WithHTTPClient(client *http.Client) *CreateWorkerPoolSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create worker pool spaces params
func (o *CreateWorkerPoolSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkerPoolResource adds the workerPoolResource to the create worker pool spaces params
func (o *CreateWorkerPoolSpacesParams) WithWorkerPoolResource(workerPoolResource *models.WorkerPoolResource) *CreateWorkerPoolSpacesParams {
	o.SetWorkerPoolResource(workerPoolResource)
	return o
}

// SetWorkerPoolResource adds the workerPoolResource to the create worker pool spaces params
func (o *CreateWorkerPoolSpacesParams) SetWorkerPoolResource(workerPoolResource *models.WorkerPoolResource) {
	o.WorkerPoolResource = workerPoolResource
}

// WithBaseSpaceID adds the baseSpaceID to the create worker pool spaces params
func (o *CreateWorkerPoolSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateWorkerPoolSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create worker pool spaces params
func (o *CreateWorkerPoolSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWorkerPoolSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.WorkerPoolResource != nil {
		if err := r.SetBodyParam(o.WorkerPoolResource); err != nil {
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
