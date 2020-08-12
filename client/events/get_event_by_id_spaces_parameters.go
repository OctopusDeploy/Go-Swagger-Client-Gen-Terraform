// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewGetEventByIDSpacesParams creates a new GetEventByIDSpacesParams object
// with the default values initialized.
func NewGetEventByIDSpacesParams() *GetEventByIDSpacesParams {
	var ()
	return &GetEventByIDSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventByIDSpacesParamsWithTimeout creates a new GetEventByIDSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventByIDSpacesParamsWithTimeout(timeout time.Duration) *GetEventByIDSpacesParams {
	var ()
	return &GetEventByIDSpacesParams{

		timeout: timeout,
	}
}

// NewGetEventByIDSpacesParamsWithContext creates a new GetEventByIDSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventByIDSpacesParamsWithContext(ctx context.Context) *GetEventByIDSpacesParams {
	var ()
	return &GetEventByIDSpacesParams{

		Context: ctx,
	}
}

// NewGetEventByIDSpacesParamsWithHTTPClient creates a new GetEventByIDSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventByIDSpacesParamsWithHTTPClient(client *http.Client) *GetEventByIDSpacesParams {
	var ()
	return &GetEventByIDSpacesParams{
		HTTPClient: client,
	}
}

/*GetEventByIDSpacesParams contains all the parameters to send to the API endpoint
for the get event by Id spaces operation typically these are written to a http.Request
*/
type GetEventByIDSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the EventResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get event by Id spaces params
func (o *GetEventByIDSpacesParams) WithTimeout(timeout time.Duration) *GetEventByIDSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get event by Id spaces params
func (o *GetEventByIDSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get event by Id spaces params
func (o *GetEventByIDSpacesParams) WithContext(ctx context.Context) *GetEventByIDSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get event by Id spaces params
func (o *GetEventByIDSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get event by Id spaces params
func (o *GetEventByIDSpacesParams) WithHTTPClient(client *http.Client) *GetEventByIDSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get event by Id spaces params
func (o *GetEventByIDSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get event by Id spaces params
func (o *GetEventByIDSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetEventByIDSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get event by Id spaces params
func (o *GetEventByIDSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get event by Id spaces params
func (o *GetEventByIDSpacesParams) WithID(id string) *GetEventByIDSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get event by Id spaces params
func (o *GetEventByIDSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventByIDSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
