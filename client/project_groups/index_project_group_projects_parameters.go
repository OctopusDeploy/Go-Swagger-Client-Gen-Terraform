// Code generated by go-swagger; DO NOT EDIT.

package project_groups

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

// NewIndexProjectGroupProjectsParams creates a new IndexProjectGroupProjectsParams object
// with the default values initialized.
func NewIndexProjectGroupProjectsParams() *IndexProjectGroupProjectsParams {
	var ()
	return &IndexProjectGroupProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexProjectGroupProjectsParamsWithTimeout creates a new IndexProjectGroupProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexProjectGroupProjectsParamsWithTimeout(timeout time.Duration) *IndexProjectGroupProjectsParams {
	var ()
	return &IndexProjectGroupProjectsParams{

		timeout: timeout,
	}
}

// NewIndexProjectGroupProjectsParamsWithContext creates a new IndexProjectGroupProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexProjectGroupProjectsParamsWithContext(ctx context.Context) *IndexProjectGroupProjectsParams {
	var ()
	return &IndexProjectGroupProjectsParams{

		Context: ctx,
	}
}

// NewIndexProjectGroupProjectsParamsWithHTTPClient creates a new IndexProjectGroupProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexProjectGroupProjectsParamsWithHTTPClient(client *http.Client) *IndexProjectGroupProjectsParams {
	var ()
	return &IndexProjectGroupProjectsParams{
		HTTPClient: client,
	}
}

/*IndexProjectGroupProjectsParams contains all the parameters to send to the API endpoint
for the index project group projects operation typically these are written to a http.Request
*/
type IndexProjectGroupProjectsParams struct {

	/*ID
	  ID of the ProjectGroup

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

// WithTimeout adds the timeout to the index project group projects params
func (o *IndexProjectGroupProjectsParams) WithTimeout(timeout time.Duration) *IndexProjectGroupProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index project group projects params
func (o *IndexProjectGroupProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index project group projects params
func (o *IndexProjectGroupProjectsParams) WithContext(ctx context.Context) *IndexProjectGroupProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index project group projects params
func (o *IndexProjectGroupProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index project group projects params
func (o *IndexProjectGroupProjectsParams) WithHTTPClient(client *http.Client) *IndexProjectGroupProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index project group projects params
func (o *IndexProjectGroupProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the index project group projects params
func (o *IndexProjectGroupProjectsParams) WithID(id string) *IndexProjectGroupProjectsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the index project group projects params
func (o *IndexProjectGroupProjectsParams) SetID(id string) {
	o.ID = id
}

// WithSkip adds the skip to the index project group projects params
func (o *IndexProjectGroupProjectsParams) WithSkip(skip *int32) *IndexProjectGroupProjectsParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index project group projects params
func (o *IndexProjectGroupProjectsParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index project group projects params
func (o *IndexProjectGroupProjectsParams) WithTake(take *int32) *IndexProjectGroupProjectsParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index project group projects params
func (o *IndexProjectGroupProjectsParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexProjectGroupProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
