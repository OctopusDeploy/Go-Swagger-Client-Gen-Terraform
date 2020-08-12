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
)

// NewDeleteWorkerPoolSpacesParams creates a new DeleteWorkerPoolSpacesParams object
// with the default values initialized.
func NewDeleteWorkerPoolSpacesParams() *DeleteWorkerPoolSpacesParams {
	var ()
	return &DeleteWorkerPoolSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkerPoolSpacesParamsWithTimeout creates a new DeleteWorkerPoolSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWorkerPoolSpacesParamsWithTimeout(timeout time.Duration) *DeleteWorkerPoolSpacesParams {
	var ()
	return &DeleteWorkerPoolSpacesParams{

		timeout: timeout,
	}
}

// NewDeleteWorkerPoolSpacesParamsWithContext creates a new DeleteWorkerPoolSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWorkerPoolSpacesParamsWithContext(ctx context.Context) *DeleteWorkerPoolSpacesParams {
	var ()
	return &DeleteWorkerPoolSpacesParams{

		Context: ctx,
	}
}

// NewDeleteWorkerPoolSpacesParamsWithHTTPClient creates a new DeleteWorkerPoolSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWorkerPoolSpacesParamsWithHTTPClient(client *http.Client) *DeleteWorkerPoolSpacesParams {
	var ()
	return &DeleteWorkerPoolSpacesParams{
		HTTPClient: client,
	}
}

/*DeleteWorkerPoolSpacesParams contains all the parameters to send to the API endpoint
for the delete worker pool spaces operation typically these are written to a http.Request
*/
type DeleteWorkerPoolSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the WorkerPoolResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete worker pool spaces params
func (o *DeleteWorkerPoolSpacesParams) WithTimeout(timeout time.Duration) *DeleteWorkerPoolSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete worker pool spaces params
func (o *DeleteWorkerPoolSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete worker pool spaces params
func (o *DeleteWorkerPoolSpacesParams) WithContext(ctx context.Context) *DeleteWorkerPoolSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete worker pool spaces params
func (o *DeleteWorkerPoolSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete worker pool spaces params
func (o *DeleteWorkerPoolSpacesParams) WithHTTPClient(client *http.Client) *DeleteWorkerPoolSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete worker pool spaces params
func (o *DeleteWorkerPoolSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the delete worker pool spaces params
func (o *DeleteWorkerPoolSpacesParams) WithBaseSpaceID(baseSpaceID string) *DeleteWorkerPoolSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the delete worker pool spaces params
func (o *DeleteWorkerPoolSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the delete worker pool spaces params
func (o *DeleteWorkerPoolSpacesParams) WithID(id string) *DeleteWorkerPoolSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete worker pool spaces params
func (o *DeleteWorkerPoolSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkerPoolSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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