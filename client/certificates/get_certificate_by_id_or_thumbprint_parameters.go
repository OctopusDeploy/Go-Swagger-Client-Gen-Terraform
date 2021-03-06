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

// NewGetCertificateByIDOrThumbprintParams creates a new GetCertificateByIDOrThumbprintParams object
// with the default values initialized.
func NewGetCertificateByIDOrThumbprintParams() *GetCertificateByIDOrThumbprintParams {
	var ()
	return &GetCertificateByIDOrThumbprintParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCertificateByIDOrThumbprintParamsWithTimeout creates a new GetCertificateByIDOrThumbprintParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCertificateByIDOrThumbprintParamsWithTimeout(timeout time.Duration) *GetCertificateByIDOrThumbprintParams {
	var ()
	return &GetCertificateByIDOrThumbprintParams{

		timeout: timeout,
	}
}

// NewGetCertificateByIDOrThumbprintParamsWithContext creates a new GetCertificateByIDOrThumbprintParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCertificateByIDOrThumbprintParamsWithContext(ctx context.Context) *GetCertificateByIDOrThumbprintParams {
	var ()
	return &GetCertificateByIDOrThumbprintParams{

		Context: ctx,
	}
}

// NewGetCertificateByIDOrThumbprintParamsWithHTTPClient creates a new GetCertificateByIDOrThumbprintParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCertificateByIDOrThumbprintParamsWithHTTPClient(client *http.Client) *GetCertificateByIDOrThumbprintParams {
	var ()
	return &GetCertificateByIDOrThumbprintParams{
		HTTPClient: client,
	}
}

/*GetCertificateByIDOrThumbprintParams contains all the parameters to send to the API endpoint
for the get certificate by Id or thumbprint operation typically these are written to a http.Request
*/
type GetCertificateByIDOrThumbprintParams struct {

	/*ID
	  ID or thumbprint of certificate

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get certificate by Id or thumbprint params
func (o *GetCertificateByIDOrThumbprintParams) WithTimeout(timeout time.Duration) *GetCertificateByIDOrThumbprintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get certificate by Id or thumbprint params
func (o *GetCertificateByIDOrThumbprintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get certificate by Id or thumbprint params
func (o *GetCertificateByIDOrThumbprintParams) WithContext(ctx context.Context) *GetCertificateByIDOrThumbprintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get certificate by Id or thumbprint params
func (o *GetCertificateByIDOrThumbprintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get certificate by Id or thumbprint params
func (o *GetCertificateByIDOrThumbprintParams) WithHTTPClient(client *http.Client) *GetCertificateByIDOrThumbprintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get certificate by Id or thumbprint params
func (o *GetCertificateByIDOrThumbprintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get certificate by Id or thumbprint params
func (o *GetCertificateByIDOrThumbprintParams) WithID(id string) *GetCertificateByIDOrThumbprintParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get certificate by Id or thumbprint params
func (o *GetCertificateByIDOrThumbprintParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCertificateByIDOrThumbprintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
