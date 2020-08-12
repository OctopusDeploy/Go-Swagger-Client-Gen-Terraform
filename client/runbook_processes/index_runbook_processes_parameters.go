// Code generated by go-swagger; DO NOT EDIT.

package runbook_processes

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
	"github.com/go-openapi/swag"
)

// NewIndexRunbookProcessesParams creates a new IndexRunbookProcessesParams object
// with the default values initialized.
func NewIndexRunbookProcessesParams() *IndexRunbookProcessesParams {
	var ()
	return &IndexRunbookProcessesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexRunbookProcessesParamsWithTimeout creates a new IndexRunbookProcessesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexRunbookProcessesParamsWithTimeout(timeout time.Duration) *IndexRunbookProcessesParams {
	var ()
	return &IndexRunbookProcessesParams{

		timeout: timeout,
	}
}

// NewIndexRunbookProcessesParamsWithContext creates a new IndexRunbookProcessesParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexRunbookProcessesParamsWithContext(ctx context.Context) *IndexRunbookProcessesParams {
	var ()
	return &IndexRunbookProcessesParams{

		Context: ctx,
	}
}

// NewIndexRunbookProcessesParamsWithHTTPClient creates a new IndexRunbookProcessesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexRunbookProcessesParamsWithHTTPClient(client *http.Client) *IndexRunbookProcessesParams {
	var ()
	return &IndexRunbookProcessesParams{
		HTTPClient: client,
	}
}

/*IndexRunbookProcessesParams contains all the parameters to send to the API endpoint
for the index runbook processes operation typically these are written to a http.Request
*/
type IndexRunbookProcessesParams struct {

	/*Skip
	  Number of items to skip

	*/
	Skip *int32
	/*Take
	  Number of items to take

	*/
	Take *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the index runbook processes params
func (o *IndexRunbookProcessesParams) WithTimeout(timeout time.Duration) *IndexRunbookProcessesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index runbook processes params
func (o *IndexRunbookProcessesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index runbook processes params
func (o *IndexRunbookProcessesParams) WithContext(ctx context.Context) *IndexRunbookProcessesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index runbook processes params
func (o *IndexRunbookProcessesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index runbook processes params
func (o *IndexRunbookProcessesParams) WithHTTPClient(client *http.Client) *IndexRunbookProcessesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index runbook processes params
func (o *IndexRunbookProcessesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkip adds the skip to the index runbook processes params
func (o *IndexRunbookProcessesParams) WithSkip(skip *int32) *IndexRunbookProcessesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index runbook processes params
func (o *IndexRunbookProcessesParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index runbook processes params
func (o *IndexRunbookProcessesParams) WithTake(take *int32) *IndexRunbookProcessesParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index runbook processes params
func (o *IndexRunbookProcessesParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexRunbookProcessesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Skip != nil {

		// query param skip
		var qrSkip int32
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt32(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	if o.Take != nil {

		// query param take
		var qrTake int32
		if o.Take != nil {
			qrTake = *o.Take
		}
		qTake := swag.FormatInt32(qrTake)
		if qTake != "" {
			if err := r.SetQueryParam("take", qTake); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}