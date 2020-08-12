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

// NewGetActionTemplateVersionParams creates a new GetActionTemplateVersionParams object
// with the default values initialized.
func NewGetActionTemplateVersionParams() *GetActionTemplateVersionParams {
	var ()
	return &GetActionTemplateVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetActionTemplateVersionParamsWithTimeout creates a new GetActionTemplateVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetActionTemplateVersionParamsWithTimeout(timeout time.Duration) *GetActionTemplateVersionParams {
	var ()
	return &GetActionTemplateVersionParams{

		timeout: timeout,
	}
}

// NewGetActionTemplateVersionParamsWithContext creates a new GetActionTemplateVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetActionTemplateVersionParamsWithContext(ctx context.Context) *GetActionTemplateVersionParams {
	var ()
	return &GetActionTemplateVersionParams{

		Context: ctx,
	}
}

// NewGetActionTemplateVersionParamsWithHTTPClient creates a new GetActionTemplateVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetActionTemplateVersionParamsWithHTTPClient(client *http.Client) *GetActionTemplateVersionParams {
	var ()
	return &GetActionTemplateVersionParams{
		HTTPClient: client,
	}
}

/*GetActionTemplateVersionParams contains all the parameters to send to the API endpoint
for the get action template version operation typically these are written to a http.Request
*/
type GetActionTemplateVersionParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get action template version params
func (o *GetActionTemplateVersionParams) WithTimeout(timeout time.Duration) *GetActionTemplateVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get action template version params
func (o *GetActionTemplateVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get action template version params
func (o *GetActionTemplateVersionParams) WithContext(ctx context.Context) *GetActionTemplateVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get action template version params
func (o *GetActionTemplateVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get action template version params
func (o *GetActionTemplateVersionParams) WithHTTPClient(client *http.Client) *GetActionTemplateVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get action template version params
func (o *GetActionTemplateVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get action template version params
func (o *GetActionTemplateVersionParams) WithID(id string) *GetActionTemplateVersionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get action template version params
func (o *GetActionTemplateVersionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetActionTemplateVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
