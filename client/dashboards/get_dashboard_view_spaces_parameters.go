// Code generated by go-swagger; DO NOT EDIT.

package dashboards

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

// NewGetDashboardViewSpacesParams creates a new GetDashboardViewSpacesParams object
// with the default values initialized.
func NewGetDashboardViewSpacesParams() *GetDashboardViewSpacesParams {
	var ()
	return &GetDashboardViewSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardViewSpacesParamsWithTimeout creates a new GetDashboardViewSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDashboardViewSpacesParamsWithTimeout(timeout time.Duration) *GetDashboardViewSpacesParams {
	var ()
	return &GetDashboardViewSpacesParams{

		timeout: timeout,
	}
}

// NewGetDashboardViewSpacesParamsWithContext creates a new GetDashboardViewSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDashboardViewSpacesParamsWithContext(ctx context.Context) *GetDashboardViewSpacesParams {
	var ()
	return &GetDashboardViewSpacesParams{

		Context: ctx,
	}
}

// NewGetDashboardViewSpacesParamsWithHTTPClient creates a new GetDashboardViewSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDashboardViewSpacesParamsWithHTTPClient(client *http.Client) *GetDashboardViewSpacesParams {
	var ()
	return &GetDashboardViewSpacesParams{
		HTTPClient: client,
	}
}

/*GetDashboardViewSpacesParams contains all the parameters to send to the API endpoint
for the get dashboard view spaces operation typically these are written to a http.Request
*/
type GetDashboardViewSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get dashboard view spaces params
func (o *GetDashboardViewSpacesParams) WithTimeout(timeout time.Duration) *GetDashboardViewSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard view spaces params
func (o *GetDashboardViewSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard view spaces params
func (o *GetDashboardViewSpacesParams) WithContext(ctx context.Context) *GetDashboardViewSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard view spaces params
func (o *GetDashboardViewSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard view spaces params
func (o *GetDashboardViewSpacesParams) WithHTTPClient(client *http.Client) *GetDashboardViewSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard view spaces params
func (o *GetDashboardViewSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get dashboard view spaces params
func (o *GetDashboardViewSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetDashboardViewSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get dashboard view spaces params
func (o *GetDashboardViewSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardViewSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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