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

// NewGetMachinePolicyByIDParams creates a new GetMachinePolicyByIDParams object
// with the default values initialized.
func NewGetMachinePolicyByIDParams() *GetMachinePolicyByIDParams {
	var ()
	return &GetMachinePolicyByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMachinePolicyByIDParamsWithTimeout creates a new GetMachinePolicyByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMachinePolicyByIDParamsWithTimeout(timeout time.Duration) *GetMachinePolicyByIDParams {
	var ()
	return &GetMachinePolicyByIDParams{

		timeout: timeout,
	}
}

// NewGetMachinePolicyByIDParamsWithContext creates a new GetMachinePolicyByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMachinePolicyByIDParamsWithContext(ctx context.Context) *GetMachinePolicyByIDParams {
	var ()
	return &GetMachinePolicyByIDParams{

		Context: ctx,
	}
}

// NewGetMachinePolicyByIDParamsWithHTTPClient creates a new GetMachinePolicyByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMachinePolicyByIDParamsWithHTTPClient(client *http.Client) *GetMachinePolicyByIDParams {
	var ()
	return &GetMachinePolicyByIDParams{
		HTTPClient: client,
	}
}

/*GetMachinePolicyByIDParams contains all the parameters to send to the API endpoint
for the get machine policy by Id operation typically these are written to a http.Request
*/
type GetMachinePolicyByIDParams struct {

	/*ID
	  ID of the MachinePolicyResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get machine policy by Id params
func (o *GetMachinePolicyByIDParams) WithTimeout(timeout time.Duration) *GetMachinePolicyByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get machine policy by Id params
func (o *GetMachinePolicyByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get machine policy by Id params
func (o *GetMachinePolicyByIDParams) WithContext(ctx context.Context) *GetMachinePolicyByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get machine policy by Id params
func (o *GetMachinePolicyByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get machine policy by Id params
func (o *GetMachinePolicyByIDParams) WithHTTPClient(client *http.Client) *GetMachinePolicyByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get machine policy by Id params
func (o *GetMachinePolicyByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get machine policy by Id params
func (o *GetMachinePolicyByIDParams) WithID(id string) *GetMachinePolicyByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get machine policy by Id params
func (o *GetMachinePolicyByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMachinePolicyByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
