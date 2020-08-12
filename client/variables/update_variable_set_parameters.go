// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewUpdateVariableSetParams creates a new UpdateVariableSetParams object
// with the default values initialized.
func NewUpdateVariableSetParams() *UpdateVariableSetParams {
	var ()
	return &UpdateVariableSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVariableSetParamsWithTimeout creates a new UpdateVariableSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVariableSetParamsWithTimeout(timeout time.Duration) *UpdateVariableSetParams {
	var ()
	return &UpdateVariableSetParams{

		timeout: timeout,
	}
}

// NewUpdateVariableSetParamsWithContext creates a new UpdateVariableSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVariableSetParamsWithContext(ctx context.Context) *UpdateVariableSetParams {
	var ()
	return &UpdateVariableSetParams{

		Context: ctx,
	}
}

// NewUpdateVariableSetParamsWithHTTPClient creates a new UpdateVariableSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateVariableSetParamsWithHTTPClient(client *http.Client) *UpdateVariableSetParams {
	var ()
	return &UpdateVariableSetParams{
		HTTPClient: client,
	}
}

/*UpdateVariableSetParams contains all the parameters to send to the API endpoint
for the update variable set operation typically these are written to a http.Request
*/
type UpdateVariableSetParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update variable set params
func (o *UpdateVariableSetParams) WithTimeout(timeout time.Duration) *UpdateVariableSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update variable set params
func (o *UpdateVariableSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update variable set params
func (o *UpdateVariableSetParams) WithContext(ctx context.Context) *UpdateVariableSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update variable set params
func (o *UpdateVariableSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update variable set params
func (o *UpdateVariableSetParams) WithHTTPClient(client *http.Client) *UpdateVariableSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update variable set params
func (o *UpdateVariableSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update variable set params
func (o *UpdateVariableSetParams) WithID(id string) *UpdateVariableSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update variable set params
func (o *UpdateVariableSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVariableSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
