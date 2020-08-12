// Code generated by go-swagger; DO NOT EDIT.

package certificates

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

// NewGetCertificateExportSpacesParams creates a new GetCertificateExportSpacesParams object
// with the default values initialized.
func NewGetCertificateExportSpacesParams() *GetCertificateExportSpacesParams {
	var ()
	return &GetCertificateExportSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCertificateExportSpacesParamsWithTimeout creates a new GetCertificateExportSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCertificateExportSpacesParamsWithTimeout(timeout time.Duration) *GetCertificateExportSpacesParams {
	var ()
	return &GetCertificateExportSpacesParams{

		timeout: timeout,
	}
}

// NewGetCertificateExportSpacesParamsWithContext creates a new GetCertificateExportSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCertificateExportSpacesParamsWithContext(ctx context.Context) *GetCertificateExportSpacesParams {
	var ()
	return &GetCertificateExportSpacesParams{

		Context: ctx,
	}
}

// NewGetCertificateExportSpacesParamsWithHTTPClient creates a new GetCertificateExportSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCertificateExportSpacesParamsWithHTTPClient(client *http.Client) *GetCertificateExportSpacesParams {
	var ()
	return &GetCertificateExportSpacesParams{
		HTTPClient: client,
	}
}

/*GetCertificateExportSpacesParams contains all the parameters to send to the API endpoint
for the get certificate export spaces operation typically these are written to a http.Request
*/
type GetCertificateExportSpacesParams struct {

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

// WithTimeout adds the timeout to the get certificate export spaces params
func (o *GetCertificateExportSpacesParams) WithTimeout(timeout time.Duration) *GetCertificateExportSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get certificate export spaces params
func (o *GetCertificateExportSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get certificate export spaces params
func (o *GetCertificateExportSpacesParams) WithContext(ctx context.Context) *GetCertificateExportSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get certificate export spaces params
func (o *GetCertificateExportSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get certificate export spaces params
func (o *GetCertificateExportSpacesParams) WithHTTPClient(client *http.Client) *GetCertificateExportSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get certificate export spaces params
func (o *GetCertificateExportSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get certificate export spaces params
func (o *GetCertificateExportSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetCertificateExportSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get certificate export spaces params
func (o *GetCertificateExportSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get certificate export spaces params
func (o *GetCertificateExportSpacesParams) WithID(id string) *GetCertificateExportSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get certificate export spaces params
func (o *GetCertificateExportSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCertificateExportSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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