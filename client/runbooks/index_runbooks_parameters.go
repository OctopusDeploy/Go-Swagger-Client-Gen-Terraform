// Code generated by go-swagger; DO NOT EDIT.

package runbooks

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

// NewIndexRunbooksParams creates a new IndexRunbooksParams object
// with the default values initialized.
func NewIndexRunbooksParams() *IndexRunbooksParams {
	var ()
	return &IndexRunbooksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexRunbooksParamsWithTimeout creates a new IndexRunbooksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexRunbooksParamsWithTimeout(timeout time.Duration) *IndexRunbooksParams {
	var ()
	return &IndexRunbooksParams{

		timeout: timeout,
	}
}

// NewIndexRunbooksParamsWithContext creates a new IndexRunbooksParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexRunbooksParamsWithContext(ctx context.Context) *IndexRunbooksParams {
	var ()
	return &IndexRunbooksParams{

		Context: ctx,
	}
}

// NewIndexRunbooksParamsWithHTTPClient creates a new IndexRunbooksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexRunbooksParamsWithHTTPClient(client *http.Client) *IndexRunbooksParams {
	var ()
	return &IndexRunbooksParams{
		HTTPClient: client,
	}
}

/*IndexRunbooksParams contains all the parameters to send to the API endpoint
for the index runbooks operation typically these are written to a http.Request
*/
type IndexRunbooksParams struct {

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

// WithTimeout adds the timeout to the index runbooks params
func (o *IndexRunbooksParams) WithTimeout(timeout time.Duration) *IndexRunbooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index runbooks params
func (o *IndexRunbooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index runbooks params
func (o *IndexRunbooksParams) WithContext(ctx context.Context) *IndexRunbooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index runbooks params
func (o *IndexRunbooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index runbooks params
func (o *IndexRunbooksParams) WithHTTPClient(client *http.Client) *IndexRunbooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index runbooks params
func (o *IndexRunbooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkip adds the skip to the index runbooks params
func (o *IndexRunbooksParams) WithSkip(skip *int32) *IndexRunbooksParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index runbooks params
func (o *IndexRunbooksParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index runbooks params
func (o *IndexRunbooksParams) WithTake(take *int32) *IndexRunbooksParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index runbooks params
func (o *IndexRunbooksParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexRunbooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
