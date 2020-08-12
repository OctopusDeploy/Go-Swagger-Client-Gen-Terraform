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
)

// NewDeleteFeedParams creates a new DeleteFeedParams object
// with the default values initialized.
func NewDeleteFeedParams() *DeleteFeedParams {
	var ()
	return &DeleteFeedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFeedParamsWithTimeout creates a new DeleteFeedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFeedParamsWithTimeout(timeout time.Duration) *DeleteFeedParams {
	var ()
	return &DeleteFeedParams{

		timeout: timeout,
	}
}

// NewDeleteFeedParamsWithContext creates a new DeleteFeedParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFeedParamsWithContext(ctx context.Context) *DeleteFeedParams {
	var ()
	return &DeleteFeedParams{

		Context: ctx,
	}
}

// NewDeleteFeedParamsWithHTTPClient creates a new DeleteFeedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFeedParamsWithHTTPClient(client *http.Client) *DeleteFeedParams {
	var ()
	return &DeleteFeedParams{
		HTTPClient: client,
	}
}

/*DeleteFeedParams contains all the parameters to send to the API endpoint
for the delete feed operation typically these are written to a http.Request
*/
type DeleteFeedParams struct {

	/*ID
	  ID of the FeedResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete feed params
func (o *DeleteFeedParams) WithTimeout(timeout time.Duration) *DeleteFeedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete feed params
func (o *DeleteFeedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete feed params
func (o *DeleteFeedParams) WithContext(ctx context.Context) *DeleteFeedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete feed params
func (o *DeleteFeedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete feed params
func (o *DeleteFeedParams) WithHTTPClient(client *http.Client) *DeleteFeedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete feed params
func (o *DeleteFeedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete feed params
func (o *DeleteFeedParams) WithID(id string) *DeleteFeedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete feed params
func (o *DeleteFeedParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFeedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}