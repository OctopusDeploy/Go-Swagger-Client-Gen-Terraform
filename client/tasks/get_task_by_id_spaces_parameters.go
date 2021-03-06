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

// NewGetTaskByIDSpacesParams creates a new GetTaskByIDSpacesParams object
// with the default values initialized.
func NewGetTaskByIDSpacesParams() *GetTaskByIDSpacesParams {
	var ()
	return &GetTaskByIDSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskByIDSpacesParamsWithTimeout creates a new GetTaskByIDSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTaskByIDSpacesParamsWithTimeout(timeout time.Duration) *GetTaskByIDSpacesParams {
	var ()
	return &GetTaskByIDSpacesParams{

		timeout: timeout,
	}
}

// NewGetTaskByIDSpacesParamsWithContext creates a new GetTaskByIDSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTaskByIDSpacesParamsWithContext(ctx context.Context) *GetTaskByIDSpacesParams {
	var ()
	return &GetTaskByIDSpacesParams{

		Context: ctx,
	}
}

// NewGetTaskByIDSpacesParamsWithHTTPClient creates a new GetTaskByIDSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTaskByIDSpacesParamsWithHTTPClient(client *http.Client) *GetTaskByIDSpacesParams {
	var ()
	return &GetTaskByIDSpacesParams{
		HTTPClient: client,
	}
}

/*GetTaskByIDSpacesParams contains all the parameters to send to the API endpoint
for the get task by Id spaces operation typically these are written to a http.Request
*/
type GetTaskByIDSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the TaskResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get task by Id spaces params
func (o *GetTaskByIDSpacesParams) WithTimeout(timeout time.Duration) *GetTaskByIDSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task by Id spaces params
func (o *GetTaskByIDSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task by Id spaces params
func (o *GetTaskByIDSpacesParams) WithContext(ctx context.Context) *GetTaskByIDSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task by Id spaces params
func (o *GetTaskByIDSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task by Id spaces params
func (o *GetTaskByIDSpacesParams) WithHTTPClient(client *http.Client) *GetTaskByIDSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task by Id spaces params
func (o *GetTaskByIDSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get task by Id spaces params
func (o *GetTaskByIDSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetTaskByIDSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get task by Id spaces params
func (o *GetTaskByIDSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get task by Id spaces params
func (o *GetTaskByIDSpacesParams) WithID(id string) *GetTaskByIDSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get task by Id spaces params
func (o *GetTaskByIDSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskByIDSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
