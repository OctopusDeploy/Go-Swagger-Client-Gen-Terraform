// Code generated by go-swagger; DO NOT EDIT.

package runbook_runs

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

// NewIndexRunbookRunsParams creates a new IndexRunbookRunsParams object
// with the default values initialized.
func NewIndexRunbookRunsParams() *IndexRunbookRunsParams {
	var ()
	return &IndexRunbookRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexRunbookRunsParamsWithTimeout creates a new IndexRunbookRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexRunbookRunsParamsWithTimeout(timeout time.Duration) *IndexRunbookRunsParams {
	var ()
	return &IndexRunbookRunsParams{

		timeout: timeout,
	}
}

// NewIndexRunbookRunsParamsWithContext creates a new IndexRunbookRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexRunbookRunsParamsWithContext(ctx context.Context) *IndexRunbookRunsParams {
	var ()
	return &IndexRunbookRunsParams{

		Context: ctx,
	}
}

// NewIndexRunbookRunsParamsWithHTTPClient creates a new IndexRunbookRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexRunbookRunsParamsWithHTTPClient(client *http.Client) *IndexRunbookRunsParams {
	var ()
	return &IndexRunbookRunsParams{
		HTTPClient: client,
	}
}

/*IndexRunbookRunsParams contains all the parameters to send to the API endpoint
for the index runbook runs operation typically these are written to a http.Request
*/
type IndexRunbookRunsParams struct {

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

// WithTimeout adds the timeout to the index runbook runs params
func (o *IndexRunbookRunsParams) WithTimeout(timeout time.Duration) *IndexRunbookRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index runbook runs params
func (o *IndexRunbookRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index runbook runs params
func (o *IndexRunbookRunsParams) WithContext(ctx context.Context) *IndexRunbookRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index runbook runs params
func (o *IndexRunbookRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index runbook runs params
func (o *IndexRunbookRunsParams) WithHTTPClient(client *http.Client) *IndexRunbookRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index runbook runs params
func (o *IndexRunbookRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkip adds the skip to the index runbook runs params
func (o *IndexRunbookRunsParams) WithSkip(skip *int32) *IndexRunbookRunsParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index runbook runs params
func (o *IndexRunbookRunsParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index runbook runs params
func (o *IndexRunbookRunsParams) WithTake(take *int32) *IndexRunbookRunsParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index runbook runs params
func (o *IndexRunbookRunsParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexRunbookRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
