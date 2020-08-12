// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewUpdateAccountParams creates a new UpdateAccountParams object
// with the default values initialized.
func NewUpdateAccountParams() *UpdateAccountParams {
	var ()
	return &UpdateAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAccountParamsWithTimeout creates a new UpdateAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAccountParamsWithTimeout(timeout time.Duration) *UpdateAccountParams {
	var ()
	return &UpdateAccountParams{

		timeout: timeout,
	}
}

// NewUpdateAccountParamsWithContext creates a new UpdateAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAccountParamsWithContext(ctx context.Context) *UpdateAccountParams {
	var ()
	return &UpdateAccountParams{

		Context: ctx,
	}
}

// NewUpdateAccountParamsWithHTTPClient creates a new UpdateAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAccountParamsWithHTTPClient(client *http.Client) *UpdateAccountParams {
	var ()
	return &UpdateAccountParams{
		HTTPClient: client,
	}
}

/*UpdateAccountParams contains all the parameters to send to the API endpoint
for the update account operation typically these are written to a http.Request
*/
type UpdateAccountParams struct {

	/*AccountResource
	  The AccountResource resource to create

	*/
	AccountResource *models.AccountResource
	/*ID
	  ID of the AccountResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update account params
func (o *UpdateAccountParams) WithTimeout(timeout time.Duration) *UpdateAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update account params
func (o *UpdateAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update account params
func (o *UpdateAccountParams) WithContext(ctx context.Context) *UpdateAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update account params
func (o *UpdateAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update account params
func (o *UpdateAccountParams) WithHTTPClient(client *http.Client) *UpdateAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update account params
func (o *UpdateAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountResource adds the accountResource to the update account params
func (o *UpdateAccountParams) WithAccountResource(accountResource *models.AccountResource) *UpdateAccountParams {
	o.SetAccountResource(accountResource)
	return o
}

// SetAccountResource adds the accountResource to the update account params
func (o *UpdateAccountParams) SetAccountResource(accountResource *models.AccountResource) {
	o.AccountResource = accountResource
}

// WithID adds the id to the update account params
func (o *UpdateAccountParams) WithID(id string) *UpdateAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update account params
func (o *UpdateAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountResource != nil {
		if err := r.SetBodyParam(o.AccountResource); err != nil {
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