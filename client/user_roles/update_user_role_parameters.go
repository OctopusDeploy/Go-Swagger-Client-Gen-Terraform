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

	"models"
)

// NewUpdateUserRoleParams creates a new UpdateUserRoleParams object
// with the default values initialized.
func NewUpdateUserRoleParams() *UpdateUserRoleParams {
	var ()
	return &UpdateUserRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserRoleParamsWithTimeout creates a new UpdateUserRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserRoleParamsWithTimeout(timeout time.Duration) *UpdateUserRoleParams {
	var ()
	return &UpdateUserRoleParams{

		timeout: timeout,
	}
}

// NewUpdateUserRoleParamsWithContext creates a new UpdateUserRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserRoleParamsWithContext(ctx context.Context) *UpdateUserRoleParams {
	var ()
	return &UpdateUserRoleParams{

		Context: ctx,
	}
}

// NewUpdateUserRoleParamsWithHTTPClient creates a new UpdateUserRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserRoleParamsWithHTTPClient(client *http.Client) *UpdateUserRoleParams {
	var ()
	return &UpdateUserRoleParams{
		HTTPClient: client,
	}
}

/*UpdateUserRoleParams contains all the parameters to send to the API endpoint
for the update user role operation typically these are written to a http.Request
*/
type UpdateUserRoleParams struct {

	/*UserRoleResource
	  The UserRoleResource resource to create

	*/
	UserRoleResource *models.UserRoleResource
	/*ID
	  ID of the UserRoleResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user role params
func (o *UpdateUserRoleParams) WithTimeout(timeout time.Duration) *UpdateUserRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user role params
func (o *UpdateUserRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user role params
func (o *UpdateUserRoleParams) WithContext(ctx context.Context) *UpdateUserRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user role params
func (o *UpdateUserRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user role params
func (o *UpdateUserRoleParams) WithHTTPClient(client *http.Client) *UpdateUserRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user role params
func (o *UpdateUserRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserRoleResource adds the userRoleResource to the update user role params
func (o *UpdateUserRoleParams) WithUserRoleResource(userRoleResource *models.UserRoleResource) *UpdateUserRoleParams {
	o.SetUserRoleResource(userRoleResource)
	return o
}

// SetUserRoleResource adds the userRoleResource to the update user role params
func (o *UpdateUserRoleParams) SetUserRoleResource(userRoleResource *models.UserRoleResource) {
	o.UserRoleResource = userRoleResource
}

// WithID adds the id to the update user role params
func (o *UpdateUserRoleParams) WithID(id string) *UpdateUserRoleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update user role params
func (o *UpdateUserRoleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserRoleResource != nil {
		if err := r.SetBodyParam(o.UserRoleResource); err != nil {
			return err
		}
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
