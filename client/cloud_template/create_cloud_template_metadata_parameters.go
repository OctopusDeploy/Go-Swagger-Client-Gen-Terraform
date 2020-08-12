// Code generated by go-swagger; DO NOT EDIT.

package cloud_template

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

// NewCreateCloudTemplateMetadataParams creates a new CreateCloudTemplateMetadataParams object
// with the default values initialized.
func NewCreateCloudTemplateMetadataParams() *CreateCloudTemplateMetadataParams {
	var ()
	return &CreateCloudTemplateMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCloudTemplateMetadataParamsWithTimeout creates a new CreateCloudTemplateMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCloudTemplateMetadataParamsWithTimeout(timeout time.Duration) *CreateCloudTemplateMetadataParams {
	var ()
	return &CreateCloudTemplateMetadataParams{

		timeout: timeout,
	}
}

// NewCreateCloudTemplateMetadataParamsWithContext creates a new CreateCloudTemplateMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCloudTemplateMetadataParamsWithContext(ctx context.Context) *CreateCloudTemplateMetadataParams {
	var ()
	return &CreateCloudTemplateMetadataParams{

		Context: ctx,
	}
}

// NewCreateCloudTemplateMetadataParamsWithHTTPClient creates a new CreateCloudTemplateMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCloudTemplateMetadataParamsWithHTTPClient(client *http.Client) *CreateCloudTemplateMetadataParams {
	var ()
	return &CreateCloudTemplateMetadataParams{
		HTTPClient: client,
	}
}

/*CreateCloudTemplateMetadataParams contains all the parameters to send to the API endpoint
for the create cloud template metadata operation typically these are written to a http.Request
*/
type CreateCloudTemplateMetadataParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cloud template metadata params
func (o *CreateCloudTemplateMetadataParams) WithTimeout(timeout time.Duration) *CreateCloudTemplateMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cloud template metadata params
func (o *CreateCloudTemplateMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cloud template metadata params
func (o *CreateCloudTemplateMetadataParams) WithContext(ctx context.Context) *CreateCloudTemplateMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cloud template metadata params
func (o *CreateCloudTemplateMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cloud template metadata params
func (o *CreateCloudTemplateMetadataParams) WithHTTPClient(client *http.Client) *CreateCloudTemplateMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cloud template metadata params
func (o *CreateCloudTemplateMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create cloud template metadata params
func (o *CreateCloudTemplateMetadataParams) WithID(id string) *CreateCloudTemplateMetadataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create cloud template metadata params
func (o *CreateCloudTemplateMetadataParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCloudTemplateMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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