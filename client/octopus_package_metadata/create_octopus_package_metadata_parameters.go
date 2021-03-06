// Code generated by go-swagger; DO NOT EDIT.

package octopus_package_metadata

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

// NewCreateOctopusPackageMetadataParams creates a new CreateOctopusPackageMetadataParams object
// with the default values initialized.
func NewCreateOctopusPackageMetadataParams() *CreateOctopusPackageMetadataParams {

	return &CreateOctopusPackageMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOctopusPackageMetadataParamsWithTimeout creates a new CreateOctopusPackageMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOctopusPackageMetadataParamsWithTimeout(timeout time.Duration) *CreateOctopusPackageMetadataParams {

	return &CreateOctopusPackageMetadataParams{

		timeout: timeout,
	}
}

// NewCreateOctopusPackageMetadataParamsWithContext creates a new CreateOctopusPackageMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOctopusPackageMetadataParamsWithContext(ctx context.Context) *CreateOctopusPackageMetadataParams {

	return &CreateOctopusPackageMetadataParams{

		Context: ctx,
	}
}

// NewCreateOctopusPackageMetadataParamsWithHTTPClient creates a new CreateOctopusPackageMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOctopusPackageMetadataParamsWithHTTPClient(client *http.Client) *CreateOctopusPackageMetadataParams {

	return &CreateOctopusPackageMetadataParams{
		HTTPClient: client,
	}
}

/*CreateOctopusPackageMetadataParams contains all the parameters to send to the API endpoint
for the create octopus package metadata operation typically these are written to a http.Request
*/
type CreateOctopusPackageMetadataParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create octopus package metadata params
func (o *CreateOctopusPackageMetadataParams) WithTimeout(timeout time.Duration) *CreateOctopusPackageMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create octopus package metadata params
func (o *CreateOctopusPackageMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create octopus package metadata params
func (o *CreateOctopusPackageMetadataParams) WithContext(ctx context.Context) *CreateOctopusPackageMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create octopus package metadata params
func (o *CreateOctopusPackageMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create octopus package metadata params
func (o *CreateOctopusPackageMetadataParams) WithHTTPClient(client *http.Client) *CreateOctopusPackageMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create octopus package metadata params
func (o *CreateOctopusPackageMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOctopusPackageMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
