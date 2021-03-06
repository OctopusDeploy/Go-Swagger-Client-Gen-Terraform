// Code generated by go-swagger; DO NOT EDIT.

package environments

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

// NewGetVariablesScopedToEnvironmentSpacesParams creates a new GetVariablesScopedToEnvironmentSpacesParams object
// with the default values initialized.
func NewGetVariablesScopedToEnvironmentSpacesParams() *GetVariablesScopedToEnvironmentSpacesParams {
	var ()
	return &GetVariablesScopedToEnvironmentSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVariablesScopedToEnvironmentSpacesParamsWithTimeout creates a new GetVariablesScopedToEnvironmentSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVariablesScopedToEnvironmentSpacesParamsWithTimeout(timeout time.Duration) *GetVariablesScopedToEnvironmentSpacesParams {
	var ()
	return &GetVariablesScopedToEnvironmentSpacesParams{

		timeout: timeout,
	}
}

// NewGetVariablesScopedToEnvironmentSpacesParamsWithContext creates a new GetVariablesScopedToEnvironmentSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVariablesScopedToEnvironmentSpacesParamsWithContext(ctx context.Context) *GetVariablesScopedToEnvironmentSpacesParams {
	var ()
	return &GetVariablesScopedToEnvironmentSpacesParams{

		Context: ctx,
	}
}

// NewGetVariablesScopedToEnvironmentSpacesParamsWithHTTPClient creates a new GetVariablesScopedToEnvironmentSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVariablesScopedToEnvironmentSpacesParamsWithHTTPClient(client *http.Client) *GetVariablesScopedToEnvironmentSpacesParams {
	var ()
	return &GetVariablesScopedToEnvironmentSpacesParams{
		HTTPClient: client,
	}
}

/*GetVariablesScopedToEnvironmentSpacesParams contains all the parameters to send to the API endpoint
for the get variables scoped to environment spaces operation typically these are written to a http.Request
*/
type GetVariablesScopedToEnvironmentSpacesParams struct {

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

// WithTimeout adds the timeout to the get variables scoped to environment spaces params
func (o *GetVariablesScopedToEnvironmentSpacesParams) WithTimeout(timeout time.Duration) *GetVariablesScopedToEnvironmentSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get variables scoped to environment spaces params
func (o *GetVariablesScopedToEnvironmentSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get variables scoped to environment spaces params
func (o *GetVariablesScopedToEnvironmentSpacesParams) WithContext(ctx context.Context) *GetVariablesScopedToEnvironmentSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get variables scoped to environment spaces params
func (o *GetVariablesScopedToEnvironmentSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get variables scoped to environment spaces params
func (o *GetVariablesScopedToEnvironmentSpacesParams) WithHTTPClient(client *http.Client) *GetVariablesScopedToEnvironmentSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get variables scoped to environment spaces params
func (o *GetVariablesScopedToEnvironmentSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get variables scoped to environment spaces params
func (o *GetVariablesScopedToEnvironmentSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetVariablesScopedToEnvironmentSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get variables scoped to environment spaces params
func (o *GetVariablesScopedToEnvironmentSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get variables scoped to environment spaces params
func (o *GetVariablesScopedToEnvironmentSpacesParams) WithID(id string) *GetVariablesScopedToEnvironmentSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get variables scoped to environment spaces params
func (o *GetVariablesScopedToEnvironmentSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVariablesScopedToEnvironmentSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
