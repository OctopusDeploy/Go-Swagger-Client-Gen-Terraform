// Code generated by go-swagger; DO NOT EDIT.

package interruptions

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

// NewCreateSubmitInterruptionParams creates a new CreateSubmitInterruptionParams object
// with the default values initialized.
func NewCreateSubmitInterruptionParams() *CreateSubmitInterruptionParams {
	var ()
	return &CreateSubmitInterruptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSubmitInterruptionParamsWithTimeout creates a new CreateSubmitInterruptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSubmitInterruptionParamsWithTimeout(timeout time.Duration) *CreateSubmitInterruptionParams {
	var ()
	return &CreateSubmitInterruptionParams{

		timeout: timeout,
	}
}

// NewCreateSubmitInterruptionParamsWithContext creates a new CreateSubmitInterruptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSubmitInterruptionParamsWithContext(ctx context.Context) *CreateSubmitInterruptionParams {
	var ()
	return &CreateSubmitInterruptionParams{

		Context: ctx,
	}
}

// NewCreateSubmitInterruptionParamsWithHTTPClient creates a new CreateSubmitInterruptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSubmitInterruptionParamsWithHTTPClient(client *http.Client) *CreateSubmitInterruptionParams {
	var ()
	return &CreateSubmitInterruptionParams{
		HTTPClient: client,
	}
}

/*CreateSubmitInterruptionParams contains all the parameters to send to the API endpoint
for the create submit interruption operation typically these are written to a http.Request
*/
type CreateSubmitInterruptionParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create submit interruption params
func (o *CreateSubmitInterruptionParams) WithTimeout(timeout time.Duration) *CreateSubmitInterruptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create submit interruption params
func (o *CreateSubmitInterruptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create submit interruption params
func (o *CreateSubmitInterruptionParams) WithContext(ctx context.Context) *CreateSubmitInterruptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create submit interruption params
func (o *CreateSubmitInterruptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create submit interruption params
func (o *CreateSubmitInterruptionParams) WithHTTPClient(client *http.Client) *CreateSubmitInterruptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create submit interruption params
func (o *CreateSubmitInterruptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create submit interruption params
func (o *CreateSubmitInterruptionParams) WithID(id string) *CreateSubmitInterruptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create submit interruption params
func (o *CreateSubmitInterruptionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubmitInterruptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
