// Code generated by go-swagger; DO NOT EDIT.

package lets_encrypt

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

// NewGetLetsEncryptConfigurationParams creates a new GetLetsEncryptConfigurationParams object
// with the default values initialized.
func NewGetLetsEncryptConfigurationParams() *GetLetsEncryptConfigurationParams {

	return &GetLetsEncryptConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLetsEncryptConfigurationParamsWithTimeout creates a new GetLetsEncryptConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLetsEncryptConfigurationParamsWithTimeout(timeout time.Duration) *GetLetsEncryptConfigurationParams {

	return &GetLetsEncryptConfigurationParams{

		timeout: timeout,
	}
}

// NewGetLetsEncryptConfigurationParamsWithContext creates a new GetLetsEncryptConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLetsEncryptConfigurationParamsWithContext(ctx context.Context) *GetLetsEncryptConfigurationParams {

	return &GetLetsEncryptConfigurationParams{

		Context: ctx,
	}
}

// NewGetLetsEncryptConfigurationParamsWithHTTPClient creates a new GetLetsEncryptConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLetsEncryptConfigurationParamsWithHTTPClient(client *http.Client) *GetLetsEncryptConfigurationParams {

	return &GetLetsEncryptConfigurationParams{
		HTTPClient: client,
	}
}

/*GetLetsEncryptConfigurationParams contains all the parameters to send to the API endpoint
for the get lets encrypt configuration operation typically these are written to a http.Request
*/
type GetLetsEncryptConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lets encrypt configuration params
func (o *GetLetsEncryptConfigurationParams) WithTimeout(timeout time.Duration) *GetLetsEncryptConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lets encrypt configuration params
func (o *GetLetsEncryptConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lets encrypt configuration params
func (o *GetLetsEncryptConfigurationParams) WithContext(ctx context.Context) *GetLetsEncryptConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lets encrypt configuration params
func (o *GetLetsEncryptConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lets encrypt configuration params
func (o *GetLetsEncryptConfigurationParams) WithHTTPClient(client *http.Client) *GetLetsEncryptConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lets encrypt configuration params
func (o *GetLetsEncryptConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLetsEncryptConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}