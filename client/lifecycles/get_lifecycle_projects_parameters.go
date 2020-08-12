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

// NewGetLifecycleProjectsParams creates a new GetLifecycleProjectsParams object
// with the default values initialized.
func NewGetLifecycleProjectsParams() *GetLifecycleProjectsParams {
	var ()
	return &GetLifecycleProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLifecycleProjectsParamsWithTimeout creates a new GetLifecycleProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLifecycleProjectsParamsWithTimeout(timeout time.Duration) *GetLifecycleProjectsParams {
	var ()
	return &GetLifecycleProjectsParams{

		timeout: timeout,
	}
}

// NewGetLifecycleProjectsParamsWithContext creates a new GetLifecycleProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLifecycleProjectsParamsWithContext(ctx context.Context) *GetLifecycleProjectsParams {
	var ()
	return &GetLifecycleProjectsParams{

		Context: ctx,
	}
}

// NewGetLifecycleProjectsParamsWithHTTPClient creates a new GetLifecycleProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLifecycleProjectsParamsWithHTTPClient(client *http.Client) *GetLifecycleProjectsParams {
	var ()
	return &GetLifecycleProjectsParams{
		HTTPClient: client,
	}
}

/*GetLifecycleProjectsParams contains all the parameters to send to the API endpoint
for the get lifecycle projects operation typically these are written to a http.Request
*/
type GetLifecycleProjectsParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lifecycle projects params
func (o *GetLifecycleProjectsParams) WithTimeout(timeout time.Duration) *GetLifecycleProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lifecycle projects params
func (o *GetLifecycleProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lifecycle projects params
func (o *GetLifecycleProjectsParams) WithContext(ctx context.Context) *GetLifecycleProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lifecycle projects params
func (o *GetLifecycleProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lifecycle projects params
func (o *GetLifecycleProjectsParams) WithHTTPClient(client *http.Client) *GetLifecycleProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lifecycle projects params
func (o *GetLifecycleProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get lifecycle projects params
func (o *GetLifecycleProjectsParams) WithID(id string) *GetLifecycleProjectsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get lifecycle projects params
func (o *GetLifecycleProjectsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLifecycleProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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