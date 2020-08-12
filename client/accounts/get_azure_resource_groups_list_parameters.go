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

// NewGetAzureResourceGroupsListParams creates a new GetAzureResourceGroupsListParams object
// with the default values initialized.
func NewGetAzureResourceGroupsListParams() *GetAzureResourceGroupsListParams {
	var ()
	return &GetAzureResourceGroupsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAzureResourceGroupsListParamsWithTimeout creates a new GetAzureResourceGroupsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAzureResourceGroupsListParamsWithTimeout(timeout time.Duration) *GetAzureResourceGroupsListParams {
	var ()
	return &GetAzureResourceGroupsListParams{

		timeout: timeout,
	}
}

// NewGetAzureResourceGroupsListParamsWithContext creates a new GetAzureResourceGroupsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAzureResourceGroupsListParamsWithContext(ctx context.Context) *GetAzureResourceGroupsListParams {
	var ()
	return &GetAzureResourceGroupsListParams{

		Context: ctx,
	}
}

// NewGetAzureResourceGroupsListParamsWithHTTPClient creates a new GetAzureResourceGroupsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAzureResourceGroupsListParamsWithHTTPClient(client *http.Client) *GetAzureResourceGroupsListParams {
	var ()
	return &GetAzureResourceGroupsListParams{
		HTTPClient: client,
	}
}

/*GetAzureResourceGroupsListParams contains all the parameters to send to the API endpoint
for the get azure resource groups list operation typically these are written to a http.Request
*/
type GetAzureResourceGroupsListParams struct {

	/*ID
	  Account Id for Azure

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get azure resource groups list params
func (o *GetAzureResourceGroupsListParams) WithTimeout(timeout time.Duration) *GetAzureResourceGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get azure resource groups list params
func (o *GetAzureResourceGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get azure resource groups list params
func (o *GetAzureResourceGroupsListParams) WithContext(ctx context.Context) *GetAzureResourceGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get azure resource groups list params
func (o *GetAzureResourceGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get azure resource groups list params
func (o *GetAzureResourceGroupsListParams) WithHTTPClient(client *http.Client) *GetAzureResourceGroupsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get azure resource groups list params
func (o *GetAzureResourceGroupsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get azure resource groups list params
func (o *GetAzureResourceGroupsListParams) WithID(id string) *GetAzureResourceGroupsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get azure resource groups list params
func (o *GetAzureResourceGroupsListParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAzureResourceGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
