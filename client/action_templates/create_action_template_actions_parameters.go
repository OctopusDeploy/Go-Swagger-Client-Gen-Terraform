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

// NewCreateActionTemplateActionsParams creates a new CreateActionTemplateActionsParams object
// with the default values initialized.
func NewCreateActionTemplateActionsParams() *CreateActionTemplateActionsParams {
	var ()
	return &CreateActionTemplateActionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateActionTemplateActionsParamsWithTimeout creates a new CreateActionTemplateActionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateActionTemplateActionsParamsWithTimeout(timeout time.Duration) *CreateActionTemplateActionsParams {
	var ()
	return &CreateActionTemplateActionsParams{

		timeout: timeout,
	}
}

// NewCreateActionTemplateActionsParamsWithContext creates a new CreateActionTemplateActionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateActionTemplateActionsParamsWithContext(ctx context.Context) *CreateActionTemplateActionsParams {
	var ()
	return &CreateActionTemplateActionsParams{

		Context: ctx,
	}
}

// NewCreateActionTemplateActionsParamsWithHTTPClient creates a new CreateActionTemplateActionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateActionTemplateActionsParamsWithHTTPClient(client *http.Client) *CreateActionTemplateActionsParams {
	var ()
	return &CreateActionTemplateActionsParams{
		HTTPClient: client,
	}
}

/*CreateActionTemplateActionsParams contains all the parameters to send to the API endpoint
for the create action template actions operation typically these are written to a http.Request
*/
type CreateActionTemplateActionsParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create action template actions params
func (o *CreateActionTemplateActionsParams) WithTimeout(timeout time.Duration) *CreateActionTemplateActionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create action template actions params
func (o *CreateActionTemplateActionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create action template actions params
func (o *CreateActionTemplateActionsParams) WithContext(ctx context.Context) *CreateActionTemplateActionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create action template actions params
func (o *CreateActionTemplateActionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create action template actions params
func (o *CreateActionTemplateActionsParams) WithHTTPClient(client *http.Client) *CreateActionTemplateActionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create action template actions params
func (o *CreateActionTemplateActionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create action template actions params
func (o *CreateActionTemplateActionsParams) WithID(id string) *CreateActionTemplateActionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create action template actions params
func (o *CreateActionTemplateActionsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateActionTemplateActionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
