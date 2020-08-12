// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewUpdateTenantParams creates a new UpdateTenantParams object
// with the default values initialized.
func NewUpdateTenantParams() *UpdateTenantParams {
	var ()
	return &UpdateTenantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTenantParamsWithTimeout creates a new UpdateTenantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTenantParamsWithTimeout(timeout time.Duration) *UpdateTenantParams {
	var ()
	return &UpdateTenantParams{

		timeout: timeout,
	}
}

// NewUpdateTenantParamsWithContext creates a new UpdateTenantParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTenantParamsWithContext(ctx context.Context) *UpdateTenantParams {
	var ()
	return &UpdateTenantParams{

		Context: ctx,
	}
}

// NewUpdateTenantParamsWithHTTPClient creates a new UpdateTenantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTenantParamsWithHTTPClient(client *http.Client) *UpdateTenantParams {
	var ()
	return &UpdateTenantParams{
		HTTPClient: client,
	}
}

/*UpdateTenantParams contains all the parameters to send to the API endpoint
for the update tenant operation typically these are written to a http.Request
*/
type UpdateTenantParams struct {

	/*TenantResource
	  The TenantResource resource to create

	*/
	TenantResource *models.TenantResource
	/*ID
	  ID of the TenantResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update tenant params
func (o *UpdateTenantParams) WithTimeout(timeout time.Duration) *UpdateTenantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tenant params
func (o *UpdateTenantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tenant params
func (o *UpdateTenantParams) WithContext(ctx context.Context) *UpdateTenantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tenant params
func (o *UpdateTenantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tenant params
func (o *UpdateTenantParams) WithHTTPClient(client *http.Client) *UpdateTenantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tenant params
func (o *UpdateTenantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantResource adds the tenantResource to the update tenant params
func (o *UpdateTenantParams) WithTenantResource(tenantResource *models.TenantResource) *UpdateTenantParams {
	o.SetTenantResource(tenantResource)
	return o
}

// SetTenantResource adds the tenantResource to the update tenant params
func (o *UpdateTenantParams) SetTenantResource(tenantResource *models.TenantResource) {
	o.TenantResource = tenantResource
}

// WithID adds the id to the update tenant params
func (o *UpdateTenantParams) WithID(id string) *UpdateTenantParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update tenant params
func (o *UpdateTenantParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTenantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TenantResource != nil {
		if err := r.SetBodyParam(o.TenantResource); err != nil {
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
