// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewIndexTenantsParams creates a new IndexTenantsParams object
// with the default values initialized.
func NewIndexTenantsParams() *IndexTenantsParams {
	var ()
	return &IndexTenantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexTenantsParamsWithTimeout creates a new IndexTenantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexTenantsParamsWithTimeout(timeout time.Duration) *IndexTenantsParams {
	var ()
	return &IndexTenantsParams{

		timeout: timeout,
	}
}

// NewIndexTenantsParamsWithContext creates a new IndexTenantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexTenantsParamsWithContext(ctx context.Context) *IndexTenantsParams {
	var ()
	return &IndexTenantsParams{

		Context: ctx,
	}
}

// NewIndexTenantsParamsWithHTTPClient creates a new IndexTenantsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexTenantsParamsWithHTTPClient(client *http.Client) *IndexTenantsParams {
	var ()
	return &IndexTenantsParams{
		HTTPClient: client,
	}
}

/*IndexTenantsParams contains all the parameters to send to the API endpoint
for the index tenants operation typically these are written to a http.Request
*/
type IndexTenantsParams struct {

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

// WithTimeout adds the timeout to the index tenants params
func (o *IndexTenantsParams) WithTimeout(timeout time.Duration) *IndexTenantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index tenants params
func (o *IndexTenantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index tenants params
func (o *IndexTenantsParams) WithContext(ctx context.Context) *IndexTenantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index tenants params
func (o *IndexTenantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index tenants params
func (o *IndexTenantsParams) WithHTTPClient(client *http.Client) *IndexTenantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index tenants params
func (o *IndexTenantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkip adds the skip to the index tenants params
func (o *IndexTenantsParams) WithSkip(skip *int32) *IndexTenantsParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index tenants params
func (o *IndexTenantsParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index tenants params
func (o *IndexTenantsParams) WithTake(take *int32) *IndexTenantsParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index tenants params
func (o *IndexTenantsParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexTenantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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