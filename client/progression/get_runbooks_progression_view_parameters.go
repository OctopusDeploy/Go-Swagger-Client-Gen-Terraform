// Code generated by go-swagger; DO NOT EDIT.

package progression

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

// NewGetRunbooksProgressionViewParams creates a new GetRunbooksProgressionViewParams object
// with the default values initialized.
func NewGetRunbooksProgressionViewParams() *GetRunbooksProgressionViewParams {
	var ()
	return &GetRunbooksProgressionViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunbooksProgressionViewParamsWithTimeout creates a new GetRunbooksProgressionViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunbooksProgressionViewParamsWithTimeout(timeout time.Duration) *GetRunbooksProgressionViewParams {
	var ()
	return &GetRunbooksProgressionViewParams{

		timeout: timeout,
	}
}

// NewGetRunbooksProgressionViewParamsWithContext creates a new GetRunbooksProgressionViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunbooksProgressionViewParamsWithContext(ctx context.Context) *GetRunbooksProgressionViewParams {
	var ()
	return &GetRunbooksProgressionViewParams{

		Context: ctx,
	}
}

// NewGetRunbooksProgressionViewParamsWithHTTPClient creates a new GetRunbooksProgressionViewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunbooksProgressionViewParamsWithHTTPClient(client *http.Client) *GetRunbooksProgressionViewParams {
	var ()
	return &GetRunbooksProgressionViewParams{
		HTTPClient: client,
	}
}

/*GetRunbooksProgressionViewParams contains all the parameters to send to the API endpoint
for the get runbooks progression view operation typically these are written to a http.Request
*/
type GetRunbooksProgressionViewParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runbooks progression view params
func (o *GetRunbooksProgressionViewParams) WithTimeout(timeout time.Duration) *GetRunbooksProgressionViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runbooks progression view params
func (o *GetRunbooksProgressionViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runbooks progression view params
func (o *GetRunbooksProgressionViewParams) WithContext(ctx context.Context) *GetRunbooksProgressionViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runbooks progression view params
func (o *GetRunbooksProgressionViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runbooks progression view params
func (o *GetRunbooksProgressionViewParams) WithHTTPClient(client *http.Client) *GetRunbooksProgressionViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runbooks progression view params
func (o *GetRunbooksProgressionViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get runbooks progression view params
func (o *GetRunbooksProgressionViewParams) WithID(id string) *GetRunbooksProgressionViewParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get runbooks progression view params
func (o *GetRunbooksProgressionViewParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunbooksProgressionViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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