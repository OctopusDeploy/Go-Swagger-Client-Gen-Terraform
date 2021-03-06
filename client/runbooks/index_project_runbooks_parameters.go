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

// NewIndexProjectRunbooksParams creates a new IndexProjectRunbooksParams object
// with the default values initialized.
func NewIndexProjectRunbooksParams() *IndexProjectRunbooksParams {
	var ()
	return &IndexProjectRunbooksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexProjectRunbooksParamsWithTimeout creates a new IndexProjectRunbooksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexProjectRunbooksParamsWithTimeout(timeout time.Duration) *IndexProjectRunbooksParams {
	var ()
	return &IndexProjectRunbooksParams{

		timeout: timeout,
	}
}

// NewIndexProjectRunbooksParamsWithContext creates a new IndexProjectRunbooksParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexProjectRunbooksParamsWithContext(ctx context.Context) *IndexProjectRunbooksParams {
	var ()
	return &IndexProjectRunbooksParams{

		Context: ctx,
	}
}

// NewIndexProjectRunbooksParamsWithHTTPClient creates a new IndexProjectRunbooksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexProjectRunbooksParamsWithHTTPClient(client *http.Client) *IndexProjectRunbooksParams {
	var ()
	return &IndexProjectRunbooksParams{
		HTTPClient: client,
	}
}

/*IndexProjectRunbooksParams contains all the parameters to send to the API endpoint
for the index project runbooks operation typically these are written to a http.Request
*/
type IndexProjectRunbooksParams struct {

	/*ID
	  ID of the Project

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

// WithTimeout adds the timeout to the index project runbooks params
func (o *IndexProjectRunbooksParams) WithTimeout(timeout time.Duration) *IndexProjectRunbooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index project runbooks params
func (o *IndexProjectRunbooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index project runbooks params
func (o *IndexProjectRunbooksParams) WithContext(ctx context.Context) *IndexProjectRunbooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index project runbooks params
func (o *IndexProjectRunbooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index project runbooks params
func (o *IndexProjectRunbooksParams) WithHTTPClient(client *http.Client) *IndexProjectRunbooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index project runbooks params
func (o *IndexProjectRunbooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the index project runbooks params
func (o *IndexProjectRunbooksParams) WithID(id string) *IndexProjectRunbooksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the index project runbooks params
func (o *IndexProjectRunbooksParams) SetID(id string) {
	o.ID = id
}

// WithSkip adds the skip to the index project runbooks params
func (o *IndexProjectRunbooksParams) WithSkip(skip *int32) *IndexProjectRunbooksParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index project runbooks params
func (o *IndexProjectRunbooksParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index project runbooks params
func (o *IndexProjectRunbooksParams) WithTake(take *int32) *IndexProjectRunbooksParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index project runbooks params
func (o *IndexProjectRunbooksParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexProjectRunbooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
