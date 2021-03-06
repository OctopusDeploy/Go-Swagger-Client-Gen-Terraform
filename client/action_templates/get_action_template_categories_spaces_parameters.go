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

// NewGetActionTemplateCategoriesSpacesParams creates a new GetActionTemplateCategoriesSpacesParams object
// with the default values initialized.
func NewGetActionTemplateCategoriesSpacesParams() *GetActionTemplateCategoriesSpacesParams {
	var ()
	return &GetActionTemplateCategoriesSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetActionTemplateCategoriesSpacesParamsWithTimeout creates a new GetActionTemplateCategoriesSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetActionTemplateCategoriesSpacesParamsWithTimeout(timeout time.Duration) *GetActionTemplateCategoriesSpacesParams {
	var ()
	return &GetActionTemplateCategoriesSpacesParams{

		timeout: timeout,
	}
}

// NewGetActionTemplateCategoriesSpacesParamsWithContext creates a new GetActionTemplateCategoriesSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetActionTemplateCategoriesSpacesParamsWithContext(ctx context.Context) *GetActionTemplateCategoriesSpacesParams {
	var ()
	return &GetActionTemplateCategoriesSpacesParams{

		Context: ctx,
	}
}

// NewGetActionTemplateCategoriesSpacesParamsWithHTTPClient creates a new GetActionTemplateCategoriesSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetActionTemplateCategoriesSpacesParamsWithHTTPClient(client *http.Client) *GetActionTemplateCategoriesSpacesParams {
	var ()
	return &GetActionTemplateCategoriesSpacesParams{
		HTTPClient: client,
	}
}

/*GetActionTemplateCategoriesSpacesParams contains all the parameters to send to the API endpoint
for the get action template categories spaces operation typically these are written to a http.Request
*/
type GetActionTemplateCategoriesSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get action template categories spaces params
func (o *GetActionTemplateCategoriesSpacesParams) WithTimeout(timeout time.Duration) *GetActionTemplateCategoriesSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get action template categories spaces params
func (o *GetActionTemplateCategoriesSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get action template categories spaces params
func (o *GetActionTemplateCategoriesSpacesParams) WithContext(ctx context.Context) *GetActionTemplateCategoriesSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get action template categories spaces params
func (o *GetActionTemplateCategoriesSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get action template categories spaces params
func (o *GetActionTemplateCategoriesSpacesParams) WithHTTPClient(client *http.Client) *GetActionTemplateCategoriesSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get action template categories spaces params
func (o *GetActionTemplateCategoriesSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get action template categories spaces params
func (o *GetActionTemplateCategoriesSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetActionTemplateCategoriesSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get action template categories spaces params
func (o *GetActionTemplateCategoriesSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetActionTemplateCategoriesSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
