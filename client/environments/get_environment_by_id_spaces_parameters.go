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

// NewGetEnvironmentByIDSpacesParams creates a new GetEnvironmentByIDSpacesParams object
// with the default values initialized.
func NewGetEnvironmentByIDSpacesParams() *GetEnvironmentByIDSpacesParams {
	var ()
	return &GetEnvironmentByIDSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnvironmentByIDSpacesParamsWithTimeout creates a new GetEnvironmentByIDSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEnvironmentByIDSpacesParamsWithTimeout(timeout time.Duration) *GetEnvironmentByIDSpacesParams {
	var ()
	return &GetEnvironmentByIDSpacesParams{

		timeout: timeout,
	}
}

// NewGetEnvironmentByIDSpacesParamsWithContext creates a new GetEnvironmentByIDSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEnvironmentByIDSpacesParamsWithContext(ctx context.Context) *GetEnvironmentByIDSpacesParams {
	var ()
	return &GetEnvironmentByIDSpacesParams{

		Context: ctx,
	}
}

// NewGetEnvironmentByIDSpacesParamsWithHTTPClient creates a new GetEnvironmentByIDSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEnvironmentByIDSpacesParamsWithHTTPClient(client *http.Client) *GetEnvironmentByIDSpacesParams {
	var ()
	return &GetEnvironmentByIDSpacesParams{
		HTTPClient: client,
	}
}

/*GetEnvironmentByIDSpacesParams contains all the parameters to send to the API endpoint
for the get environment by Id spaces operation typically these are written to a http.Request
*/
type GetEnvironmentByIDSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the EnvironmentResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get environment by Id spaces params
func (o *GetEnvironmentByIDSpacesParams) WithTimeout(timeout time.Duration) *GetEnvironmentByIDSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get environment by Id spaces params
func (o *GetEnvironmentByIDSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get environment by Id spaces params
func (o *GetEnvironmentByIDSpacesParams) WithContext(ctx context.Context) *GetEnvironmentByIDSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get environment by Id spaces params
func (o *GetEnvironmentByIDSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get environment by Id spaces params
func (o *GetEnvironmentByIDSpacesParams) WithHTTPClient(client *http.Client) *GetEnvironmentByIDSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get environment by Id spaces params
func (o *GetEnvironmentByIDSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get environment by Id spaces params
func (o *GetEnvironmentByIDSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetEnvironmentByIDSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get environment by Id spaces params
func (o *GetEnvironmentByIDSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get environment by Id spaces params
func (o *GetEnvironmentByIDSpacesParams) WithID(id string) *GetEnvironmentByIDSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get environment by Id spaces params
func (o *GetEnvironmentByIDSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnvironmentByIDSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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