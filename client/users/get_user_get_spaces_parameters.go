// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetUserGetSpacesParams creates a new GetUserGetSpacesParams object
// with the default values initialized.
func NewGetUserGetSpacesParams() *GetUserGetSpacesParams {
	var ()
	return &GetUserGetSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserGetSpacesParamsWithTimeout creates a new GetUserGetSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserGetSpacesParamsWithTimeout(timeout time.Duration) *GetUserGetSpacesParams {
	var ()
	return &GetUserGetSpacesParams{

		timeout: timeout,
	}
}

// NewGetUserGetSpacesParamsWithContext creates a new GetUserGetSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserGetSpacesParamsWithContext(ctx context.Context) *GetUserGetSpacesParams {
	var ()
	return &GetUserGetSpacesParams{

		Context: ctx,
	}
}

// NewGetUserGetSpacesParamsWithHTTPClient creates a new GetUserGetSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserGetSpacesParamsWithHTTPClient(client *http.Client) *GetUserGetSpacesParams {
	var ()
	return &GetUserGetSpacesParams{
		HTTPClient: client,
	}
}

/*GetUserGetSpacesParams contains all the parameters to send to the API endpoint
for the get user get spaces operation typically these are written to a http.Request
*/
type GetUserGetSpacesParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user get spaces params
func (o *GetUserGetSpacesParams) WithTimeout(timeout time.Duration) *GetUserGetSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user get spaces params
func (o *GetUserGetSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user get spaces params
func (o *GetUserGetSpacesParams) WithContext(ctx context.Context) *GetUserGetSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user get spaces params
func (o *GetUserGetSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user get spaces params
func (o *GetUserGetSpacesParams) WithHTTPClient(client *http.Client) *GetUserGetSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user get spaces params
func (o *GetUserGetSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get user get spaces params
func (o *GetUserGetSpacesParams) WithID(id string) *GetUserGetSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get user get spaces params
func (o *GetUserGetSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserGetSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
