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

// NewListAllAccountsSpacesParams creates a new ListAllAccountsSpacesParams object
// with the default values initialized.
func NewListAllAccountsSpacesParams() *ListAllAccountsSpacesParams {
	var ()
	return &ListAllAccountsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllAccountsSpacesParamsWithTimeout creates a new ListAllAccountsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllAccountsSpacesParamsWithTimeout(timeout time.Duration) *ListAllAccountsSpacesParams {
	var ()
	return &ListAllAccountsSpacesParams{

		timeout: timeout,
	}
}

// NewListAllAccountsSpacesParamsWithContext creates a new ListAllAccountsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllAccountsSpacesParamsWithContext(ctx context.Context) *ListAllAccountsSpacesParams {
	var ()
	return &ListAllAccountsSpacesParams{

		Context: ctx,
	}
}

// NewListAllAccountsSpacesParamsWithHTTPClient creates a new ListAllAccountsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllAccountsSpacesParamsWithHTTPClient(client *http.Client) *ListAllAccountsSpacesParams {
	var ()
	return &ListAllAccountsSpacesParams{
		HTTPClient: client,
	}
}

/*ListAllAccountsSpacesParams contains all the parameters to send to the API endpoint
for the list all accounts spaces operation typically these are written to a http.Request
*/
type ListAllAccountsSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all accounts spaces params
func (o *ListAllAccountsSpacesParams) WithTimeout(timeout time.Duration) *ListAllAccountsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all accounts spaces params
func (o *ListAllAccountsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all accounts spaces params
func (o *ListAllAccountsSpacesParams) WithContext(ctx context.Context) *ListAllAccountsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all accounts spaces params
func (o *ListAllAccountsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all accounts spaces params
func (o *ListAllAccountsSpacesParams) WithHTTPClient(client *http.Client) *ListAllAccountsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all accounts spaces params
func (o *ListAllAccountsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the list all accounts spaces params
func (o *ListAllAccountsSpacesParams) WithBaseSpaceID(baseSpaceID string) *ListAllAccountsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the list all accounts spaces params
func (o *ListAllAccountsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllAccountsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
