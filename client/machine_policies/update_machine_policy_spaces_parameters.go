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

	"models"
)

// NewUpdateMachinePolicySpacesParams creates a new UpdateMachinePolicySpacesParams object
// with the default values initialized.
func NewUpdateMachinePolicySpacesParams() *UpdateMachinePolicySpacesParams {
	var ()
	return &UpdateMachinePolicySpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMachinePolicySpacesParamsWithTimeout creates a new UpdateMachinePolicySpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMachinePolicySpacesParamsWithTimeout(timeout time.Duration) *UpdateMachinePolicySpacesParams {
	var ()
	return &UpdateMachinePolicySpacesParams{

		timeout: timeout,
	}
}

// NewUpdateMachinePolicySpacesParamsWithContext creates a new UpdateMachinePolicySpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMachinePolicySpacesParamsWithContext(ctx context.Context) *UpdateMachinePolicySpacesParams {
	var ()
	return &UpdateMachinePolicySpacesParams{

		Context: ctx,
	}
}

// NewUpdateMachinePolicySpacesParamsWithHTTPClient creates a new UpdateMachinePolicySpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMachinePolicySpacesParamsWithHTTPClient(client *http.Client) *UpdateMachinePolicySpacesParams {
	var ()
	return &UpdateMachinePolicySpacesParams{
		HTTPClient: client,
	}
}

/*UpdateMachinePolicySpacesParams contains all the parameters to send to the API endpoint
for the update machine policy spaces operation typically these are written to a http.Request
*/
type UpdateMachinePolicySpacesParams struct {

	/*MachinePolicyResource
	  The MachinePolicyResource resource to create

	*/
	MachinePolicyResource *models.MachinePolicyResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the MachinePolicyResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) WithTimeout(timeout time.Duration) *UpdateMachinePolicySpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) WithContext(ctx context.Context) *UpdateMachinePolicySpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) WithHTTPClient(client *http.Client) *UpdateMachinePolicySpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMachinePolicyResource adds the machinePolicyResource to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) WithMachinePolicyResource(machinePolicyResource *models.MachinePolicyResource) *UpdateMachinePolicySpacesParams {
	o.SetMachinePolicyResource(machinePolicyResource)
	return o
}

// SetMachinePolicyResource adds the machinePolicyResource to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) SetMachinePolicyResource(machinePolicyResource *models.MachinePolicyResource) {
	o.MachinePolicyResource = machinePolicyResource
}

// WithBaseSpaceID adds the baseSpaceID to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) WithBaseSpaceID(baseSpaceID string) *UpdateMachinePolicySpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) WithID(id string) *UpdateMachinePolicySpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update machine policy spaces params
func (o *UpdateMachinePolicySpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMachinePolicySpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MachinePolicyResource != nil {
		if err := r.SetBodyParam(o.MachinePolicyResource); err != nil {
			return err
		}
	}

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
