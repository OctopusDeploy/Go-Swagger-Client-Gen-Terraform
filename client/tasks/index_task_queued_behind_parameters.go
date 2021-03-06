// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewIndexTaskQueuedBehindParams creates a new IndexTaskQueuedBehindParams object
// with the default values initialized.
func NewIndexTaskQueuedBehindParams() *IndexTaskQueuedBehindParams {
	var ()
	return &IndexTaskQueuedBehindParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexTaskQueuedBehindParamsWithTimeout creates a new IndexTaskQueuedBehindParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexTaskQueuedBehindParamsWithTimeout(timeout time.Duration) *IndexTaskQueuedBehindParams {
	var ()
	return &IndexTaskQueuedBehindParams{

		timeout: timeout,
	}
}

// NewIndexTaskQueuedBehindParamsWithContext creates a new IndexTaskQueuedBehindParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexTaskQueuedBehindParamsWithContext(ctx context.Context) *IndexTaskQueuedBehindParams {
	var ()
	return &IndexTaskQueuedBehindParams{

		Context: ctx,
	}
}

// NewIndexTaskQueuedBehindParamsWithHTTPClient creates a new IndexTaskQueuedBehindParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexTaskQueuedBehindParamsWithHTTPClient(client *http.Client) *IndexTaskQueuedBehindParams {
	var ()
	return &IndexTaskQueuedBehindParams{
		HTTPClient: client,
	}
}

/*IndexTaskQueuedBehindParams contains all the parameters to send to the API endpoint
for the index task queued behind operation typically these are written to a http.Request
*/
type IndexTaskQueuedBehindParams struct {

	/*ID
	  ID of the Task

	*/
	ID string
	/*Skip
	  Number of items to skip

	*/
	Skip *int32
	/*Take
	  Number of items per page

	*/
	Take *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) WithTimeout(timeout time.Duration) *IndexTaskQueuedBehindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) WithContext(ctx context.Context) *IndexTaskQueuedBehindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) WithHTTPClient(client *http.Client) *IndexTaskQueuedBehindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) WithID(id string) *IndexTaskQueuedBehindParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) SetID(id string) {
	o.ID = id
}

// WithSkip adds the skip to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) WithSkip(skip *int32) *IndexTaskQueuedBehindParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) WithTake(take *int32) *IndexTaskQueuedBehindParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index task queued behind params
func (o *IndexTaskQueuedBehindParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexTaskQueuedBehindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

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
