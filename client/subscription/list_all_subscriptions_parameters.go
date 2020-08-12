// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// NewListAllSubscriptionsParams creates a new ListAllSubscriptionsParams object
// with the default values initialized.
func NewListAllSubscriptionsParams() *ListAllSubscriptionsParams {

	return &ListAllSubscriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllSubscriptionsParamsWithTimeout creates a new ListAllSubscriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllSubscriptionsParamsWithTimeout(timeout time.Duration) *ListAllSubscriptionsParams {

	return &ListAllSubscriptionsParams{

		timeout: timeout,
	}
}

// NewListAllSubscriptionsParamsWithContext creates a new ListAllSubscriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllSubscriptionsParamsWithContext(ctx context.Context) *ListAllSubscriptionsParams {

	return &ListAllSubscriptionsParams{

		Context: ctx,
	}
}

// NewListAllSubscriptionsParamsWithHTTPClient creates a new ListAllSubscriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllSubscriptionsParamsWithHTTPClient(client *http.Client) *ListAllSubscriptionsParams {

	return &ListAllSubscriptionsParams{
		HTTPClient: client,
	}
}

/*ListAllSubscriptionsParams contains all the parameters to send to the API endpoint
for the list all subscriptions operation typically these are written to a http.Request
*/
type ListAllSubscriptionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all subscriptions params
func (o *ListAllSubscriptionsParams) WithTimeout(timeout time.Duration) *ListAllSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all subscriptions params
func (o *ListAllSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all subscriptions params
func (o *ListAllSubscriptionsParams) WithContext(ctx context.Context) *ListAllSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all subscriptions params
func (o *ListAllSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all subscriptions params
func (o *ListAllSubscriptionsParams) WithHTTPClient(client *http.Client) *ListAllSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all subscriptions params
func (o *ListAllSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}