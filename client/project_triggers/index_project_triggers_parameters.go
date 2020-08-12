// Code generated by go-swagger; DO NOT EDIT.

package project_triggers

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

// NewIndexProjectTriggersParams creates a new IndexProjectTriggersParams object
// with the default values initialized.
func NewIndexProjectTriggersParams() *IndexProjectTriggersParams {
	var ()
	return &IndexProjectTriggersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexProjectTriggersParamsWithTimeout creates a new IndexProjectTriggersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexProjectTriggersParamsWithTimeout(timeout time.Duration) *IndexProjectTriggersParams {
	var ()
	return &IndexProjectTriggersParams{

		timeout: timeout,
	}
}

// NewIndexProjectTriggersParamsWithContext creates a new IndexProjectTriggersParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexProjectTriggersParamsWithContext(ctx context.Context) *IndexProjectTriggersParams {
	var ()
	return &IndexProjectTriggersParams{

		Context: ctx,
	}
}

// NewIndexProjectTriggersParamsWithHTTPClient creates a new IndexProjectTriggersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexProjectTriggersParamsWithHTTPClient(client *http.Client) *IndexProjectTriggersParams {
	var ()
	return &IndexProjectTriggersParams{
		HTTPClient: client,
	}
}

/*IndexProjectTriggersParams contains all the parameters to send to the API endpoint
for the index project triggers operation typically these are written to a http.Request
*/
type IndexProjectTriggersParams struct {

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

// WithTimeout adds the timeout to the index project triggers params
func (o *IndexProjectTriggersParams) WithTimeout(timeout time.Duration) *IndexProjectTriggersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index project triggers params
func (o *IndexProjectTriggersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index project triggers params
func (o *IndexProjectTriggersParams) WithContext(ctx context.Context) *IndexProjectTriggersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index project triggers params
func (o *IndexProjectTriggersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index project triggers params
func (o *IndexProjectTriggersParams) WithHTTPClient(client *http.Client) *IndexProjectTriggersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index project triggers params
func (o *IndexProjectTriggersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkip adds the skip to the index project triggers params
func (o *IndexProjectTriggersParams) WithSkip(skip *int32) *IndexProjectTriggersParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index project triggers params
func (o *IndexProjectTriggersParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index project triggers params
func (o *IndexProjectTriggersParams) WithTake(take *int32) *IndexProjectTriggersParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index project triggers params
func (o *IndexProjectTriggersParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexProjectTriggersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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