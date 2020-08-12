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

// NewGetActionTemplateUsageParams creates a new GetActionTemplateUsageParams object
// with the default values initialized.
func NewGetActionTemplateUsageParams() *GetActionTemplateUsageParams {
	var ()
	return &GetActionTemplateUsageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetActionTemplateUsageParamsWithTimeout creates a new GetActionTemplateUsageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetActionTemplateUsageParamsWithTimeout(timeout time.Duration) *GetActionTemplateUsageParams {
	var ()
	return &GetActionTemplateUsageParams{

		timeout: timeout,
	}
}

// NewGetActionTemplateUsageParamsWithContext creates a new GetActionTemplateUsageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetActionTemplateUsageParamsWithContext(ctx context.Context) *GetActionTemplateUsageParams {
	var ()
	return &GetActionTemplateUsageParams{

		Context: ctx,
	}
}

// NewGetActionTemplateUsageParamsWithHTTPClient creates a new GetActionTemplateUsageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetActionTemplateUsageParamsWithHTTPClient(client *http.Client) *GetActionTemplateUsageParams {
	var ()
	return &GetActionTemplateUsageParams{
		HTTPClient: client,
	}
}

/*GetActionTemplateUsageParams contains all the parameters to send to the API endpoint
for the get action template usage operation typically these are written to a http.Request
*/
type GetActionTemplateUsageParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get action template usage params
func (o *GetActionTemplateUsageParams) WithTimeout(timeout time.Duration) *GetActionTemplateUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get action template usage params
func (o *GetActionTemplateUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get action template usage params
func (o *GetActionTemplateUsageParams) WithContext(ctx context.Context) *GetActionTemplateUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get action template usage params
func (o *GetActionTemplateUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get action template usage params
func (o *GetActionTemplateUsageParams) WithHTTPClient(client *http.Client) *GetActionTemplateUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get action template usage params
func (o *GetActionTemplateUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get action template usage params
func (o *GetActionTemplateUsageParams) WithID(id string) *GetActionTemplateUsageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get action template usage params
func (o *GetActionTemplateUsageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetActionTemplateUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
