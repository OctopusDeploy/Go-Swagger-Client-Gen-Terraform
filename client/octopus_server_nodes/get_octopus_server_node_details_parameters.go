// Code generated by go-swagger; DO NOT EDIT.

package octopus_server_nodes

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

// NewGetOctopusServerNodeDetailsParams creates a new GetOctopusServerNodeDetailsParams object
// with the default values initialized.
func NewGetOctopusServerNodeDetailsParams() *GetOctopusServerNodeDetailsParams {
	var ()
	return &GetOctopusServerNodeDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOctopusServerNodeDetailsParamsWithTimeout creates a new GetOctopusServerNodeDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOctopusServerNodeDetailsParamsWithTimeout(timeout time.Duration) *GetOctopusServerNodeDetailsParams {
	var ()
	return &GetOctopusServerNodeDetailsParams{

		timeout: timeout,
	}
}

// NewGetOctopusServerNodeDetailsParamsWithContext creates a new GetOctopusServerNodeDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOctopusServerNodeDetailsParamsWithContext(ctx context.Context) *GetOctopusServerNodeDetailsParams {
	var ()
	return &GetOctopusServerNodeDetailsParams{

		Context: ctx,
	}
}

// NewGetOctopusServerNodeDetailsParamsWithHTTPClient creates a new GetOctopusServerNodeDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOctopusServerNodeDetailsParamsWithHTTPClient(client *http.Client) *GetOctopusServerNodeDetailsParams {
	var ()
	return &GetOctopusServerNodeDetailsParams{
		HTTPClient: client,
	}
}

/*GetOctopusServerNodeDetailsParams contains all the parameters to send to the API endpoint
for the get octopus server node details operation typically these are written to a http.Request
*/
type GetOctopusServerNodeDetailsParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get octopus server node details params
func (o *GetOctopusServerNodeDetailsParams) WithTimeout(timeout time.Duration) *GetOctopusServerNodeDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get octopus server node details params
func (o *GetOctopusServerNodeDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get octopus server node details params
func (o *GetOctopusServerNodeDetailsParams) WithContext(ctx context.Context) *GetOctopusServerNodeDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get octopus server node details params
func (o *GetOctopusServerNodeDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get octopus server node details params
func (o *GetOctopusServerNodeDetailsParams) WithHTTPClient(client *http.Client) *GetOctopusServerNodeDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get octopus server node details params
func (o *GetOctopusServerNodeDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get octopus server node details params
func (o *GetOctopusServerNodeDetailsParams) WithID(id string) *GetOctopusServerNodeDetailsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get octopus server node details params
func (o *GetOctopusServerNodeDetailsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetOctopusServerNodeDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
