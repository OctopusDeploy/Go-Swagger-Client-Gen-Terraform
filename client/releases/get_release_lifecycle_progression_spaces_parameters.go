// Code generated by go-swagger; DO NOT EDIT.

package releases

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

// NewGetReleaseLifecycleProgressionSpacesParams creates a new GetReleaseLifecycleProgressionSpacesParams object
// with the default values initialized.
func NewGetReleaseLifecycleProgressionSpacesParams() *GetReleaseLifecycleProgressionSpacesParams {
	var ()
	return &GetReleaseLifecycleProgressionSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReleaseLifecycleProgressionSpacesParamsWithTimeout creates a new GetReleaseLifecycleProgressionSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReleaseLifecycleProgressionSpacesParamsWithTimeout(timeout time.Duration) *GetReleaseLifecycleProgressionSpacesParams {
	var ()
	return &GetReleaseLifecycleProgressionSpacesParams{

		timeout: timeout,
	}
}

// NewGetReleaseLifecycleProgressionSpacesParamsWithContext creates a new GetReleaseLifecycleProgressionSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReleaseLifecycleProgressionSpacesParamsWithContext(ctx context.Context) *GetReleaseLifecycleProgressionSpacesParams {
	var ()
	return &GetReleaseLifecycleProgressionSpacesParams{

		Context: ctx,
	}
}

// NewGetReleaseLifecycleProgressionSpacesParamsWithHTTPClient creates a new GetReleaseLifecycleProgressionSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReleaseLifecycleProgressionSpacesParamsWithHTTPClient(client *http.Client) *GetReleaseLifecycleProgressionSpacesParams {
	var ()
	return &GetReleaseLifecycleProgressionSpacesParams{
		HTTPClient: client,
	}
}

/*GetReleaseLifecycleProgressionSpacesParams contains all the parameters to send to the API endpoint
for the get release lifecycle progression spaces operation typically these are written to a http.Request
*/
type GetReleaseLifecycleProgressionSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get release lifecycle progression spaces params
func (o *GetReleaseLifecycleProgressionSpacesParams) WithTimeout(timeout time.Duration) *GetReleaseLifecycleProgressionSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get release lifecycle progression spaces params
func (o *GetReleaseLifecycleProgressionSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get release lifecycle progression spaces params
func (o *GetReleaseLifecycleProgressionSpacesParams) WithContext(ctx context.Context) *GetReleaseLifecycleProgressionSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get release lifecycle progression spaces params
func (o *GetReleaseLifecycleProgressionSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get release lifecycle progression spaces params
func (o *GetReleaseLifecycleProgressionSpacesParams) WithHTTPClient(client *http.Client) *GetReleaseLifecycleProgressionSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get release lifecycle progression spaces params
func (o *GetReleaseLifecycleProgressionSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get release lifecycle progression spaces params
func (o *GetReleaseLifecycleProgressionSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetReleaseLifecycleProgressionSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get release lifecycle progression spaces params
func (o *GetReleaseLifecycleProgressionSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get release lifecycle progression spaces params
func (o *GetReleaseLifecycleProgressionSpacesParams) WithID(id string) *GetReleaseLifecycleProgressionSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get release lifecycle progression spaces params
func (o *GetReleaseLifecycleProgressionSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetReleaseLifecycleProgressionSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}