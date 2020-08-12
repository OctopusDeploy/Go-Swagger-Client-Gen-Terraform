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

// NewGetArtifactByIDParams creates a new GetArtifactByIDParams object
// with the default values initialized.
func NewGetArtifactByIDParams() *GetArtifactByIDParams {
	var ()
	return &GetArtifactByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArtifactByIDParamsWithTimeout creates a new GetArtifactByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArtifactByIDParamsWithTimeout(timeout time.Duration) *GetArtifactByIDParams {
	var ()
	return &GetArtifactByIDParams{

		timeout: timeout,
	}
}

// NewGetArtifactByIDParamsWithContext creates a new GetArtifactByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArtifactByIDParamsWithContext(ctx context.Context) *GetArtifactByIDParams {
	var ()
	return &GetArtifactByIDParams{

		Context: ctx,
	}
}

// NewGetArtifactByIDParamsWithHTTPClient creates a new GetArtifactByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArtifactByIDParamsWithHTTPClient(client *http.Client) *GetArtifactByIDParams {
	var ()
	return &GetArtifactByIDParams{
		HTTPClient: client,
	}
}

/*GetArtifactByIDParams contains all the parameters to send to the API endpoint
for the get artifact by Id operation typically these are written to a http.Request
*/
type GetArtifactByIDParams struct {

	/*ID
	  ID of the ArtifactResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get artifact by Id params
func (o *GetArtifactByIDParams) WithTimeout(timeout time.Duration) *GetArtifactByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get artifact by Id params
func (o *GetArtifactByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get artifact by Id params
func (o *GetArtifactByIDParams) WithContext(ctx context.Context) *GetArtifactByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get artifact by Id params
func (o *GetArtifactByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get artifact by Id params
func (o *GetArtifactByIDParams) WithHTTPClient(client *http.Client) *GetArtifactByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get artifact by Id params
func (o *GetArtifactByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get artifact by Id params
func (o *GetArtifactByIDParams) WithID(id string) *GetArtifactByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get artifact by Id params
func (o *GetArtifactByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetArtifactByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}