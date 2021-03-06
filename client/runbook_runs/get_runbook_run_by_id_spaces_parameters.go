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

// NewGetRunbookRunByIDSpacesParams creates a new GetRunbookRunByIDSpacesParams object
// with the default values initialized.
func NewGetRunbookRunByIDSpacesParams() *GetRunbookRunByIDSpacesParams {
	var ()
	return &GetRunbookRunByIDSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunbookRunByIDSpacesParamsWithTimeout creates a new GetRunbookRunByIDSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunbookRunByIDSpacesParamsWithTimeout(timeout time.Duration) *GetRunbookRunByIDSpacesParams {
	var ()
	return &GetRunbookRunByIDSpacesParams{

		timeout: timeout,
	}
}

// NewGetRunbookRunByIDSpacesParamsWithContext creates a new GetRunbookRunByIDSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunbookRunByIDSpacesParamsWithContext(ctx context.Context) *GetRunbookRunByIDSpacesParams {
	var ()
	return &GetRunbookRunByIDSpacesParams{

		Context: ctx,
	}
}

// NewGetRunbookRunByIDSpacesParamsWithHTTPClient creates a new GetRunbookRunByIDSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunbookRunByIDSpacesParamsWithHTTPClient(client *http.Client) *GetRunbookRunByIDSpacesParams {
	var ()
	return &GetRunbookRunByIDSpacesParams{
		HTTPClient: client,
	}
}

/*GetRunbookRunByIDSpacesParams contains all the parameters to send to the API endpoint
for the get runbook run by Id spaces operation typically these are written to a http.Request
*/
type GetRunbookRunByIDSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the RunbookRunResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runbook run by Id spaces params
func (o *GetRunbookRunByIDSpacesParams) WithTimeout(timeout time.Duration) *GetRunbookRunByIDSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runbook run by Id spaces params
func (o *GetRunbookRunByIDSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runbook run by Id spaces params
func (o *GetRunbookRunByIDSpacesParams) WithContext(ctx context.Context) *GetRunbookRunByIDSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runbook run by Id spaces params
func (o *GetRunbookRunByIDSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runbook run by Id spaces params
func (o *GetRunbookRunByIDSpacesParams) WithHTTPClient(client *http.Client) *GetRunbookRunByIDSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runbook run by Id spaces params
func (o *GetRunbookRunByIDSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get runbook run by Id spaces params
func (o *GetRunbookRunByIDSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetRunbookRunByIDSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get runbook run by Id spaces params
func (o *GetRunbookRunByIDSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get runbook run by Id spaces params
func (o *GetRunbookRunByIDSpacesParams) WithID(id string) *GetRunbookRunByIDSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get runbook run by Id spaces params
func (o *GetRunbookRunByIDSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunbookRunByIDSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
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
