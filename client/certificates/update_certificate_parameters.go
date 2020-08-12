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

	"models"
)

// NewUpdateCertificateParams creates a new UpdateCertificateParams object
// with the default values initialized.
func NewUpdateCertificateParams() *UpdateCertificateParams {
	var ()
	return &UpdateCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCertificateParamsWithTimeout creates a new UpdateCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCertificateParamsWithTimeout(timeout time.Duration) *UpdateCertificateParams {
	var ()
	return &UpdateCertificateParams{

		timeout: timeout,
	}
}

// NewUpdateCertificateParamsWithContext creates a new UpdateCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCertificateParamsWithContext(ctx context.Context) *UpdateCertificateParams {
	var ()
	return &UpdateCertificateParams{

		Context: ctx,
	}
}

// NewUpdateCertificateParamsWithHTTPClient creates a new UpdateCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCertificateParamsWithHTTPClient(client *http.Client) *UpdateCertificateParams {
	var ()
	return &UpdateCertificateParams{
		HTTPClient: client,
	}
}

/*UpdateCertificateParams contains all the parameters to send to the API endpoint
for the update certificate operation typically these are written to a http.Request
*/
type UpdateCertificateParams struct {

	/*CertificateResource
	  The CertificateResource resource to create

	*/
	CertificateResource *models.CertificateResource
	/*ID
	  ID of the CertificateResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update certificate params
func (o *UpdateCertificateParams) WithTimeout(timeout time.Duration) *UpdateCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update certificate params
func (o *UpdateCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update certificate params
func (o *UpdateCertificateParams) WithContext(ctx context.Context) *UpdateCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update certificate params
func (o *UpdateCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update certificate params
func (o *UpdateCertificateParams) WithHTTPClient(client *http.Client) *UpdateCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update certificate params
func (o *UpdateCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateResource adds the certificateResource to the update certificate params
func (o *UpdateCertificateParams) WithCertificateResource(certificateResource *models.CertificateResource) *UpdateCertificateParams {
	o.SetCertificateResource(certificateResource)
	return o
}

// SetCertificateResource adds the certificateResource to the update certificate params
func (o *UpdateCertificateParams) SetCertificateResource(certificateResource *models.CertificateResource) {
	o.CertificateResource = certificateResource
}

// WithID adds the id to the update certificate params
func (o *UpdateCertificateParams) WithID(id string) *UpdateCertificateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update certificate params
func (o *UpdateCertificateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CertificateResource != nil {
		if err := r.SetBodyParam(o.CertificateResource); err != nil {
			return err
		}
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