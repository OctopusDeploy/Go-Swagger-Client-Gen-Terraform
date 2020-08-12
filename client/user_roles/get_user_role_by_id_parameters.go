// Code generated by go-swagger; DO NOT EDIT.

package user_roles

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

// NewGetUserRoleByIDParams creates a new GetUserRoleByIDParams object
// with the default values initialized.
func NewGetUserRoleByIDParams() *GetUserRoleByIDParams {
	var ()
	return &GetUserRoleByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserRoleByIDParamsWithTimeout creates a new GetUserRoleByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserRoleByIDParamsWithTimeout(timeout time.Duration) *GetUserRoleByIDParams {
	var ()
	return &GetUserRoleByIDParams{

		timeout: timeout,
	}
}

// NewGetUserRoleByIDParamsWithContext creates a new GetUserRoleByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserRoleByIDParamsWithContext(ctx context.Context) *GetUserRoleByIDParams {
	var ()
	return &GetUserRoleByIDParams{

		Context: ctx,
	}
}

// NewGetUserRoleByIDParamsWithHTTPClient creates a new GetUserRoleByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserRoleByIDParamsWithHTTPClient(client *http.Client) *GetUserRoleByIDParams {
	var ()
	return &GetUserRoleByIDParams{
		HTTPClient: client,
	}
}

/*GetUserRoleByIDParams contains all the parameters to send to the API endpoint
for the get user role by Id operation typically these are written to a http.Request
*/
type GetUserRoleByIDParams struct {

	/*ID
	  ID of the UserRoleResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user role by Id params
func (o *GetUserRoleByIDParams) WithTimeout(timeout time.Duration) *GetUserRoleByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user role by Id params
func (o *GetUserRoleByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user role by Id params
func (o *GetUserRoleByIDParams) WithContext(ctx context.Context) *GetUserRoleByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user role by Id params
func (o *GetUserRoleByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user role by Id params
func (o *GetUserRoleByIDParams) WithHTTPClient(client *http.Client) *GetUserRoleByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user role by Id params
func (o *GetUserRoleByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get user role by Id params
func (o *GetUserRoleByIDParams) WithID(id string) *GetUserRoleByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get user role by Id params
func (o *GetUserRoleByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserRoleByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
