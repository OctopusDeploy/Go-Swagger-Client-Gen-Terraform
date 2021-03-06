// Code generated by go-swagger; DO NOT EDIT.

package packages

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

// NewGetPackageSearchActionOldSpacesParams creates a new GetPackageSearchActionOldSpacesParams object
// with the default values initialized.
func NewGetPackageSearchActionOldSpacesParams() *GetPackageSearchActionOldSpacesParams {
	var ()
	return &GetPackageSearchActionOldSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackageSearchActionOldSpacesParamsWithTimeout creates a new GetPackageSearchActionOldSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPackageSearchActionOldSpacesParamsWithTimeout(timeout time.Duration) *GetPackageSearchActionOldSpacesParams {
	var ()
	return &GetPackageSearchActionOldSpacesParams{

		timeout: timeout,
	}
}

// NewGetPackageSearchActionOldSpacesParamsWithContext creates a new GetPackageSearchActionOldSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPackageSearchActionOldSpacesParamsWithContext(ctx context.Context) *GetPackageSearchActionOldSpacesParams {
	var ()
	return &GetPackageSearchActionOldSpacesParams{

		Context: ctx,
	}
}

// NewGetPackageSearchActionOldSpacesParamsWithHTTPClient creates a new GetPackageSearchActionOldSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPackageSearchActionOldSpacesParamsWithHTTPClient(client *http.Client) *GetPackageSearchActionOldSpacesParams {
	var ()
	return &GetPackageSearchActionOldSpacesParams{
		HTTPClient: client,
	}
}

/*GetPackageSearchActionOldSpacesParams contains all the parameters to send to the API endpoint
for the get package search action old spaces operation typically these are written to a http.Request
*/
type GetPackageSearchActionOldSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the feed

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get package search action old spaces params
func (o *GetPackageSearchActionOldSpacesParams) WithTimeout(timeout time.Duration) *GetPackageSearchActionOldSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get package search action old spaces params
func (o *GetPackageSearchActionOldSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get package search action old spaces params
func (o *GetPackageSearchActionOldSpacesParams) WithContext(ctx context.Context) *GetPackageSearchActionOldSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get package search action old spaces params
func (o *GetPackageSearchActionOldSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get package search action old spaces params
func (o *GetPackageSearchActionOldSpacesParams) WithHTTPClient(client *http.Client) *GetPackageSearchActionOldSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get package search action old spaces params
func (o *GetPackageSearchActionOldSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get package search action old spaces params
func (o *GetPackageSearchActionOldSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetPackageSearchActionOldSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get package search action old spaces params
func (o *GetPackageSearchActionOldSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get package search action old spaces params
func (o *GetPackageSearchActionOldSpacesParams) WithID(id string) *GetPackageSearchActionOldSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get package search action old spaces params
func (o *GetPackageSearchActionOldSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackageSearchActionOldSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
