// Code generated by go-swagger; DO NOT EDIT.

package project_groups

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

// NewListAllProjectGroupsParams creates a new ListAllProjectGroupsParams object
// with the default values initialized.
func NewListAllProjectGroupsParams() *ListAllProjectGroupsParams {

	return &ListAllProjectGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllProjectGroupsParamsWithTimeout creates a new ListAllProjectGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllProjectGroupsParamsWithTimeout(timeout time.Duration) *ListAllProjectGroupsParams {

	return &ListAllProjectGroupsParams{

		timeout: timeout,
	}
}

// NewListAllProjectGroupsParamsWithContext creates a new ListAllProjectGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllProjectGroupsParamsWithContext(ctx context.Context) *ListAllProjectGroupsParams {

	return &ListAllProjectGroupsParams{

		Context: ctx,
	}
}

// NewListAllProjectGroupsParamsWithHTTPClient creates a new ListAllProjectGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllProjectGroupsParamsWithHTTPClient(client *http.Client) *ListAllProjectGroupsParams {

	return &ListAllProjectGroupsParams{
		HTTPClient: client,
	}
}

/*ListAllProjectGroupsParams contains all the parameters to send to the API endpoint
for the list all project groups operation typically these are written to a http.Request
*/
type ListAllProjectGroupsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all project groups params
func (o *ListAllProjectGroupsParams) WithTimeout(timeout time.Duration) *ListAllProjectGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all project groups params
func (o *ListAllProjectGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all project groups params
func (o *ListAllProjectGroupsParams) WithContext(ctx context.Context) *ListAllProjectGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all project groups params
func (o *ListAllProjectGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all project groups params
func (o *ListAllProjectGroupsParams) WithHTTPClient(client *http.Client) *ListAllProjectGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all project groups params
func (o *ListAllProjectGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllProjectGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
