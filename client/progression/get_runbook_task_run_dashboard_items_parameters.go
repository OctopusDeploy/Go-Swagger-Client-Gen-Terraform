// Code generated by go-swagger; DO NOT EDIT.

package progression

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

// NewGetRunbookTaskRunDashboardItemsParams creates a new GetRunbookTaskRunDashboardItemsParams object
// with the default values initialized.
func NewGetRunbookTaskRunDashboardItemsParams() *GetRunbookTaskRunDashboardItemsParams {

	return &GetRunbookTaskRunDashboardItemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunbookTaskRunDashboardItemsParamsWithTimeout creates a new GetRunbookTaskRunDashboardItemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunbookTaskRunDashboardItemsParamsWithTimeout(timeout time.Duration) *GetRunbookTaskRunDashboardItemsParams {

	return &GetRunbookTaskRunDashboardItemsParams{

		timeout: timeout,
	}
}

// NewGetRunbookTaskRunDashboardItemsParamsWithContext creates a new GetRunbookTaskRunDashboardItemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunbookTaskRunDashboardItemsParamsWithContext(ctx context.Context) *GetRunbookTaskRunDashboardItemsParams {

	return &GetRunbookTaskRunDashboardItemsParams{

		Context: ctx,
	}
}

// NewGetRunbookTaskRunDashboardItemsParamsWithHTTPClient creates a new GetRunbookTaskRunDashboardItemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunbookTaskRunDashboardItemsParamsWithHTTPClient(client *http.Client) *GetRunbookTaskRunDashboardItemsParams {

	return &GetRunbookTaskRunDashboardItemsParams{
		HTTPClient: client,
	}
}

/*GetRunbookTaskRunDashboardItemsParams contains all the parameters to send to the API endpoint
for the get runbook task run dashboard items operation typically these are written to a http.Request
*/
type GetRunbookTaskRunDashboardItemsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runbook task run dashboard items params
func (o *GetRunbookTaskRunDashboardItemsParams) WithTimeout(timeout time.Duration) *GetRunbookTaskRunDashboardItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runbook task run dashboard items params
func (o *GetRunbookTaskRunDashboardItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runbook task run dashboard items params
func (o *GetRunbookTaskRunDashboardItemsParams) WithContext(ctx context.Context) *GetRunbookTaskRunDashboardItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runbook task run dashboard items params
func (o *GetRunbookTaskRunDashboardItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runbook task run dashboard items params
func (o *GetRunbookTaskRunDashboardItemsParams) WithHTTPClient(client *http.Client) *GetRunbookTaskRunDashboardItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runbook task run dashboard items params
func (o *GetRunbookTaskRunDashboardItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunbookTaskRunDashboardItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
