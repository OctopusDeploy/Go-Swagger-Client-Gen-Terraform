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

// NewGetCommunityActionTemplateLogoParams creates a new GetCommunityActionTemplateLogoParams object
// with the default values initialized.
func NewGetCommunityActionTemplateLogoParams() *GetCommunityActionTemplateLogoParams {
	var ()
	return &GetCommunityActionTemplateLogoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCommunityActionTemplateLogoParamsWithTimeout creates a new GetCommunityActionTemplateLogoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCommunityActionTemplateLogoParamsWithTimeout(timeout time.Duration) *GetCommunityActionTemplateLogoParams {
	var ()
	return &GetCommunityActionTemplateLogoParams{

		timeout: timeout,
	}
}

// NewGetCommunityActionTemplateLogoParamsWithContext creates a new GetCommunityActionTemplateLogoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCommunityActionTemplateLogoParamsWithContext(ctx context.Context) *GetCommunityActionTemplateLogoParams {
	var ()
	return &GetCommunityActionTemplateLogoParams{

		Context: ctx,
	}
}

// NewGetCommunityActionTemplateLogoParamsWithHTTPClient creates a new GetCommunityActionTemplateLogoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCommunityActionTemplateLogoParamsWithHTTPClient(client *http.Client) *GetCommunityActionTemplateLogoParams {
	var ()
	return &GetCommunityActionTemplateLogoParams{
		HTTPClient: client,
	}
}

/*GetCommunityActionTemplateLogoParams contains all the parameters to send to the API endpoint
for the get community action template logo operation typically these are written to a http.Request
*/
type GetCommunityActionTemplateLogoParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get community action template logo params
func (o *GetCommunityActionTemplateLogoParams) WithTimeout(timeout time.Duration) *GetCommunityActionTemplateLogoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get community action template logo params
func (o *GetCommunityActionTemplateLogoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get community action template logo params
func (o *GetCommunityActionTemplateLogoParams) WithContext(ctx context.Context) *GetCommunityActionTemplateLogoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get community action template logo params
func (o *GetCommunityActionTemplateLogoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get community action template logo params
func (o *GetCommunityActionTemplateLogoParams) WithHTTPClient(client *http.Client) *GetCommunityActionTemplateLogoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get community action template logo params
func (o *GetCommunityActionTemplateLogoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get community action template logo params
func (o *GetCommunityActionTemplateLogoParams) WithID(id string) *GetCommunityActionTemplateLogoParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get community action template logo params
func (o *GetCommunityActionTemplateLogoParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCommunityActionTemplateLogoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
