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

	"models"
)

// NewCreateScopedUserRoleParams creates a new CreateScopedUserRoleParams object
// with the default values initialized.
func NewCreateScopedUserRoleParams() *CreateScopedUserRoleParams {
	var ()
	return &CreateScopedUserRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateScopedUserRoleParamsWithTimeout creates a new CreateScopedUserRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateScopedUserRoleParamsWithTimeout(timeout time.Duration) *CreateScopedUserRoleParams {
	var ()
	return &CreateScopedUserRoleParams{

		timeout: timeout,
	}
}

// NewCreateScopedUserRoleParamsWithContext creates a new CreateScopedUserRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateScopedUserRoleParamsWithContext(ctx context.Context) *CreateScopedUserRoleParams {
	var ()
	return &CreateScopedUserRoleParams{

		Context: ctx,
	}
}

// NewCreateScopedUserRoleParamsWithHTTPClient creates a new CreateScopedUserRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateScopedUserRoleParamsWithHTTPClient(client *http.Client) *CreateScopedUserRoleParams {
	var ()
	return &CreateScopedUserRoleParams{
		HTTPClient: client,
	}
}

/*CreateScopedUserRoleParams contains all the parameters to send to the API endpoint
for the create scoped user role operation typically these are written to a http.Request
*/
type CreateScopedUserRoleParams struct {

	/*ScopedUserRoleResource
	  The ScopedUserRoleResource resource to create

	*/
	ScopedUserRoleResource *models.ScopedUserRoleResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create scoped user role params
func (o *CreateScopedUserRoleParams) WithTimeout(timeout time.Duration) *CreateScopedUserRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create scoped user role params
func (o *CreateScopedUserRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create scoped user role params
func (o *CreateScopedUserRoleParams) WithContext(ctx context.Context) *CreateScopedUserRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create scoped user role params
func (o *CreateScopedUserRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create scoped user role params
func (o *CreateScopedUserRoleParams) WithHTTPClient(client *http.Client) *CreateScopedUserRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create scoped user role params
func (o *CreateScopedUserRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScopedUserRoleResource adds the scopedUserRoleResource to the create scoped user role params
func (o *CreateScopedUserRoleParams) WithScopedUserRoleResource(scopedUserRoleResource *models.ScopedUserRoleResource) *CreateScopedUserRoleParams {
	o.SetScopedUserRoleResource(scopedUserRoleResource)
	return o
}

// SetScopedUserRoleResource adds the scopedUserRoleResource to the create scoped user role params
func (o *CreateScopedUserRoleParams) SetScopedUserRoleResource(scopedUserRoleResource *models.ScopedUserRoleResource) {
	o.ScopedUserRoleResource = scopedUserRoleResource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateScopedUserRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ScopedUserRoleResource != nil {
		if err := r.SetBodyParam(o.ScopedUserRoleResource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
