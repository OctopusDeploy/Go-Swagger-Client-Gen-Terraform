// Code generated by go-swagger; DO NOT EDIT.

package runbook_processes

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

// NewGetRunbookProcessByIDSpacesParams creates a new GetRunbookProcessByIDSpacesParams object
// with the default values initialized.
func NewGetRunbookProcessByIDSpacesParams() *GetRunbookProcessByIDSpacesParams {
	var ()
	return &GetRunbookProcessByIDSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunbookProcessByIDSpacesParamsWithTimeout creates a new GetRunbookProcessByIDSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunbookProcessByIDSpacesParamsWithTimeout(timeout time.Duration) *GetRunbookProcessByIDSpacesParams {
	var ()
	return &GetRunbookProcessByIDSpacesParams{

		timeout: timeout,
	}
}

// NewGetRunbookProcessByIDSpacesParamsWithContext creates a new GetRunbookProcessByIDSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunbookProcessByIDSpacesParamsWithContext(ctx context.Context) *GetRunbookProcessByIDSpacesParams {
	var ()
	return &GetRunbookProcessByIDSpacesParams{

		Context: ctx,
	}
}

// NewGetRunbookProcessByIDSpacesParamsWithHTTPClient creates a new GetRunbookProcessByIDSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunbookProcessByIDSpacesParamsWithHTTPClient(client *http.Client) *GetRunbookProcessByIDSpacesParams {
	var ()
	return &GetRunbookProcessByIDSpacesParams{
		HTTPClient: client,
	}
}

/*GetRunbookProcessByIDSpacesParams contains all the parameters to send to the API endpoint
for the get runbook process by Id spaces operation typically these are written to a http.Request
*/
type GetRunbookProcessByIDSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the RunbookProcessResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runbook process by Id spaces params
func (o *GetRunbookProcessByIDSpacesParams) WithTimeout(timeout time.Duration) *GetRunbookProcessByIDSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runbook process by Id spaces params
func (o *GetRunbookProcessByIDSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runbook process by Id spaces params
func (o *GetRunbookProcessByIDSpacesParams) WithContext(ctx context.Context) *GetRunbookProcessByIDSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runbook process by Id spaces params
func (o *GetRunbookProcessByIDSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runbook process by Id spaces params
func (o *GetRunbookProcessByIDSpacesParams) WithHTTPClient(client *http.Client) *GetRunbookProcessByIDSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runbook process by Id spaces params
func (o *GetRunbookProcessByIDSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get runbook process by Id spaces params
func (o *GetRunbookProcessByIDSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetRunbookProcessByIDSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get runbook process by Id spaces params
func (o *GetRunbookProcessByIDSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get runbook process by Id spaces params
func (o *GetRunbookProcessByIDSpacesParams) WithID(id string) *GetRunbookProcessByIDSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get runbook process by Id spaces params
func (o *GetRunbookProcessByIDSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunbookProcessByIDSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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