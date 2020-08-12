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

// NewListAllAccountsParams creates a new ListAllAccountsParams object
// with the default values initialized.
func NewListAllAccountsParams() *ListAllAccountsParams {

	return &ListAllAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllAccountsParamsWithTimeout creates a new ListAllAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllAccountsParamsWithTimeout(timeout time.Duration) *ListAllAccountsParams {

	return &ListAllAccountsParams{

		timeout: timeout,
	}
}

// NewListAllAccountsParamsWithContext creates a new ListAllAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllAccountsParamsWithContext(ctx context.Context) *ListAllAccountsParams {

	return &ListAllAccountsParams{

		Context: ctx,
	}
}

// NewListAllAccountsParamsWithHTTPClient creates a new ListAllAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllAccountsParamsWithHTTPClient(client *http.Client) *ListAllAccountsParams {

	return &ListAllAccountsParams{
		HTTPClient: client,
	}
}

/*ListAllAccountsParams contains all the parameters to send to the API endpoint
for the list all accounts operation typically these are written to a http.Request
*/
type ListAllAccountsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all accounts params
func (o *ListAllAccountsParams) WithTimeout(timeout time.Duration) *ListAllAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all accounts params
func (o *ListAllAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all accounts params
func (o *ListAllAccountsParams) WithContext(ctx context.Context) *ListAllAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all accounts params
func (o *ListAllAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all accounts params
func (o *ListAllAccountsParams) WithHTTPClient(client *http.Client) *ListAllAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all accounts params
func (o *ListAllAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
