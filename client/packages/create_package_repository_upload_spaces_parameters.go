// Code generated by go-swagger; DO NOT EDIT.

package packages

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

// NewCreatePackageRepositoryUploadSpacesParams creates a new CreatePackageRepositoryUploadSpacesParams object
// with the default values initialized.
func NewCreatePackageRepositoryUploadSpacesParams() *CreatePackageRepositoryUploadSpacesParams {
	var ()
	return &CreatePackageRepositoryUploadSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePackageRepositoryUploadSpacesParamsWithTimeout creates a new CreatePackageRepositoryUploadSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePackageRepositoryUploadSpacesParamsWithTimeout(timeout time.Duration) *CreatePackageRepositoryUploadSpacesParams {
	var ()
	return &CreatePackageRepositoryUploadSpacesParams{

		timeout: timeout,
	}
}

// NewCreatePackageRepositoryUploadSpacesParamsWithContext creates a new CreatePackageRepositoryUploadSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePackageRepositoryUploadSpacesParamsWithContext(ctx context.Context) *CreatePackageRepositoryUploadSpacesParams {
	var ()
	return &CreatePackageRepositoryUploadSpacesParams{

		Context: ctx,
	}
}

// NewCreatePackageRepositoryUploadSpacesParamsWithHTTPClient creates a new CreatePackageRepositoryUploadSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePackageRepositoryUploadSpacesParamsWithHTTPClient(client *http.Client) *CreatePackageRepositoryUploadSpacesParams {
	var ()
	return &CreatePackageRepositoryUploadSpacesParams{
		HTTPClient: client,
	}
}

/*CreatePackageRepositoryUploadSpacesParams contains all the parameters to send to the API endpoint
for the create package repository upload spaces operation typically these are written to a http.Request
*/
type CreatePackageRepositoryUploadSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create package repository upload spaces params
func (o *CreatePackageRepositoryUploadSpacesParams) WithTimeout(timeout time.Duration) *CreatePackageRepositoryUploadSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create package repository upload spaces params
func (o *CreatePackageRepositoryUploadSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create package repository upload spaces params
func (o *CreatePackageRepositoryUploadSpacesParams) WithContext(ctx context.Context) *CreatePackageRepositoryUploadSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create package repository upload spaces params
func (o *CreatePackageRepositoryUploadSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create package repository upload spaces params
func (o *CreatePackageRepositoryUploadSpacesParams) WithHTTPClient(client *http.Client) *CreatePackageRepositoryUploadSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create package repository upload spaces params
func (o *CreatePackageRepositoryUploadSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the create package repository upload spaces params
func (o *CreatePackageRepositoryUploadSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreatePackageRepositoryUploadSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create package repository upload spaces params
func (o *CreatePackageRepositoryUploadSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePackageRepositoryUploadSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
