// Code generated by go-swagger; DO NOT EDIT.

package scoped_user_role

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

// NewGetScopedUserRoleByIDParams creates a new GetScopedUserRoleByIDParams object
// with the default values initialized.
func NewGetScopedUserRoleByIDParams() *GetScopedUserRoleByIDParams {
	var ()
	return &GetScopedUserRoleByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScopedUserRoleByIDParamsWithTimeout creates a new GetScopedUserRoleByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScopedUserRoleByIDParamsWithTimeout(timeout time.Duration) *GetScopedUserRoleByIDParams {
	var ()
	return &GetScopedUserRoleByIDParams{

		timeout: timeout,
	}
}

// NewGetScopedUserRoleByIDParamsWithContext creates a new GetScopedUserRoleByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScopedUserRoleByIDParamsWithContext(ctx context.Context) *GetScopedUserRoleByIDParams {
	var ()
	return &GetScopedUserRoleByIDParams{

		Context: ctx,
	}
}

// NewGetScopedUserRoleByIDParamsWithHTTPClient creates a new GetScopedUserRoleByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScopedUserRoleByIDParamsWithHTTPClient(client *http.Client) *GetScopedUserRoleByIDParams {
	var ()
	return &GetScopedUserRoleByIDParams{
		HTTPClient: client,
	}
}

/*GetScopedUserRoleByIDParams contains all the parameters to send to the API endpoint
for the get scoped user role by Id operation typically these are written to a http.Request
*/
type GetScopedUserRoleByIDParams struct {

	/*ID
	  ID of the ScopedUserRoleResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scoped user role by Id params
func (o *GetScopedUserRoleByIDParams) WithTimeout(timeout time.Duration) *GetScopedUserRoleByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scoped user role by Id params
func (o *GetScopedUserRoleByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scoped user role by Id params
func (o *GetScopedUserRoleByIDParams) WithContext(ctx context.Context) *GetScopedUserRoleByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scoped user role by Id params
func (o *GetScopedUserRoleByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scoped user role by Id params
func (o *GetScopedUserRoleByIDParams) WithHTTPClient(client *http.Client) *GetScopedUserRoleByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scoped user role by Id params
func (o *GetScopedUserRoleByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get scoped user role by Id params
func (o *GetScopedUserRoleByIDParams) WithID(id string) *GetScopedUserRoleByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get scoped user role by Id params
func (o *GetScopedUserRoleByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetScopedUserRoleByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
