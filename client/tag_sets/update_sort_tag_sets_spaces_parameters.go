// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

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

// NewUpdateSortTagSetsSpacesParams creates a new UpdateSortTagSetsSpacesParams object
// with the default values initialized.
func NewUpdateSortTagSetsSpacesParams() *UpdateSortTagSetsSpacesParams {
	var ()
	return &UpdateSortTagSetsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSortTagSetsSpacesParamsWithTimeout creates a new UpdateSortTagSetsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSortTagSetsSpacesParamsWithTimeout(timeout time.Duration) *UpdateSortTagSetsSpacesParams {
	var ()
	return &UpdateSortTagSetsSpacesParams{

		timeout: timeout,
	}
}

// NewUpdateSortTagSetsSpacesParamsWithContext creates a new UpdateSortTagSetsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSortTagSetsSpacesParamsWithContext(ctx context.Context) *UpdateSortTagSetsSpacesParams {
	var ()
	return &UpdateSortTagSetsSpacesParams{

		Context: ctx,
	}
}

// NewUpdateSortTagSetsSpacesParamsWithHTTPClient creates a new UpdateSortTagSetsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSortTagSetsSpacesParamsWithHTTPClient(client *http.Client) *UpdateSortTagSetsSpacesParams {
	var ()
	return &UpdateSortTagSetsSpacesParams{
		HTTPClient: client,
	}
}

/*UpdateSortTagSetsSpacesParams contains all the parameters to send to the API endpoint
for the update sort tag sets spaces operation typically these are written to a http.Request
*/
type UpdateSortTagSetsSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update sort tag sets spaces params
func (o *UpdateSortTagSetsSpacesParams) WithTimeout(timeout time.Duration) *UpdateSortTagSetsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update sort tag sets spaces params
func (o *UpdateSortTagSetsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update sort tag sets spaces params
func (o *UpdateSortTagSetsSpacesParams) WithContext(ctx context.Context) *UpdateSortTagSetsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update sort tag sets spaces params
func (o *UpdateSortTagSetsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update sort tag sets spaces params
func (o *UpdateSortTagSetsSpacesParams) WithHTTPClient(client *http.Client) *UpdateSortTagSetsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update sort tag sets spaces params
func (o *UpdateSortTagSetsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the update sort tag sets spaces params
func (o *UpdateSortTagSetsSpacesParams) WithBaseSpaceID(baseSpaceID string) *UpdateSortTagSetsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the update sort tag sets spaces params
func (o *UpdateSortTagSetsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSortTagSetsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
