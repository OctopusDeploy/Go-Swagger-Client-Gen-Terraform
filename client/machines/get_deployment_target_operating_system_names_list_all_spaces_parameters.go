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

// NewGetDeploymentTargetOperatingSystemNamesListAllSpacesParams creates a new GetDeploymentTargetOperatingSystemNamesListAllSpacesParams object
// with the default values initialized.
func NewGetDeploymentTargetOperatingSystemNamesListAllSpacesParams() *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams {
	var ()
	return &GetDeploymentTargetOperatingSystemNamesListAllSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentTargetOperatingSystemNamesListAllSpacesParamsWithTimeout creates a new GetDeploymentTargetOperatingSystemNamesListAllSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentTargetOperatingSystemNamesListAllSpacesParamsWithTimeout(timeout time.Duration) *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams {
	var ()
	return &GetDeploymentTargetOperatingSystemNamesListAllSpacesParams{

		timeout: timeout,
	}
}

// NewGetDeploymentTargetOperatingSystemNamesListAllSpacesParamsWithContext creates a new GetDeploymentTargetOperatingSystemNamesListAllSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentTargetOperatingSystemNamesListAllSpacesParamsWithContext(ctx context.Context) *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams {
	var ()
	return &GetDeploymentTargetOperatingSystemNamesListAllSpacesParams{

		Context: ctx,
	}
}

// NewGetDeploymentTargetOperatingSystemNamesListAllSpacesParamsWithHTTPClient creates a new GetDeploymentTargetOperatingSystemNamesListAllSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentTargetOperatingSystemNamesListAllSpacesParamsWithHTTPClient(client *http.Client) *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams {
	var ()
	return &GetDeploymentTargetOperatingSystemNamesListAllSpacesParams{
		HTTPClient: client,
	}
}

/*GetDeploymentTargetOperatingSystemNamesListAllSpacesParams contains all the parameters to send to the API endpoint
for the get deployment target operating system names list all spaces operation typically these are written to a http.Request
*/
type GetDeploymentTargetOperatingSystemNamesListAllSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment target operating system names list all spaces params
func (o *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams) WithTimeout(timeout time.Duration) *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment target operating system names list all spaces params
func (o *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment target operating system names list all spaces params
func (o *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams) WithContext(ctx context.Context) *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment target operating system names list all spaces params
func (o *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment target operating system names list all spaces params
func (o *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams) WithHTTPClient(client *http.Client) *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment target operating system names list all spaces params
func (o *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get deployment target operating system names list all spaces params
func (o *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get deployment target operating system names list all spaces params
func (o *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
