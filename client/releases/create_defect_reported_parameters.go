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

// NewCreateDefectReportedParams creates a new CreateDefectReportedParams object
// with the default values initialized.
func NewCreateDefectReportedParams() *CreateDefectReportedParams {
	var ()
	return &CreateDefectReportedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDefectReportedParamsWithTimeout creates a new CreateDefectReportedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDefectReportedParamsWithTimeout(timeout time.Duration) *CreateDefectReportedParams {
	var ()
	return &CreateDefectReportedParams{

		timeout: timeout,
	}
}

// NewCreateDefectReportedParamsWithContext creates a new CreateDefectReportedParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDefectReportedParamsWithContext(ctx context.Context) *CreateDefectReportedParams {
	var ()
	return &CreateDefectReportedParams{

		Context: ctx,
	}
}

// NewCreateDefectReportedParamsWithHTTPClient creates a new CreateDefectReportedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDefectReportedParamsWithHTTPClient(client *http.Client) *CreateDefectReportedParams {
	var ()
	return &CreateDefectReportedParams{
		HTTPClient: client,
	}
}

/*CreateDefectReportedParams contains all the parameters to send to the API endpoint
for the create defect reported operation typically these are written to a http.Request
*/
type CreateDefectReportedParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create defect reported params
func (o *CreateDefectReportedParams) WithTimeout(timeout time.Duration) *CreateDefectReportedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create defect reported params
func (o *CreateDefectReportedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create defect reported params
func (o *CreateDefectReportedParams) WithContext(ctx context.Context) *CreateDefectReportedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create defect reported params
func (o *CreateDefectReportedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create defect reported params
func (o *CreateDefectReportedParams) WithHTTPClient(client *http.Client) *CreateDefectReportedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create defect reported params
func (o *CreateDefectReportedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create defect reported params
func (o *CreateDefectReportedParams) WithID(id string) *CreateDefectReportedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create defect reported params
func (o *CreateDefectReportedParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDefectReportedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
