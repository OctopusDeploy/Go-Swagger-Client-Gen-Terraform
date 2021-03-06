// Code generated by go-swagger; DO NOT EDIT.

package community_action_templates

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

// NewCreateCommunityActionTemplateInstallationParams creates a new CreateCommunityActionTemplateInstallationParams object
// with the default values initialized.
func NewCreateCommunityActionTemplateInstallationParams() *CreateCommunityActionTemplateInstallationParams {
	var ()
	return &CreateCommunityActionTemplateInstallationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCommunityActionTemplateInstallationParamsWithTimeout creates a new CreateCommunityActionTemplateInstallationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCommunityActionTemplateInstallationParamsWithTimeout(timeout time.Duration) *CreateCommunityActionTemplateInstallationParams {
	var ()
	return &CreateCommunityActionTemplateInstallationParams{

		timeout: timeout,
	}
}

// NewCreateCommunityActionTemplateInstallationParamsWithContext creates a new CreateCommunityActionTemplateInstallationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCommunityActionTemplateInstallationParamsWithContext(ctx context.Context) *CreateCommunityActionTemplateInstallationParams {
	var ()
	return &CreateCommunityActionTemplateInstallationParams{

		Context: ctx,
	}
}

// NewCreateCommunityActionTemplateInstallationParamsWithHTTPClient creates a new CreateCommunityActionTemplateInstallationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCommunityActionTemplateInstallationParamsWithHTTPClient(client *http.Client) *CreateCommunityActionTemplateInstallationParams {
	var ()
	return &CreateCommunityActionTemplateInstallationParams{
		HTTPClient: client,
	}
}

/*CreateCommunityActionTemplateInstallationParams contains all the parameters to send to the API endpoint
for the create community action template installation operation typically these are written to a http.Request
*/
type CreateCommunityActionTemplateInstallationParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create community action template installation params
func (o *CreateCommunityActionTemplateInstallationParams) WithTimeout(timeout time.Duration) *CreateCommunityActionTemplateInstallationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create community action template installation params
func (o *CreateCommunityActionTemplateInstallationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create community action template installation params
func (o *CreateCommunityActionTemplateInstallationParams) WithContext(ctx context.Context) *CreateCommunityActionTemplateInstallationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create community action template installation params
func (o *CreateCommunityActionTemplateInstallationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create community action template installation params
func (o *CreateCommunityActionTemplateInstallationParams) WithHTTPClient(client *http.Client) *CreateCommunityActionTemplateInstallationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create community action template installation params
func (o *CreateCommunityActionTemplateInstallationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create community action template installation params
func (o *CreateCommunityActionTemplateInstallationParams) WithID(id string) *CreateCommunityActionTemplateInstallationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create community action template installation params
func (o *CreateCommunityActionTemplateInstallationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCommunityActionTemplateInstallationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
