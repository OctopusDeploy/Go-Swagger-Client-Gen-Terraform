// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewGetTenantTagTestSpacesParams creates a new GetTenantTagTestSpacesParams object
// with the default values initialized.
func NewGetTenantTagTestSpacesParams() *GetTenantTagTestSpacesParams {
	var ()
	return &GetTenantTagTestSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantTagTestSpacesParamsWithTimeout creates a new GetTenantTagTestSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenantTagTestSpacesParamsWithTimeout(timeout time.Duration) *GetTenantTagTestSpacesParams {
	var ()
	return &GetTenantTagTestSpacesParams{

		timeout: timeout,
	}
}

// NewGetTenantTagTestSpacesParamsWithContext creates a new GetTenantTagTestSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenantTagTestSpacesParamsWithContext(ctx context.Context) *GetTenantTagTestSpacesParams {
	var ()
	return &GetTenantTagTestSpacesParams{

		Context: ctx,
	}
}

// NewGetTenantTagTestSpacesParamsWithHTTPClient creates a new GetTenantTagTestSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenantTagTestSpacesParamsWithHTTPClient(client *http.Client) *GetTenantTagTestSpacesParams {
	var ()
	return &GetTenantTagTestSpacesParams{
		HTTPClient: client,
	}
}

/*GetTenantTagTestSpacesParams contains all the parameters to send to the API endpoint
for the get tenant tag test spaces operation typically these are written to a http.Request
*/
type GetTenantTagTestSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tenant tag test spaces params
func (o *GetTenantTagTestSpacesParams) WithTimeout(timeout time.Duration) *GetTenantTagTestSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenant tag test spaces params
func (o *GetTenantTagTestSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenant tag test spaces params
func (o *GetTenantTagTestSpacesParams) WithContext(ctx context.Context) *GetTenantTagTestSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenant tag test spaces params
func (o *GetTenantTagTestSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenant tag test spaces params
func (o *GetTenantTagTestSpacesParams) WithHTTPClient(client *http.Client) *GetTenantTagTestSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenant tag test spaces params
func (o *GetTenantTagTestSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get tenant tag test spaces params
func (o *GetTenantTagTestSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetTenantTagTestSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get tenant tag test spaces params
func (o *GetTenantTagTestSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantTagTestSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
