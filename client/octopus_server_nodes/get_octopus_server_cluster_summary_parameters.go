// Code generated by go-swagger; DO NOT EDIT.

package octopus_server_nodes

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

// NewGetOctopusServerClusterSummaryParams creates a new GetOctopusServerClusterSummaryParams object
// with the default values initialized.
func NewGetOctopusServerClusterSummaryParams() *GetOctopusServerClusterSummaryParams {

	return &GetOctopusServerClusterSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOctopusServerClusterSummaryParamsWithTimeout creates a new GetOctopusServerClusterSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOctopusServerClusterSummaryParamsWithTimeout(timeout time.Duration) *GetOctopusServerClusterSummaryParams {

	return &GetOctopusServerClusterSummaryParams{

		timeout: timeout,
	}
}

// NewGetOctopusServerClusterSummaryParamsWithContext creates a new GetOctopusServerClusterSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOctopusServerClusterSummaryParamsWithContext(ctx context.Context) *GetOctopusServerClusterSummaryParams {

	return &GetOctopusServerClusterSummaryParams{

		Context: ctx,
	}
}

// NewGetOctopusServerClusterSummaryParamsWithHTTPClient creates a new GetOctopusServerClusterSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOctopusServerClusterSummaryParamsWithHTTPClient(client *http.Client) *GetOctopusServerClusterSummaryParams {

	return &GetOctopusServerClusterSummaryParams{
		HTTPClient: client,
	}
}

/*GetOctopusServerClusterSummaryParams contains all the parameters to send to the API endpoint
for the get octopus server cluster summary operation typically these are written to a http.Request
*/
type GetOctopusServerClusterSummaryParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get octopus server cluster summary params
func (o *GetOctopusServerClusterSummaryParams) WithTimeout(timeout time.Duration) *GetOctopusServerClusterSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get octopus server cluster summary params
func (o *GetOctopusServerClusterSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get octopus server cluster summary params
func (o *GetOctopusServerClusterSummaryParams) WithContext(ctx context.Context) *GetOctopusServerClusterSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get octopus server cluster summary params
func (o *GetOctopusServerClusterSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get octopus server cluster summary params
func (o *GetOctopusServerClusterSummaryParams) WithHTTPClient(client *http.Client) *GetOctopusServerClusterSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get octopus server cluster summary params
func (o *GetOctopusServerClusterSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOctopusServerClusterSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}