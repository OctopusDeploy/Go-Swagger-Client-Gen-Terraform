// Code generated by go-swagger; DO NOT EDIT.

package machines

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

// NewGetDeploymentTargetOperatingSystemShellNameListAllSpacesParams creates a new GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams object
// with the default values initialized.
func NewGetDeploymentTargetOperatingSystemShellNameListAllSpacesParams() *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams {
	var ()
	return &GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentTargetOperatingSystemShellNameListAllSpacesParamsWithTimeout creates a new GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentTargetOperatingSystemShellNameListAllSpacesParamsWithTimeout(timeout time.Duration) *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams {
	var ()
	return &GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams{

		timeout: timeout,
	}
}

// NewGetDeploymentTargetOperatingSystemShellNameListAllSpacesParamsWithContext creates a new GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentTargetOperatingSystemShellNameListAllSpacesParamsWithContext(ctx context.Context) *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams {
	var ()
	return &GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams{

		Context: ctx,
	}
}

// NewGetDeploymentTargetOperatingSystemShellNameListAllSpacesParamsWithHTTPClient creates a new GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentTargetOperatingSystemShellNameListAllSpacesParamsWithHTTPClient(client *http.Client) *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams {
	var ()
	return &GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams{
		HTTPClient: client,
	}
}

/*GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams contains all the parameters to send to the API endpoint
for the get deployment target operating system shell name list all spaces operation typically these are written to a http.Request
*/
type GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment target operating system shell name list all spaces params
func (o *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams) WithTimeout(timeout time.Duration) *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment target operating system shell name list all spaces params
func (o *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment target operating system shell name list all spaces params
func (o *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams) WithContext(ctx context.Context) *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment target operating system shell name list all spaces params
func (o *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment target operating system shell name list all spaces params
func (o *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams) WithHTTPClient(client *http.Client) *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment target operating system shell name list all spaces params
func (o *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get deployment target operating system shell name list all spaces params
func (o *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get deployment target operating system shell name list all spaces params
func (o *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
