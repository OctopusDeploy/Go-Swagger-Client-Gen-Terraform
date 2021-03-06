// Code generated by go-swagger; DO NOT EDIT.

package interruptions

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

// NewCreateSubmitInterruptionSpacesParams creates a new CreateSubmitInterruptionSpacesParams object
// with the default values initialized.
func NewCreateSubmitInterruptionSpacesParams() *CreateSubmitInterruptionSpacesParams {
	var ()
	return &CreateSubmitInterruptionSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSubmitInterruptionSpacesParamsWithTimeout creates a new CreateSubmitInterruptionSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSubmitInterruptionSpacesParamsWithTimeout(timeout time.Duration) *CreateSubmitInterruptionSpacesParams {
	var ()
	return &CreateSubmitInterruptionSpacesParams{

		timeout: timeout,
	}
}

// NewCreateSubmitInterruptionSpacesParamsWithContext creates a new CreateSubmitInterruptionSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSubmitInterruptionSpacesParamsWithContext(ctx context.Context) *CreateSubmitInterruptionSpacesParams {
	var ()
	return &CreateSubmitInterruptionSpacesParams{

		Context: ctx,
	}
}

// NewCreateSubmitInterruptionSpacesParamsWithHTTPClient creates a new CreateSubmitInterruptionSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSubmitInterruptionSpacesParamsWithHTTPClient(client *http.Client) *CreateSubmitInterruptionSpacesParams {
	var ()
	return &CreateSubmitInterruptionSpacesParams{
		HTTPClient: client,
	}
}

/*CreateSubmitInterruptionSpacesParams contains all the parameters to send to the API endpoint
for the create submit interruption spaces operation typically these are written to a http.Request
*/
type CreateSubmitInterruptionSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create submit interruption spaces params
func (o *CreateSubmitInterruptionSpacesParams) WithTimeout(timeout time.Duration) *CreateSubmitInterruptionSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create submit interruption spaces params
func (o *CreateSubmitInterruptionSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create submit interruption spaces params
func (o *CreateSubmitInterruptionSpacesParams) WithContext(ctx context.Context) *CreateSubmitInterruptionSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create submit interruption spaces params
func (o *CreateSubmitInterruptionSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create submit interruption spaces params
func (o *CreateSubmitInterruptionSpacesParams) WithHTTPClient(client *http.Client) *CreateSubmitInterruptionSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create submit interruption spaces params
func (o *CreateSubmitInterruptionSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the create submit interruption spaces params
func (o *CreateSubmitInterruptionSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateSubmitInterruptionSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create submit interruption spaces params
func (o *CreateSubmitInterruptionSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the create submit interruption spaces params
func (o *CreateSubmitInterruptionSpacesParams) WithID(id string) *CreateSubmitInterruptionSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create submit interruption spaces params
func (o *CreateSubmitInterruptionSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubmitInterruptionSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
