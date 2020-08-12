// Code generated by go-swagger; DO NOT EDIT.

package machine_policies

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

// NewDeleteMachinePolicySpacesParams creates a new DeleteMachinePolicySpacesParams object
// with the default values initialized.
func NewDeleteMachinePolicySpacesParams() *DeleteMachinePolicySpacesParams {
	var ()
	return &DeleteMachinePolicySpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMachinePolicySpacesParamsWithTimeout creates a new DeleteMachinePolicySpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMachinePolicySpacesParamsWithTimeout(timeout time.Duration) *DeleteMachinePolicySpacesParams {
	var ()
	return &DeleteMachinePolicySpacesParams{

		timeout: timeout,
	}
}

// NewDeleteMachinePolicySpacesParamsWithContext creates a new DeleteMachinePolicySpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMachinePolicySpacesParamsWithContext(ctx context.Context) *DeleteMachinePolicySpacesParams {
	var ()
	return &DeleteMachinePolicySpacesParams{

		Context: ctx,
	}
}

// NewDeleteMachinePolicySpacesParamsWithHTTPClient creates a new DeleteMachinePolicySpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMachinePolicySpacesParamsWithHTTPClient(client *http.Client) *DeleteMachinePolicySpacesParams {
	var ()
	return &DeleteMachinePolicySpacesParams{
		HTTPClient: client,
	}
}

/*DeleteMachinePolicySpacesParams contains all the parameters to send to the API endpoint
for the delete machine policy spaces operation typically these are written to a http.Request
*/
type DeleteMachinePolicySpacesParams struct {

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

// WithTimeout adds the timeout to the delete machine policy spaces params
func (o *DeleteMachinePolicySpacesParams) WithTimeout(timeout time.Duration) *DeleteMachinePolicySpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete machine policy spaces params
func (o *DeleteMachinePolicySpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete machine policy spaces params
func (o *DeleteMachinePolicySpacesParams) WithContext(ctx context.Context) *DeleteMachinePolicySpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete machine policy spaces params
func (o *DeleteMachinePolicySpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete machine policy spaces params
func (o *DeleteMachinePolicySpacesParams) WithHTTPClient(client *http.Client) *DeleteMachinePolicySpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete machine policy spaces params
func (o *DeleteMachinePolicySpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the delete machine policy spaces params
func (o *DeleteMachinePolicySpacesParams) WithBaseSpaceID(baseSpaceID string) *DeleteMachinePolicySpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the delete machine policy spaces params
func (o *DeleteMachinePolicySpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the delete machine policy spaces params
func (o *DeleteMachinePolicySpacesParams) WithID(id string) *DeleteMachinePolicySpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete machine policy spaces params
func (o *DeleteMachinePolicySpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMachinePolicySpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
