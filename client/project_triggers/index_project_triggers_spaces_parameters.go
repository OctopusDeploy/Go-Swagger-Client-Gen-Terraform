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

// NewIndexProjectTriggersSpacesParams creates a new IndexProjectTriggersSpacesParams object
// with the default values initialized.
func NewIndexProjectTriggersSpacesParams() *IndexProjectTriggersSpacesParams {
	var ()
	return &IndexProjectTriggersSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexProjectTriggersSpacesParamsWithTimeout creates a new IndexProjectTriggersSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexProjectTriggersSpacesParamsWithTimeout(timeout time.Duration) *IndexProjectTriggersSpacesParams {
	var ()
	return &IndexProjectTriggersSpacesParams{

		timeout: timeout,
	}
}

// NewIndexProjectTriggersSpacesParamsWithContext creates a new IndexProjectTriggersSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexProjectTriggersSpacesParamsWithContext(ctx context.Context) *IndexProjectTriggersSpacesParams {
	var ()
	return &IndexProjectTriggersSpacesParams{

		Context: ctx,
	}
}

// NewIndexProjectTriggersSpacesParamsWithHTTPClient creates a new IndexProjectTriggersSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexProjectTriggersSpacesParamsWithHTTPClient(client *http.Client) *IndexProjectTriggersSpacesParams {
	var ()
	return &IndexProjectTriggersSpacesParams{
		HTTPClient: client,
	}
}

/*IndexProjectTriggersSpacesParams contains all the parameters to send to the API endpoint
for the index project triggers spaces operation typically these are written to a http.Request
*/
type IndexProjectTriggersSpacesParams struct {

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

// WithTimeout adds the timeout to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) WithTimeout(timeout time.Duration) *IndexProjectTriggersSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) WithContext(ctx context.Context) *IndexProjectTriggersSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) WithHTTPClient(client *http.Client) *IndexProjectTriggersSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) WithBaseSpaceID(baseSpaceID string) *IndexProjectTriggersSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithSkip adds the skip to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) WithSkip(skip *int32) *IndexProjectTriggersSpacesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) WithTake(take *int32) *IndexProjectTriggersSpacesParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index project triggers spaces params
func (o *IndexProjectTriggersSpacesParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexProjectTriggersSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
