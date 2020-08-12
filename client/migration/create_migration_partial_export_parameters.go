// Code generated by go-swagger; DO NOT EDIT.

package migration

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

// NewCreateMigrationPartialExportParams creates a new CreateMigrationPartialExportParams object
// with the default values initialized.
func NewCreateMigrationPartialExportParams() *CreateMigrationPartialExportParams {

	return &CreateMigrationPartialExportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMigrationPartialExportParamsWithTimeout creates a new CreateMigrationPartialExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMigrationPartialExportParamsWithTimeout(timeout time.Duration) *CreateMigrationPartialExportParams {

	return &CreateMigrationPartialExportParams{

		timeout: timeout,
	}
}

// NewCreateMigrationPartialExportParamsWithContext creates a new CreateMigrationPartialExportParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMigrationPartialExportParamsWithContext(ctx context.Context) *CreateMigrationPartialExportParams {

	return &CreateMigrationPartialExportParams{

		Context: ctx,
	}
}

// NewCreateMigrationPartialExportParamsWithHTTPClient creates a new CreateMigrationPartialExportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMigrationPartialExportParamsWithHTTPClient(client *http.Client) *CreateMigrationPartialExportParams {

	return &CreateMigrationPartialExportParams{
		HTTPClient: client,
	}
}

/*CreateMigrationPartialExportParams contains all the parameters to send to the API endpoint
for the create migration partial export operation typically these are written to a http.Request
*/
type CreateMigrationPartialExportParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create migration partial export params
func (o *CreateMigrationPartialExportParams) WithTimeout(timeout time.Duration) *CreateMigrationPartialExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create migration partial export params
func (o *CreateMigrationPartialExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create migration partial export params
func (o *CreateMigrationPartialExportParams) WithContext(ctx context.Context) *CreateMigrationPartialExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create migration partial export params
func (o *CreateMigrationPartialExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create migration partial export params
func (o *CreateMigrationPartialExportParams) WithHTTPClient(client *http.Client) *CreateMigrationPartialExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create migration partial export params
func (o *CreateMigrationPartialExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMigrationPartialExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
