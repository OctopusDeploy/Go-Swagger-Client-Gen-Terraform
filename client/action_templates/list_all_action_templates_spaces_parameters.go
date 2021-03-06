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

// NewListAllActionTemplatesSpacesParams creates a new ListAllActionTemplatesSpacesParams object
// with the default values initialized.
func NewListAllActionTemplatesSpacesParams() *ListAllActionTemplatesSpacesParams {
	var ()
	return &ListAllActionTemplatesSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllActionTemplatesSpacesParamsWithTimeout creates a new ListAllActionTemplatesSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllActionTemplatesSpacesParamsWithTimeout(timeout time.Duration) *ListAllActionTemplatesSpacesParams {
	var ()
	return &ListAllActionTemplatesSpacesParams{

		timeout: timeout,
	}
}

// NewListAllActionTemplatesSpacesParamsWithContext creates a new ListAllActionTemplatesSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllActionTemplatesSpacesParamsWithContext(ctx context.Context) *ListAllActionTemplatesSpacesParams {
	var ()
	return &ListAllActionTemplatesSpacesParams{

		Context: ctx,
	}
}

// NewListAllActionTemplatesSpacesParamsWithHTTPClient creates a new ListAllActionTemplatesSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllActionTemplatesSpacesParamsWithHTTPClient(client *http.Client) *ListAllActionTemplatesSpacesParams {
	var ()
	return &ListAllActionTemplatesSpacesParams{
		HTTPClient: client,
	}
}

/*ListAllActionTemplatesSpacesParams contains all the parameters to send to the API endpoint
for the list all action templates spaces operation typically these are written to a http.Request
*/
type ListAllActionTemplatesSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all action templates spaces params
func (o *ListAllActionTemplatesSpacesParams) WithTimeout(timeout time.Duration) *ListAllActionTemplatesSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all action templates spaces params
func (o *ListAllActionTemplatesSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all action templates spaces params
func (o *ListAllActionTemplatesSpacesParams) WithContext(ctx context.Context) *ListAllActionTemplatesSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all action templates spaces params
func (o *ListAllActionTemplatesSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all action templates spaces params
func (o *ListAllActionTemplatesSpacesParams) WithHTTPClient(client *http.Client) *ListAllActionTemplatesSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all action templates spaces params
func (o *ListAllActionTemplatesSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the list all action templates spaces params
func (o *ListAllActionTemplatesSpacesParams) WithBaseSpaceID(baseSpaceID string) *ListAllActionTemplatesSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the list all action templates spaces params
func (o *ListAllActionTemplatesSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllActionTemplatesSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
