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

	"models"
)

// NewUpdateTagSetSpacesParams creates a new UpdateTagSetSpacesParams object
// with the default values initialized.
func NewUpdateTagSetSpacesParams() *UpdateTagSetSpacesParams {
	var ()
	return &UpdateTagSetSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTagSetSpacesParamsWithTimeout creates a new UpdateTagSetSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTagSetSpacesParamsWithTimeout(timeout time.Duration) *UpdateTagSetSpacesParams {
	var ()
	return &UpdateTagSetSpacesParams{

		timeout: timeout,
	}
}

// NewUpdateTagSetSpacesParamsWithContext creates a new UpdateTagSetSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTagSetSpacesParamsWithContext(ctx context.Context) *UpdateTagSetSpacesParams {
	var ()
	return &UpdateTagSetSpacesParams{

		Context: ctx,
	}
}

// NewUpdateTagSetSpacesParamsWithHTTPClient creates a new UpdateTagSetSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTagSetSpacesParamsWithHTTPClient(client *http.Client) *UpdateTagSetSpacesParams {
	var ()
	return &UpdateTagSetSpacesParams{
		HTTPClient: client,
	}
}

/*UpdateTagSetSpacesParams contains all the parameters to send to the API endpoint
for the update tag set spaces operation typically these are written to a http.Request
*/
type UpdateTagSetSpacesParams struct {

	/*TagSetResource
	  The TagSetResource resource to create

	*/
	TagSetResource *models.TagSetResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the TagSetResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) WithTimeout(timeout time.Duration) *UpdateTagSetSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) WithContext(ctx context.Context) *UpdateTagSetSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) WithHTTPClient(client *http.Client) *UpdateTagSetSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagSetResource adds the tagSetResource to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) WithTagSetResource(tagSetResource *models.TagSetResource) *UpdateTagSetSpacesParams {
	o.SetTagSetResource(tagSetResource)
	return o
}

// SetTagSetResource adds the tagSetResource to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) SetTagSetResource(tagSetResource *models.TagSetResource) {
	o.TagSetResource = tagSetResource
}

// WithBaseSpaceID adds the baseSpaceID to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) WithBaseSpaceID(baseSpaceID string) *UpdateTagSetSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) WithID(id string) *UpdateTagSetSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update tag set spaces params
func (o *UpdateTagSetSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTagSetSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TagSetResource != nil {
		if err := r.SetBodyParam(o.TagSetResource); err != nil {
			return err
		}
	}

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
