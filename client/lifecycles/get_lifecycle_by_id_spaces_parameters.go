// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

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

// NewGetLifecycleByIDSpacesParams creates a new GetLifecycleByIDSpacesParams object
// with the default values initialized.
func NewGetLifecycleByIDSpacesParams() *GetLifecycleByIDSpacesParams {
	var ()
	return &GetLifecycleByIDSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLifecycleByIDSpacesParamsWithTimeout creates a new GetLifecycleByIDSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLifecycleByIDSpacesParamsWithTimeout(timeout time.Duration) *GetLifecycleByIDSpacesParams {
	var ()
	return &GetLifecycleByIDSpacesParams{

		timeout: timeout,
	}
}

// NewGetLifecycleByIDSpacesParamsWithContext creates a new GetLifecycleByIDSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLifecycleByIDSpacesParamsWithContext(ctx context.Context) *GetLifecycleByIDSpacesParams {
	var ()
	return &GetLifecycleByIDSpacesParams{

		Context: ctx,
	}
}

// NewGetLifecycleByIDSpacesParamsWithHTTPClient creates a new GetLifecycleByIDSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLifecycleByIDSpacesParamsWithHTTPClient(client *http.Client) *GetLifecycleByIDSpacesParams {
	var ()
	return &GetLifecycleByIDSpacesParams{
		HTTPClient: client,
	}
}

/*GetLifecycleByIDSpacesParams contains all the parameters to send to the API endpoint
for the get lifecycle by Id spaces operation typically these are written to a http.Request
*/
type GetLifecycleByIDSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the LifecycleResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lifecycle by Id spaces params
func (o *GetLifecycleByIDSpacesParams) WithTimeout(timeout time.Duration) *GetLifecycleByIDSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lifecycle by Id spaces params
func (o *GetLifecycleByIDSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lifecycle by Id spaces params
func (o *GetLifecycleByIDSpacesParams) WithContext(ctx context.Context) *GetLifecycleByIDSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lifecycle by Id spaces params
func (o *GetLifecycleByIDSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lifecycle by Id spaces params
func (o *GetLifecycleByIDSpacesParams) WithHTTPClient(client *http.Client) *GetLifecycleByIDSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lifecycle by Id spaces params
func (o *GetLifecycleByIDSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get lifecycle by Id spaces params
func (o *GetLifecycleByIDSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetLifecycleByIDSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get lifecycle by Id spaces params
func (o *GetLifecycleByIDSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get lifecycle by Id spaces params
func (o *GetLifecycleByIDSpacesParams) WithID(id string) *GetLifecycleByIDSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get lifecycle by Id spaces params
func (o *GetLifecycleByIDSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLifecycleByIDSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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