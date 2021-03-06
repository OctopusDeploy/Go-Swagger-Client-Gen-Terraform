// Code generated by go-swagger; DO NOT EDIT.

package api_keys

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

// NewIndexAPIKeysParams creates a new IndexAPIKeysParams object
// with the default values initialized.
func NewIndexAPIKeysParams() *IndexAPIKeysParams {
	var ()
	return &IndexAPIKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexAPIKeysParamsWithTimeout creates a new IndexAPIKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexAPIKeysParamsWithTimeout(timeout time.Duration) *IndexAPIKeysParams {
	var ()
	return &IndexAPIKeysParams{

		timeout: timeout,
	}
}

// NewIndexAPIKeysParamsWithContext creates a new IndexAPIKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexAPIKeysParamsWithContext(ctx context.Context) *IndexAPIKeysParams {
	var ()
	return &IndexAPIKeysParams{

		Context: ctx,
	}
}

// NewIndexAPIKeysParamsWithHTTPClient creates a new IndexAPIKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexAPIKeysParamsWithHTTPClient(client *http.Client) *IndexAPIKeysParams {
	var ()
	return &IndexAPIKeysParams{
		HTTPClient: client,
	}
}

/*IndexAPIKeysParams contains all the parameters to send to the API endpoint
for the index Api keys operation typically these are written to a http.Request
*/
type IndexAPIKeysParams struct {

	/*Skip
	  Number of items to skip

	*/
	Skip *int32
	/*Take
	  Number of items to take

	*/
	Take *int32
	/*UserID
	  ID of the user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the index Api keys params
func (o *IndexAPIKeysParams) WithTimeout(timeout time.Duration) *IndexAPIKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index Api keys params
func (o *IndexAPIKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index Api keys params
func (o *IndexAPIKeysParams) WithContext(ctx context.Context) *IndexAPIKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index Api keys params
func (o *IndexAPIKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index Api keys params
func (o *IndexAPIKeysParams) WithHTTPClient(client *http.Client) *IndexAPIKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index Api keys params
func (o *IndexAPIKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkip adds the skip to the index Api keys params
func (o *IndexAPIKeysParams) WithSkip(skip *int32) *IndexAPIKeysParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index Api keys params
func (o *IndexAPIKeysParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index Api keys params
func (o *IndexAPIKeysParams) WithTake(take *int32) *IndexAPIKeysParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index Api keys params
func (o *IndexAPIKeysParams) SetTake(take *int32) {
	o.Take = take
}

// WithUserID adds the userID to the index Api keys params
func (o *IndexAPIKeysParams) WithUserID(userID string) *IndexAPIKeysParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the index Api keys params
func (o *IndexAPIKeysParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *IndexAPIKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
