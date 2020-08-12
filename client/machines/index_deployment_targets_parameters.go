// Code generated by go-swagger; DO NOT EDIT.

package machines

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

// NewIndexDeploymentTargetsParams creates a new IndexDeploymentTargetsParams object
// with the default values initialized.
func NewIndexDeploymentTargetsParams() *IndexDeploymentTargetsParams {
	var ()
	return &IndexDeploymentTargetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexDeploymentTargetsParamsWithTimeout creates a new IndexDeploymentTargetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexDeploymentTargetsParamsWithTimeout(timeout time.Duration) *IndexDeploymentTargetsParams {
	var ()
	return &IndexDeploymentTargetsParams{

		timeout: timeout,
	}
}

// NewIndexDeploymentTargetsParamsWithContext creates a new IndexDeploymentTargetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexDeploymentTargetsParamsWithContext(ctx context.Context) *IndexDeploymentTargetsParams {
	var ()
	return &IndexDeploymentTargetsParams{

		Context: ctx,
	}
}

// NewIndexDeploymentTargetsParamsWithHTTPClient creates a new IndexDeploymentTargetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexDeploymentTargetsParamsWithHTTPClient(client *http.Client) *IndexDeploymentTargetsParams {
	var ()
	return &IndexDeploymentTargetsParams{
		HTTPClient: client,
	}
}

/*IndexDeploymentTargetsParams contains all the parameters to send to the API endpoint
for the index deployment targets operation typically these are written to a http.Request
*/
type IndexDeploymentTargetsParams struct {

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

// WithTimeout adds the timeout to the index deployment targets params
func (o *IndexDeploymentTargetsParams) WithTimeout(timeout time.Duration) *IndexDeploymentTargetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index deployment targets params
func (o *IndexDeploymentTargetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index deployment targets params
func (o *IndexDeploymentTargetsParams) WithContext(ctx context.Context) *IndexDeploymentTargetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index deployment targets params
func (o *IndexDeploymentTargetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index deployment targets params
func (o *IndexDeploymentTargetsParams) WithHTTPClient(client *http.Client) *IndexDeploymentTargetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index deployment targets params
func (o *IndexDeploymentTargetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkip adds the skip to the index deployment targets params
func (o *IndexDeploymentTargetsParams) WithSkip(skip *int32) *IndexDeploymentTargetsParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index deployment targets params
func (o *IndexDeploymentTargetsParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index deployment targets params
func (o *IndexDeploymentTargetsParams) WithTake(take *int32) *IndexDeploymentTargetsParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index deployment targets params
func (o *IndexDeploymentTargetsParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexDeploymentTargetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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