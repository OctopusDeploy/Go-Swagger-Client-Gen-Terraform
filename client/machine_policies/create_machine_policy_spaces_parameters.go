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

// NewCreateMachinePolicySpacesParams creates a new CreateMachinePolicySpacesParams object
// with the default values initialized.
func NewCreateMachinePolicySpacesParams() *CreateMachinePolicySpacesParams {
	var ()
	return &CreateMachinePolicySpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMachinePolicySpacesParamsWithTimeout creates a new CreateMachinePolicySpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMachinePolicySpacesParamsWithTimeout(timeout time.Duration) *CreateMachinePolicySpacesParams {
	var ()
	return &CreateMachinePolicySpacesParams{

		timeout: timeout,
	}
}

// NewCreateMachinePolicySpacesParamsWithContext creates a new CreateMachinePolicySpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMachinePolicySpacesParamsWithContext(ctx context.Context) *CreateMachinePolicySpacesParams {
	var ()
	return &CreateMachinePolicySpacesParams{

		Context: ctx,
	}
}

// NewCreateMachinePolicySpacesParamsWithHTTPClient creates a new CreateMachinePolicySpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMachinePolicySpacesParamsWithHTTPClient(client *http.Client) *CreateMachinePolicySpacesParams {
	var ()
	return &CreateMachinePolicySpacesParams{
		HTTPClient: client,
	}
}

/*CreateMachinePolicySpacesParams contains all the parameters to send to the API endpoint
for the create machine policy spaces operation typically these are written to a http.Request
*/
type CreateMachinePolicySpacesParams struct {

	/*MachinePolicyResource
	  The MachinePolicyResource resource to create

	*/
	MachinePolicyResource *models.MachinePolicyResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create machine policy spaces params
func (o *CreateMachinePolicySpacesParams) WithTimeout(timeout time.Duration) *CreateMachinePolicySpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create machine policy spaces params
func (o *CreateMachinePolicySpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create machine policy spaces params
func (o *CreateMachinePolicySpacesParams) WithContext(ctx context.Context) *CreateMachinePolicySpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create machine policy spaces params
func (o *CreateMachinePolicySpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create machine policy spaces params
func (o *CreateMachinePolicySpacesParams) WithHTTPClient(client *http.Client) *CreateMachinePolicySpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create machine policy spaces params
func (o *CreateMachinePolicySpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMachinePolicyResource adds the machinePolicyResource to the create machine policy spaces params
func (o *CreateMachinePolicySpacesParams) WithMachinePolicyResource(machinePolicyResource *models.MachinePolicyResource) *CreateMachinePolicySpacesParams {
	o.SetMachinePolicyResource(machinePolicyResource)
	return o
}

// SetMachinePolicyResource adds the machinePolicyResource to the create machine policy spaces params
func (o *CreateMachinePolicySpacesParams) SetMachinePolicyResource(machinePolicyResource *models.MachinePolicyResource) {
	o.MachinePolicyResource = machinePolicyResource
}

// WithBaseSpaceID adds the baseSpaceID to the create machine policy spaces params
func (o *CreateMachinePolicySpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateMachinePolicySpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create machine policy spaces params
func (o *CreateMachinePolicySpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMachinePolicySpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
