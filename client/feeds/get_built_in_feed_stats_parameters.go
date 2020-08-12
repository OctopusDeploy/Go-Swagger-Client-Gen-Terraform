// Code generated by go-swagger; DO NOT EDIT.

package feeds

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

// NewGetBuiltInFeedStatsParams creates a new GetBuiltInFeedStatsParams object
// with the default values initialized.
func NewGetBuiltInFeedStatsParams() *GetBuiltInFeedStatsParams {

	return &GetBuiltInFeedStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuiltInFeedStatsParamsWithTimeout creates a new GetBuiltInFeedStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBuiltInFeedStatsParamsWithTimeout(timeout time.Duration) *GetBuiltInFeedStatsParams {

	return &GetBuiltInFeedStatsParams{

		timeout: timeout,
	}
}

// NewGetBuiltInFeedStatsParamsWithContext creates a new GetBuiltInFeedStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBuiltInFeedStatsParamsWithContext(ctx context.Context) *GetBuiltInFeedStatsParams {

	return &GetBuiltInFeedStatsParams{

		Context: ctx,
	}
}

// NewGetBuiltInFeedStatsParamsWithHTTPClient creates a new GetBuiltInFeedStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBuiltInFeedStatsParamsWithHTTPClient(client *http.Client) *GetBuiltInFeedStatsParams {

	return &GetBuiltInFeedStatsParams{
		HTTPClient: client,
	}
}

/*GetBuiltInFeedStatsParams contains all the parameters to send to the API endpoint
for the get built in feed stats operation typically these are written to a http.Request
*/
type GetBuiltInFeedStatsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get built in feed stats params
func (o *GetBuiltInFeedStatsParams) WithTimeout(timeout time.Duration) *GetBuiltInFeedStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get built in feed stats params
func (o *GetBuiltInFeedStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get built in feed stats params
func (o *GetBuiltInFeedStatsParams) WithContext(ctx context.Context) *GetBuiltInFeedStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get built in feed stats params
func (o *GetBuiltInFeedStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get built in feed stats params
func (o *GetBuiltInFeedStatsParams) WithHTTPClient(client *http.Client) *GetBuiltInFeedStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get built in feed stats params
func (o *GetBuiltInFeedStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuiltInFeedStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}