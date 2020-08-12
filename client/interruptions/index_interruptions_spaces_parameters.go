// Code generated by go-swagger; DO NOT EDIT.

package interruptions

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

// NewIndexInterruptionsSpacesParams creates a new IndexInterruptionsSpacesParams object
// with the default values initialized.
func NewIndexInterruptionsSpacesParams() *IndexInterruptionsSpacesParams {
	var ()
	return &IndexInterruptionsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexInterruptionsSpacesParamsWithTimeout creates a new IndexInterruptionsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexInterruptionsSpacesParamsWithTimeout(timeout time.Duration) *IndexInterruptionsSpacesParams {
	var ()
	return &IndexInterruptionsSpacesParams{

		timeout: timeout,
	}
}

// NewIndexInterruptionsSpacesParamsWithContext creates a new IndexInterruptionsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexInterruptionsSpacesParamsWithContext(ctx context.Context) *IndexInterruptionsSpacesParams {
	var ()
	return &IndexInterruptionsSpacesParams{

		Context: ctx,
	}
}

// NewIndexInterruptionsSpacesParamsWithHTTPClient creates a new IndexInterruptionsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexInterruptionsSpacesParamsWithHTTPClient(client *http.Client) *IndexInterruptionsSpacesParams {
	var ()
	return &IndexInterruptionsSpacesParams{
		HTTPClient: client,
	}
}

/*IndexInterruptionsSpacesParams contains all the parameters to send to the API endpoint
for the index interruptions spaces operation typically these are written to a http.Request
*/
type IndexInterruptionsSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
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

// WithTimeout adds the timeout to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) WithTimeout(timeout time.Duration) *IndexInterruptionsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) WithContext(ctx context.Context) *IndexInterruptionsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) WithHTTPClient(client *http.Client) *IndexInterruptionsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) WithBaseSpaceID(baseSpaceID string) *IndexInterruptionsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithSkip adds the skip to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) WithSkip(skip *int32) *IndexInterruptionsSpacesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) WithTake(take *int32) *IndexInterruptionsSpacesParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index interruptions spaces params
func (o *IndexInterruptionsSpacesParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexInterruptionsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
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
