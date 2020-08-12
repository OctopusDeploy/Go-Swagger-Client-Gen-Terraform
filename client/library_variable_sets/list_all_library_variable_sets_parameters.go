// Code generated by go-swagger; DO NOT EDIT.

package library_variable_sets

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

// NewListAllLibraryVariableSetsParams creates a new ListAllLibraryVariableSetsParams object
// with the default values initialized.
func NewListAllLibraryVariableSetsParams() *ListAllLibraryVariableSetsParams {

	return &ListAllLibraryVariableSetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllLibraryVariableSetsParamsWithTimeout creates a new ListAllLibraryVariableSetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllLibraryVariableSetsParamsWithTimeout(timeout time.Duration) *ListAllLibraryVariableSetsParams {

	return &ListAllLibraryVariableSetsParams{

		timeout: timeout,
	}
}

// NewListAllLibraryVariableSetsParamsWithContext creates a new ListAllLibraryVariableSetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllLibraryVariableSetsParamsWithContext(ctx context.Context) *ListAllLibraryVariableSetsParams {

	return &ListAllLibraryVariableSetsParams{

		Context: ctx,
	}
}

// NewListAllLibraryVariableSetsParamsWithHTTPClient creates a new ListAllLibraryVariableSetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllLibraryVariableSetsParamsWithHTTPClient(client *http.Client) *ListAllLibraryVariableSetsParams {

	return &ListAllLibraryVariableSetsParams{
		HTTPClient: client,
	}
}

/*ListAllLibraryVariableSetsParams contains all the parameters to send to the API endpoint
for the list all library variable sets operation typically these are written to a http.Request
*/
type ListAllLibraryVariableSetsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all library variable sets params
func (o *ListAllLibraryVariableSetsParams) WithTimeout(timeout time.Duration) *ListAllLibraryVariableSetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all library variable sets params
func (o *ListAllLibraryVariableSetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all library variable sets params
func (o *ListAllLibraryVariableSetsParams) WithContext(ctx context.Context) *ListAllLibraryVariableSetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all library variable sets params
func (o *ListAllLibraryVariableSetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all library variable sets params
func (o *ListAllLibraryVariableSetsParams) WithHTTPClient(client *http.Client) *ListAllLibraryVariableSetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all library variable sets params
func (o *ListAllLibraryVariableSetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllLibraryVariableSetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}