// Code generated by go-swagger; DO NOT EDIT.

package action_templates

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

// NewCreateActionTemplateLogoParams creates a new CreateActionTemplateLogoParams object
// with the default values initialized.
func NewCreateActionTemplateLogoParams() *CreateActionTemplateLogoParams {
	var ()
	return &CreateActionTemplateLogoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateActionTemplateLogoParamsWithTimeout creates a new CreateActionTemplateLogoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateActionTemplateLogoParamsWithTimeout(timeout time.Duration) *CreateActionTemplateLogoParams {
	var ()
	return &CreateActionTemplateLogoParams{

		timeout: timeout,
	}
}

// NewCreateActionTemplateLogoParamsWithContext creates a new CreateActionTemplateLogoParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateActionTemplateLogoParamsWithContext(ctx context.Context) *CreateActionTemplateLogoParams {
	var ()
	return &CreateActionTemplateLogoParams{

		Context: ctx,
	}
}

// NewCreateActionTemplateLogoParamsWithHTTPClient creates a new CreateActionTemplateLogoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateActionTemplateLogoParamsWithHTTPClient(client *http.Client) *CreateActionTemplateLogoParams {
	var ()
	return &CreateActionTemplateLogoParams{
		HTTPClient: client,
	}
}

/*CreateActionTemplateLogoParams contains all the parameters to send to the API endpoint
for the create action template logo operation typically these are written to a http.Request
*/
type CreateActionTemplateLogoParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create action template logo params
func (o *CreateActionTemplateLogoParams) WithTimeout(timeout time.Duration) *CreateActionTemplateLogoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create action template logo params
func (o *CreateActionTemplateLogoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create action template logo params
func (o *CreateActionTemplateLogoParams) WithContext(ctx context.Context) *CreateActionTemplateLogoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create action template logo params
func (o *CreateActionTemplateLogoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create action template logo params
func (o *CreateActionTemplateLogoParams) WithHTTPClient(client *http.Client) *CreateActionTemplateLogoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create action template logo params
func (o *CreateActionTemplateLogoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create action template logo params
func (o *CreateActionTemplateLogoParams) WithID(id string) *CreateActionTemplateLogoParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create action template logo params
func (o *CreateActionTemplateLogoParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateActionTemplateLogoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
