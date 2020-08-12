// Code generated by go-swagger; DO NOT EDIT.

package artifacts

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

// NewGetArtifactContentSpacesParams creates a new GetArtifactContentSpacesParams object
// with the default values initialized.
func NewGetArtifactContentSpacesParams() *GetArtifactContentSpacesParams {
	var ()
	return &GetArtifactContentSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArtifactContentSpacesParamsWithTimeout creates a new GetArtifactContentSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArtifactContentSpacesParamsWithTimeout(timeout time.Duration) *GetArtifactContentSpacesParams {
	var ()
	return &GetArtifactContentSpacesParams{

		timeout: timeout,
	}
}

// NewGetArtifactContentSpacesParamsWithContext creates a new GetArtifactContentSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArtifactContentSpacesParamsWithContext(ctx context.Context) *GetArtifactContentSpacesParams {
	var ()
	return &GetArtifactContentSpacesParams{

		Context: ctx,
	}
}

// NewGetArtifactContentSpacesParamsWithHTTPClient creates a new GetArtifactContentSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArtifactContentSpacesParamsWithHTTPClient(client *http.Client) *GetArtifactContentSpacesParams {
	var ()
	return &GetArtifactContentSpacesParams{
		HTTPClient: client,
	}
}

/*GetArtifactContentSpacesParams contains all the parameters to send to the API endpoint
for the get artifact content spaces operation typically these are written to a http.Request
*/
type GetArtifactContentSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the artifact

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get artifact content spaces params
func (o *GetArtifactContentSpacesParams) WithTimeout(timeout time.Duration) *GetArtifactContentSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get artifact content spaces params
func (o *GetArtifactContentSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get artifact content spaces params
func (o *GetArtifactContentSpacesParams) WithContext(ctx context.Context) *GetArtifactContentSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get artifact content spaces params
func (o *GetArtifactContentSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get artifact content spaces params
func (o *GetArtifactContentSpacesParams) WithHTTPClient(client *http.Client) *GetArtifactContentSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get artifact content spaces params
func (o *GetArtifactContentSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get artifact content spaces params
func (o *GetArtifactContentSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetArtifactContentSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get artifact content spaces params
func (o *GetArtifactContentSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get artifact content spaces params
func (o *GetArtifactContentSpacesParams) WithID(id string) *GetArtifactContentSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get artifact content spaces params
func (o *GetArtifactContentSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetArtifactContentSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
