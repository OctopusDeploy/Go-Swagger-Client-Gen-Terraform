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

// NewGetActionTemplateBasedOnCommunityActionTemplateParams creates a new GetActionTemplateBasedOnCommunityActionTemplateParams object
// with the default values initialized.
func NewGetActionTemplateBasedOnCommunityActionTemplateParams() *GetActionTemplateBasedOnCommunityActionTemplateParams {
	var ()
	return &GetActionTemplateBasedOnCommunityActionTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetActionTemplateBasedOnCommunityActionTemplateParamsWithTimeout creates a new GetActionTemplateBasedOnCommunityActionTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetActionTemplateBasedOnCommunityActionTemplateParamsWithTimeout(timeout time.Duration) *GetActionTemplateBasedOnCommunityActionTemplateParams {
	var ()
	return &GetActionTemplateBasedOnCommunityActionTemplateParams{

		timeout: timeout,
	}
}

// NewGetActionTemplateBasedOnCommunityActionTemplateParamsWithContext creates a new GetActionTemplateBasedOnCommunityActionTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetActionTemplateBasedOnCommunityActionTemplateParamsWithContext(ctx context.Context) *GetActionTemplateBasedOnCommunityActionTemplateParams {
	var ()
	return &GetActionTemplateBasedOnCommunityActionTemplateParams{

		Context: ctx,
	}
}

// NewGetActionTemplateBasedOnCommunityActionTemplateParamsWithHTTPClient creates a new GetActionTemplateBasedOnCommunityActionTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetActionTemplateBasedOnCommunityActionTemplateParamsWithHTTPClient(client *http.Client) *GetActionTemplateBasedOnCommunityActionTemplateParams {
	var ()
	return &GetActionTemplateBasedOnCommunityActionTemplateParams{
		HTTPClient: client,
	}
}

/*GetActionTemplateBasedOnCommunityActionTemplateParams contains all the parameters to send to the API endpoint
for the get action template based on community action template operation typically these are written to a http.Request
*/
type GetActionTemplateBasedOnCommunityActionTemplateParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get action template based on community action template params
func (o *GetActionTemplateBasedOnCommunityActionTemplateParams) WithTimeout(timeout time.Duration) *GetActionTemplateBasedOnCommunityActionTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get action template based on community action template params
func (o *GetActionTemplateBasedOnCommunityActionTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get action template based on community action template params
func (o *GetActionTemplateBasedOnCommunityActionTemplateParams) WithContext(ctx context.Context) *GetActionTemplateBasedOnCommunityActionTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get action template based on community action template params
func (o *GetActionTemplateBasedOnCommunityActionTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get action template based on community action template params
func (o *GetActionTemplateBasedOnCommunityActionTemplateParams) WithHTTPClient(client *http.Client) *GetActionTemplateBasedOnCommunityActionTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get action template based on community action template params
func (o *GetActionTemplateBasedOnCommunityActionTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get action template based on community action template params
func (o *GetActionTemplateBasedOnCommunityActionTemplateParams) WithID(id string) *GetActionTemplateBasedOnCommunityActionTemplateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get action template based on community action template params
func (o *GetActionTemplateBasedOnCommunityActionTemplateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetActionTemplateBasedOnCommunityActionTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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