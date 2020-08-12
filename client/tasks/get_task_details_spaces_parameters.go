// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewGetTaskDetailsSpacesParams creates a new GetTaskDetailsSpacesParams object
// with the default values initialized.
func NewGetTaskDetailsSpacesParams() *GetTaskDetailsSpacesParams {
	var ()
	return &GetTaskDetailsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskDetailsSpacesParamsWithTimeout creates a new GetTaskDetailsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTaskDetailsSpacesParamsWithTimeout(timeout time.Duration) *GetTaskDetailsSpacesParams {
	var ()
	return &GetTaskDetailsSpacesParams{

		timeout: timeout,
	}
}

// NewGetTaskDetailsSpacesParamsWithContext creates a new GetTaskDetailsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTaskDetailsSpacesParamsWithContext(ctx context.Context) *GetTaskDetailsSpacesParams {
	var ()
	return &GetTaskDetailsSpacesParams{

		Context: ctx,
	}
}

// NewGetTaskDetailsSpacesParamsWithHTTPClient creates a new GetTaskDetailsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTaskDetailsSpacesParamsWithHTTPClient(client *http.Client) *GetTaskDetailsSpacesParams {
	var ()
	return &GetTaskDetailsSpacesParams{
		HTTPClient: client,
	}
}

/*GetTaskDetailsSpacesParams contains all the parameters to send to the API endpoint
for the get task details spaces operation typically these are written to a http.Request
*/
type GetTaskDetailsSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get task details spaces params
func (o *GetTaskDetailsSpacesParams) WithTimeout(timeout time.Duration) *GetTaskDetailsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task details spaces params
func (o *GetTaskDetailsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task details spaces params
func (o *GetTaskDetailsSpacesParams) WithContext(ctx context.Context) *GetTaskDetailsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task details spaces params
func (o *GetTaskDetailsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task details spaces params
func (o *GetTaskDetailsSpacesParams) WithHTTPClient(client *http.Client) *GetTaskDetailsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task details spaces params
func (o *GetTaskDetailsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get task details spaces params
func (o *GetTaskDetailsSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetTaskDetailsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get task details spaces params
func (o *GetTaskDetailsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get task details spaces params
func (o *GetTaskDetailsSpacesParams) WithID(id string) *GetTaskDetailsSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get task details spaces params
func (o *GetTaskDetailsSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskDetailsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
