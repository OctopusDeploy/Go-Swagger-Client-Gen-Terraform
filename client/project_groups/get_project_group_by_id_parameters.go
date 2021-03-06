// Code generated by go-swagger; DO NOT EDIT.

package project_groups

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

// NewGetProjectGroupByIDParams creates a new GetProjectGroupByIDParams object
// with the default values initialized.
func NewGetProjectGroupByIDParams() *GetProjectGroupByIDParams {
	var ()
	return &GetProjectGroupByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectGroupByIDParamsWithTimeout creates a new GetProjectGroupByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectGroupByIDParamsWithTimeout(timeout time.Duration) *GetProjectGroupByIDParams {
	var ()
	return &GetProjectGroupByIDParams{

		timeout: timeout,
	}
}

// NewGetProjectGroupByIDParamsWithContext creates a new GetProjectGroupByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectGroupByIDParamsWithContext(ctx context.Context) *GetProjectGroupByIDParams {
	var ()
	return &GetProjectGroupByIDParams{

		Context: ctx,
	}
}

// NewGetProjectGroupByIDParamsWithHTTPClient creates a new GetProjectGroupByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProjectGroupByIDParamsWithHTTPClient(client *http.Client) *GetProjectGroupByIDParams {
	var ()
	return &GetProjectGroupByIDParams{
		HTTPClient: client,
	}
}

/*GetProjectGroupByIDParams contains all the parameters to send to the API endpoint
for the get project group by Id operation typically these are written to a http.Request
*/
type GetProjectGroupByIDParams struct {

	/*ID
	  ID of the ProjectGroupResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get project group by Id params
func (o *GetProjectGroupByIDParams) WithTimeout(timeout time.Duration) *GetProjectGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project group by Id params
func (o *GetProjectGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project group by Id params
func (o *GetProjectGroupByIDParams) WithContext(ctx context.Context) *GetProjectGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project group by Id params
func (o *GetProjectGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project group by Id params
func (o *GetProjectGroupByIDParams) WithHTTPClient(client *http.Client) *GetProjectGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project group by Id params
func (o *GetProjectGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get project group by Id params
func (o *GetProjectGroupByIDParams) WithID(id string) *GetProjectGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get project group by Id params
func (o *GetProjectGroupByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
