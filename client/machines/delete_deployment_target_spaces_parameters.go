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

// NewDeleteDeploymentTargetSpacesParams creates a new DeleteDeploymentTargetSpacesParams object
// with the default values initialized.
func NewDeleteDeploymentTargetSpacesParams() *DeleteDeploymentTargetSpacesParams {
	var ()
	return &DeleteDeploymentTargetSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeploymentTargetSpacesParamsWithTimeout creates a new DeleteDeploymentTargetSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeploymentTargetSpacesParamsWithTimeout(timeout time.Duration) *DeleteDeploymentTargetSpacesParams {
	var ()
	return &DeleteDeploymentTargetSpacesParams{

		timeout: timeout,
	}
}

// NewDeleteDeploymentTargetSpacesParamsWithContext creates a new DeleteDeploymentTargetSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeploymentTargetSpacesParamsWithContext(ctx context.Context) *DeleteDeploymentTargetSpacesParams {
	var ()
	return &DeleteDeploymentTargetSpacesParams{

		Context: ctx,
	}
}

// NewDeleteDeploymentTargetSpacesParamsWithHTTPClient creates a new DeleteDeploymentTargetSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeploymentTargetSpacesParamsWithHTTPClient(client *http.Client) *DeleteDeploymentTargetSpacesParams {
	var ()
	return &DeleteDeploymentTargetSpacesParams{
		HTTPClient: client,
	}
}

/*DeleteDeploymentTargetSpacesParams contains all the parameters to send to the API endpoint
for the delete deployment target spaces operation typically these are written to a http.Request
*/
type DeleteDeploymentTargetSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the DeploymentTargetResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete deployment target spaces params
func (o *DeleteDeploymentTargetSpacesParams) WithTimeout(timeout time.Duration) *DeleteDeploymentTargetSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete deployment target spaces params
func (o *DeleteDeploymentTargetSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete deployment target spaces params
func (o *DeleteDeploymentTargetSpacesParams) WithContext(ctx context.Context) *DeleteDeploymentTargetSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete deployment target spaces params
func (o *DeleteDeploymentTargetSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete deployment target spaces params
func (o *DeleteDeploymentTargetSpacesParams) WithHTTPClient(client *http.Client) *DeleteDeploymentTargetSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete deployment target spaces params
func (o *DeleteDeploymentTargetSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the delete deployment target spaces params
func (o *DeleteDeploymentTargetSpacesParams) WithBaseSpaceID(baseSpaceID string) *DeleteDeploymentTargetSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the delete deployment target spaces params
func (o *DeleteDeploymentTargetSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the delete deployment target spaces params
func (o *DeleteDeploymentTargetSpacesParams) WithID(id string) *DeleteDeploymentTargetSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete deployment target spaces params
func (o *DeleteDeploymentTargetSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeploymentTargetSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
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
