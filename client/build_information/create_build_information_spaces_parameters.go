// Code generated by go-swagger; DO NOT EDIT.

package build_information

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

// NewCreateBuildInformationSpacesParams creates a new CreateBuildInformationSpacesParams object
// with the default values initialized.
func NewCreateBuildInformationSpacesParams() *CreateBuildInformationSpacesParams {
	var ()
	return &CreateBuildInformationSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBuildInformationSpacesParamsWithTimeout creates a new CreateBuildInformationSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBuildInformationSpacesParamsWithTimeout(timeout time.Duration) *CreateBuildInformationSpacesParams {
	var ()
	return &CreateBuildInformationSpacesParams{

		timeout: timeout,
	}
}

// NewCreateBuildInformationSpacesParamsWithContext creates a new CreateBuildInformationSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBuildInformationSpacesParamsWithContext(ctx context.Context) *CreateBuildInformationSpacesParams {
	var ()
	return &CreateBuildInformationSpacesParams{

		Context: ctx,
	}
}

// NewCreateBuildInformationSpacesParamsWithHTTPClient creates a new CreateBuildInformationSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBuildInformationSpacesParamsWithHTTPClient(client *http.Client) *CreateBuildInformationSpacesParams {
	var ()
	return &CreateBuildInformationSpacesParams{
		HTTPClient: client,
	}
}

/*CreateBuildInformationSpacesParams contains all the parameters to send to the API endpoint
for the create build information spaces operation typically these are written to a http.Request
*/
type CreateBuildInformationSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create build information spaces params
func (o *CreateBuildInformationSpacesParams) WithTimeout(timeout time.Duration) *CreateBuildInformationSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create build information spaces params
func (o *CreateBuildInformationSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create build information spaces params
func (o *CreateBuildInformationSpacesParams) WithContext(ctx context.Context) *CreateBuildInformationSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create build information spaces params
func (o *CreateBuildInformationSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create build information spaces params
func (o *CreateBuildInformationSpacesParams) WithHTTPClient(client *http.Client) *CreateBuildInformationSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create build information spaces params
func (o *CreateBuildInformationSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the create build information spaces params
func (o *CreateBuildInformationSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateBuildInformationSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create build information spaces params
func (o *CreateBuildInformationSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBuildInformationSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
