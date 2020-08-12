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

// NewGetCertificateUsageSpacesParams creates a new GetCertificateUsageSpacesParams object
// with the default values initialized.
func NewGetCertificateUsageSpacesParams() *GetCertificateUsageSpacesParams {
	var ()
	return &GetCertificateUsageSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCertificateUsageSpacesParamsWithTimeout creates a new GetCertificateUsageSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCertificateUsageSpacesParamsWithTimeout(timeout time.Duration) *GetCertificateUsageSpacesParams {
	var ()
	return &GetCertificateUsageSpacesParams{

		timeout: timeout,
	}
}

// NewGetCertificateUsageSpacesParamsWithContext creates a new GetCertificateUsageSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCertificateUsageSpacesParamsWithContext(ctx context.Context) *GetCertificateUsageSpacesParams {
	var ()
	return &GetCertificateUsageSpacesParams{

		Context: ctx,
	}
}

// NewGetCertificateUsageSpacesParamsWithHTTPClient creates a new GetCertificateUsageSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCertificateUsageSpacesParamsWithHTTPClient(client *http.Client) *GetCertificateUsageSpacesParams {
	var ()
	return &GetCertificateUsageSpacesParams{
		HTTPClient: client,
	}
}

/*GetCertificateUsageSpacesParams contains all the parameters to send to the API endpoint
for the get certificate usage spaces operation typically these are written to a http.Request
*/
type GetCertificateUsageSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the certificate

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get certificate usage spaces params
func (o *GetCertificateUsageSpacesParams) WithTimeout(timeout time.Duration) *GetCertificateUsageSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get certificate usage spaces params
func (o *GetCertificateUsageSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get certificate usage spaces params
func (o *GetCertificateUsageSpacesParams) WithContext(ctx context.Context) *GetCertificateUsageSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get certificate usage spaces params
func (o *GetCertificateUsageSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get certificate usage spaces params
func (o *GetCertificateUsageSpacesParams) WithHTTPClient(client *http.Client) *GetCertificateUsageSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get certificate usage spaces params
func (o *GetCertificateUsageSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get certificate usage spaces params
func (o *GetCertificateUsageSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetCertificateUsageSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get certificate usage spaces params
func (o *GetCertificateUsageSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get certificate usage spaces params
func (o *GetCertificateUsageSpacesParams) WithID(id string) *GetCertificateUsageSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get certificate usage spaces params
func (o *GetCertificateUsageSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCertificateUsageSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
