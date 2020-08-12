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

// NewGetAccountUsageListParams creates a new GetAccountUsageListParams object
// with the default values initialized.
func NewGetAccountUsageListParams() *GetAccountUsageListParams {
	var ()
	return &GetAccountUsageListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountUsageListParamsWithTimeout creates a new GetAccountUsageListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountUsageListParamsWithTimeout(timeout time.Duration) *GetAccountUsageListParams {
	var ()
	return &GetAccountUsageListParams{

		timeout: timeout,
	}
}

// NewGetAccountUsageListParamsWithContext creates a new GetAccountUsageListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountUsageListParamsWithContext(ctx context.Context) *GetAccountUsageListParams {
	var ()
	return &GetAccountUsageListParams{

		Context: ctx,
	}
}

// NewGetAccountUsageListParamsWithHTTPClient creates a new GetAccountUsageListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountUsageListParamsWithHTTPClient(client *http.Client) *GetAccountUsageListParams {
	var ()
	return &GetAccountUsageListParams{
		HTTPClient: client,
	}
}

/*GetAccountUsageListParams contains all the parameters to send to the API endpoint
for the get account usage list operation typically these are written to a http.Request
*/
type GetAccountUsageListParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account usage list params
func (o *GetAccountUsageListParams) WithTimeout(timeout time.Duration) *GetAccountUsageListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account usage list params
func (o *GetAccountUsageListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account usage list params
func (o *GetAccountUsageListParams) WithContext(ctx context.Context) *GetAccountUsageListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account usage list params
func (o *GetAccountUsageListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account usage list params
func (o *GetAccountUsageListParams) WithHTTPClient(client *http.Client) *GetAccountUsageListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account usage list params
func (o *GetAccountUsageListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get account usage list params
func (o *GetAccountUsageListParams) WithID(id string) *GetAccountUsageListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get account usage list params
func (o *GetAccountUsageListParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountUsageListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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