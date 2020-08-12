// Code generated by go-swagger; DO NOT EDIT.

package releases

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

// NewCreateDefectResolvedParams creates a new CreateDefectResolvedParams object
// with the default values initialized.
func NewCreateDefectResolvedParams() *CreateDefectResolvedParams {
	var ()
	return &CreateDefectResolvedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDefectResolvedParamsWithTimeout creates a new CreateDefectResolvedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDefectResolvedParamsWithTimeout(timeout time.Duration) *CreateDefectResolvedParams {
	var ()
	return &CreateDefectResolvedParams{

		timeout: timeout,
	}
}

// NewCreateDefectResolvedParamsWithContext creates a new CreateDefectResolvedParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDefectResolvedParamsWithContext(ctx context.Context) *CreateDefectResolvedParams {
	var ()
	return &CreateDefectResolvedParams{

		Context: ctx,
	}
}

// NewCreateDefectResolvedParamsWithHTTPClient creates a new CreateDefectResolvedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDefectResolvedParamsWithHTTPClient(client *http.Client) *CreateDefectResolvedParams {
	var ()
	return &CreateDefectResolvedParams{
		HTTPClient: client,
	}
}

/*CreateDefectResolvedParams contains all the parameters to send to the API endpoint
for the create defect resolved operation typically these are written to a http.Request
*/
type CreateDefectResolvedParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create defect resolved params
func (o *CreateDefectResolvedParams) WithTimeout(timeout time.Duration) *CreateDefectResolvedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create defect resolved params
func (o *CreateDefectResolvedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create defect resolved params
func (o *CreateDefectResolvedParams) WithContext(ctx context.Context) *CreateDefectResolvedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create defect resolved params
func (o *CreateDefectResolvedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create defect resolved params
func (o *CreateDefectResolvedParams) WithHTTPClient(client *http.Client) *CreateDefectResolvedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create defect resolved params
func (o *CreateDefectResolvedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create defect resolved params
func (o *CreateDefectResolvedParams) WithID(id string) *CreateDefectResolvedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create defect resolved params
func (o *CreateDefectResolvedParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDefectResolvedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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