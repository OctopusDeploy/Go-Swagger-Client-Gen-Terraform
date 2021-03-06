// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewGetAccountPublicKeyDownloadParams creates a new GetAccountPublicKeyDownloadParams object
// with the default values initialized.
func NewGetAccountPublicKeyDownloadParams() *GetAccountPublicKeyDownloadParams {
	var ()
	return &GetAccountPublicKeyDownloadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountPublicKeyDownloadParamsWithTimeout creates a new GetAccountPublicKeyDownloadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountPublicKeyDownloadParamsWithTimeout(timeout time.Duration) *GetAccountPublicKeyDownloadParams {
	var ()
	return &GetAccountPublicKeyDownloadParams{

		timeout: timeout,
	}
}

// NewGetAccountPublicKeyDownloadParamsWithContext creates a new GetAccountPublicKeyDownloadParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountPublicKeyDownloadParamsWithContext(ctx context.Context) *GetAccountPublicKeyDownloadParams {
	var ()
	return &GetAccountPublicKeyDownloadParams{

		Context: ctx,
	}
}

// NewGetAccountPublicKeyDownloadParamsWithHTTPClient creates a new GetAccountPublicKeyDownloadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountPublicKeyDownloadParamsWithHTTPClient(client *http.Client) *GetAccountPublicKeyDownloadParams {
	var ()
	return &GetAccountPublicKeyDownloadParams{
		HTTPClient: client,
	}
}

/*GetAccountPublicKeyDownloadParams contains all the parameters to send to the API endpoint
for the get account public key download operation typically these are written to a http.Request
*/
type GetAccountPublicKeyDownloadParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account public key download params
func (o *GetAccountPublicKeyDownloadParams) WithTimeout(timeout time.Duration) *GetAccountPublicKeyDownloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account public key download params
func (o *GetAccountPublicKeyDownloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account public key download params
func (o *GetAccountPublicKeyDownloadParams) WithContext(ctx context.Context) *GetAccountPublicKeyDownloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account public key download params
func (o *GetAccountPublicKeyDownloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account public key download params
func (o *GetAccountPublicKeyDownloadParams) WithHTTPClient(client *http.Client) *GetAccountPublicKeyDownloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account public key download params
func (o *GetAccountPublicKeyDownloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get account public key download params
func (o *GetAccountPublicKeyDownloadParams) WithID(id string) *GetAccountPublicKeyDownloadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get account public key download params
func (o *GetAccountPublicKeyDownloadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountPublicKeyDownloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
