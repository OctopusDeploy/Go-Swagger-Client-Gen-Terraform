// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// NewDeleteDeploymentSpacesParams creates a new DeleteDeploymentSpacesParams object
// with the default values initialized.
func NewDeleteDeploymentSpacesParams() *DeleteDeploymentSpacesParams {
	var ()
	return &DeleteDeploymentSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeploymentSpacesParamsWithTimeout creates a new DeleteDeploymentSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeploymentSpacesParamsWithTimeout(timeout time.Duration) *DeleteDeploymentSpacesParams {
	var ()
	return &DeleteDeploymentSpacesParams{

		timeout: timeout,
	}
}

// NewDeleteDeploymentSpacesParamsWithContext creates a new DeleteDeploymentSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeploymentSpacesParamsWithContext(ctx context.Context) *DeleteDeploymentSpacesParams {
	var ()
	return &DeleteDeploymentSpacesParams{

		Context: ctx,
	}
}

// NewDeleteDeploymentSpacesParamsWithHTTPClient creates a new DeleteDeploymentSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeploymentSpacesParamsWithHTTPClient(client *http.Client) *DeleteDeploymentSpacesParams {
	var ()
	return &DeleteDeploymentSpacesParams{
		HTTPClient: client,
	}
}

/*DeleteDeploymentSpacesParams contains all the parameters to send to the API endpoint
for the delete deployment spaces operation typically these are written to a http.Request
*/
type DeleteDeploymentSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the DeploymentResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete deployment spaces params
func (o *DeleteDeploymentSpacesParams) WithTimeout(timeout time.Duration) *DeleteDeploymentSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete deployment spaces params
func (o *DeleteDeploymentSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete deployment spaces params
func (o *DeleteDeploymentSpacesParams) WithContext(ctx context.Context) *DeleteDeploymentSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete deployment spaces params
func (o *DeleteDeploymentSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete deployment spaces params
func (o *DeleteDeploymentSpacesParams) WithHTTPClient(client *http.Client) *DeleteDeploymentSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete deployment spaces params
func (o *DeleteDeploymentSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the delete deployment spaces params
func (o *DeleteDeploymentSpacesParams) WithBaseSpaceID(baseSpaceID string) *DeleteDeploymentSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the delete deployment spaces params
func (o *DeleteDeploymentSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the delete deployment spaces params
func (o *DeleteDeploymentSpacesParams) WithID(id string) *DeleteDeploymentSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete deployment spaces params
func (o *DeleteDeploymentSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeploymentSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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