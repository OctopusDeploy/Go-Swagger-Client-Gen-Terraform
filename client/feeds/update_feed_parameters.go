// Code generated by go-swagger; DO NOT EDIT.

package feeds

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

	"models"
)

// NewUpdateFeedParams creates a new UpdateFeedParams object
// with the default values initialized.
func NewUpdateFeedParams() *UpdateFeedParams {
	var ()
	return &UpdateFeedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFeedParamsWithTimeout creates a new UpdateFeedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFeedParamsWithTimeout(timeout time.Duration) *UpdateFeedParams {
	var ()
	return &UpdateFeedParams{

		timeout: timeout,
	}
}

// NewUpdateFeedParamsWithContext creates a new UpdateFeedParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFeedParamsWithContext(ctx context.Context) *UpdateFeedParams {
	var ()
	return &UpdateFeedParams{

		Context: ctx,
	}
}

// NewUpdateFeedParamsWithHTTPClient creates a new UpdateFeedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFeedParamsWithHTTPClient(client *http.Client) *UpdateFeedParams {
	var ()
	return &UpdateFeedParams{
		HTTPClient: client,
	}
}

/*UpdateFeedParams contains all the parameters to send to the API endpoint
for the update feed operation typically these are written to a http.Request
*/
type UpdateFeedParams struct {

	/*FeedResource
	  The FeedResource resource to create

	*/
	FeedResource *models.FeedResource
	/*ID
	  ID of the FeedResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update feed params
func (o *UpdateFeedParams) WithTimeout(timeout time.Duration) *UpdateFeedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update feed params
func (o *UpdateFeedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update feed params
func (o *UpdateFeedParams) WithContext(ctx context.Context) *UpdateFeedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update feed params
func (o *UpdateFeedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update feed params
func (o *UpdateFeedParams) WithHTTPClient(client *http.Client) *UpdateFeedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update feed params
func (o *UpdateFeedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeedResource adds the feedResource to the update feed params
func (o *UpdateFeedParams) WithFeedResource(feedResource *models.FeedResource) *UpdateFeedParams {
	o.SetFeedResource(feedResource)
	return o
}

// SetFeedResource adds the feedResource to the update feed params
func (o *UpdateFeedParams) SetFeedResource(feedResource *models.FeedResource) {
	o.FeedResource = feedResource
}

// WithID adds the id to the update feed params
func (o *UpdateFeedParams) WithID(id string) *UpdateFeedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update feed params
func (o *UpdateFeedParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFeedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FeedResource != nil {
		if err := r.SetBodyParam(o.FeedResource); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
