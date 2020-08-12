// Code generated by go-swagger; DO NOT EDIT.

package runbook_runs

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

// NewGetRunbookRunByIDParams creates a new GetRunbookRunByIDParams object
// with the default values initialized.
func NewGetRunbookRunByIDParams() *GetRunbookRunByIDParams {
	var ()
	return &GetRunbookRunByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunbookRunByIDParamsWithTimeout creates a new GetRunbookRunByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunbookRunByIDParamsWithTimeout(timeout time.Duration) *GetRunbookRunByIDParams {
	var ()
	return &GetRunbookRunByIDParams{

		timeout: timeout,
	}
}

// NewGetRunbookRunByIDParamsWithContext creates a new GetRunbookRunByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunbookRunByIDParamsWithContext(ctx context.Context) *GetRunbookRunByIDParams {
	var ()
	return &GetRunbookRunByIDParams{

		Context: ctx,
	}
}

// NewGetRunbookRunByIDParamsWithHTTPClient creates a new GetRunbookRunByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunbookRunByIDParamsWithHTTPClient(client *http.Client) *GetRunbookRunByIDParams {
	var ()
	return &GetRunbookRunByIDParams{
		HTTPClient: client,
	}
}

/*GetRunbookRunByIDParams contains all the parameters to send to the API endpoint
for the get runbook run by Id operation typically these are written to a http.Request
*/
type GetRunbookRunByIDParams struct {

	/*ID
	  ID of the RunbookRunResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runbook run by Id params
func (o *GetRunbookRunByIDParams) WithTimeout(timeout time.Duration) *GetRunbookRunByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runbook run by Id params
func (o *GetRunbookRunByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runbook run by Id params
func (o *GetRunbookRunByIDParams) WithContext(ctx context.Context) *GetRunbookRunByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runbook run by Id params
func (o *GetRunbookRunByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runbook run by Id params
func (o *GetRunbookRunByIDParams) WithHTTPClient(client *http.Client) *GetRunbookRunByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runbook run by Id params
func (o *GetRunbookRunByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get runbook run by Id params
func (o *GetRunbookRunByIDParams) WithID(id string) *GetRunbookRunByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get runbook run by Id params
func (o *GetRunbookRunByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunbookRunByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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