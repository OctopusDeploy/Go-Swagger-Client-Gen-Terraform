// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

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

// NewIndexTagSetsSpacesParams creates a new IndexTagSetsSpacesParams object
// with the default values initialized.
func NewIndexTagSetsSpacesParams() *IndexTagSetsSpacesParams {
	var ()
	return &IndexTagSetsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexTagSetsSpacesParamsWithTimeout creates a new IndexTagSetsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexTagSetsSpacesParamsWithTimeout(timeout time.Duration) *IndexTagSetsSpacesParams {
	var ()
	return &IndexTagSetsSpacesParams{

		timeout: timeout,
	}
}

// NewIndexTagSetsSpacesParamsWithContext creates a new IndexTagSetsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexTagSetsSpacesParamsWithContext(ctx context.Context) *IndexTagSetsSpacesParams {
	var ()
	return &IndexTagSetsSpacesParams{

		Context: ctx,
	}
}

// NewIndexTagSetsSpacesParamsWithHTTPClient creates a new IndexTagSetsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexTagSetsSpacesParamsWithHTTPClient(client *http.Client) *IndexTagSetsSpacesParams {
	var ()
	return &IndexTagSetsSpacesParams{
		HTTPClient: client,
	}
}

/*IndexTagSetsSpacesParams contains all the parameters to send to the API endpoint
for the index tag sets spaces operation typically these are written to a http.Request
*/
type IndexTagSetsSpacesParams struct {

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

// WithTimeout adds the timeout to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) WithTimeout(timeout time.Duration) *IndexTagSetsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) WithContext(ctx context.Context) *IndexTagSetsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) WithHTTPClient(client *http.Client) *IndexTagSetsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) WithBaseSpaceID(baseSpaceID string) *IndexTagSetsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithSkip adds the skip to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) WithSkip(skip *int32) *IndexTagSetsSpacesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) WithTake(take *int32) *IndexTagSetsSpacesParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index tag sets spaces params
func (o *IndexTagSetsSpacesParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexTagSetsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
