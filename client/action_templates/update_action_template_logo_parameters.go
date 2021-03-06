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

// NewUpdateActionTemplateLogoParams creates a new UpdateActionTemplateLogoParams object
// with the default values initialized.
func NewUpdateActionTemplateLogoParams() *UpdateActionTemplateLogoParams {
	var ()
	return &UpdateActionTemplateLogoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateActionTemplateLogoParamsWithTimeout creates a new UpdateActionTemplateLogoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateActionTemplateLogoParamsWithTimeout(timeout time.Duration) *UpdateActionTemplateLogoParams {
	var ()
	return &UpdateActionTemplateLogoParams{

		timeout: timeout,
	}
}

// NewUpdateActionTemplateLogoParamsWithContext creates a new UpdateActionTemplateLogoParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateActionTemplateLogoParamsWithContext(ctx context.Context) *UpdateActionTemplateLogoParams {
	var ()
	return &UpdateActionTemplateLogoParams{

		Context: ctx,
	}
}

// NewUpdateActionTemplateLogoParamsWithHTTPClient creates a new UpdateActionTemplateLogoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateActionTemplateLogoParamsWithHTTPClient(client *http.Client) *UpdateActionTemplateLogoParams {
	var ()
	return &UpdateActionTemplateLogoParams{
		HTTPClient: client,
	}
}

/*UpdateActionTemplateLogoParams contains all the parameters to send to the API endpoint
for the update action template logo operation typically these are written to a http.Request
*/
type UpdateActionTemplateLogoParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update action template logo params
func (o *UpdateActionTemplateLogoParams) WithTimeout(timeout time.Duration) *UpdateActionTemplateLogoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update action template logo params
func (o *UpdateActionTemplateLogoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update action template logo params
func (o *UpdateActionTemplateLogoParams) WithContext(ctx context.Context) *UpdateActionTemplateLogoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update action template logo params
func (o *UpdateActionTemplateLogoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update action template logo params
func (o *UpdateActionTemplateLogoParams) WithHTTPClient(client *http.Client) *UpdateActionTemplateLogoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update action template logo params
func (o *UpdateActionTemplateLogoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update action template logo params
func (o *UpdateActionTemplateLogoParams) WithID(id string) *UpdateActionTemplateLogoParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update action template logo params
func (o *UpdateActionTemplateLogoParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateActionTemplateLogoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
