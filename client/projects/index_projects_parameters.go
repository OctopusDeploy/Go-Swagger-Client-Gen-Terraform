// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewIndexProjectsParams creates a new IndexProjectsParams object
// with the default values initialized.
func NewIndexProjectsParams() *IndexProjectsParams {
	var ()
	return &IndexProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexProjectsParamsWithTimeout creates a new IndexProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexProjectsParamsWithTimeout(timeout time.Duration) *IndexProjectsParams {
	var ()
	return &IndexProjectsParams{

		timeout: timeout,
	}
}

// NewIndexProjectsParamsWithContext creates a new IndexProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexProjectsParamsWithContext(ctx context.Context) *IndexProjectsParams {
	var ()
	return &IndexProjectsParams{

		Context: ctx,
	}
}

// NewIndexProjectsParamsWithHTTPClient creates a new IndexProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexProjectsParamsWithHTTPClient(client *http.Client) *IndexProjectsParams {
	var ()
	return &IndexProjectsParams{
		HTTPClient: client,
	}
}

/*IndexProjectsParams contains all the parameters to send to the API endpoint
for the index projects operation typically these are written to a http.Request
*/
type IndexProjectsParams struct {

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

// WithTimeout adds the timeout to the index projects params
func (o *IndexProjectsParams) WithTimeout(timeout time.Duration) *IndexProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index projects params
func (o *IndexProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index projects params
func (o *IndexProjectsParams) WithContext(ctx context.Context) *IndexProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index projects params
func (o *IndexProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index projects params
func (o *IndexProjectsParams) WithHTTPClient(client *http.Client) *IndexProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index projects params
func (o *IndexProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkip adds the skip to the index projects params
func (o *IndexProjectsParams) WithSkip(skip *int32) *IndexProjectsParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index projects params
func (o *IndexProjectsParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index projects params
func (o *IndexProjectsParams) WithTake(take *int32) *IndexProjectsParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index projects params
func (o *IndexProjectsParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
