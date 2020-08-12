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

// NewCreateTagSetSpacesParams creates a new CreateTagSetSpacesParams object
// with the default values initialized.
func NewCreateTagSetSpacesParams() *CreateTagSetSpacesParams {
	var ()
	return &CreateTagSetSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTagSetSpacesParamsWithTimeout creates a new CreateTagSetSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTagSetSpacesParamsWithTimeout(timeout time.Duration) *CreateTagSetSpacesParams {
	var ()
	return &CreateTagSetSpacesParams{

		timeout: timeout,
	}
}

// NewCreateTagSetSpacesParamsWithContext creates a new CreateTagSetSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTagSetSpacesParamsWithContext(ctx context.Context) *CreateTagSetSpacesParams {
	var ()
	return &CreateTagSetSpacesParams{

		Context: ctx,
	}
}

// NewCreateTagSetSpacesParamsWithHTTPClient creates a new CreateTagSetSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTagSetSpacesParamsWithHTTPClient(client *http.Client) *CreateTagSetSpacesParams {
	var ()
	return &CreateTagSetSpacesParams{
		HTTPClient: client,
	}
}

/*CreateTagSetSpacesParams contains all the parameters to send to the API endpoint
for the create tag set spaces operation typically these are written to a http.Request
*/
type CreateTagSetSpacesParams struct {

	/*TagSetResource
	  The TagSetResource resource to create

	*/
	TagSetResource *models.TagSetResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create tag set spaces params
func (o *CreateTagSetSpacesParams) WithTimeout(timeout time.Duration) *CreateTagSetSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create tag set spaces params
func (o *CreateTagSetSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create tag set spaces params
func (o *CreateTagSetSpacesParams) WithContext(ctx context.Context) *CreateTagSetSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create tag set spaces params
func (o *CreateTagSetSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create tag set spaces params
func (o *CreateTagSetSpacesParams) WithHTTPClient(client *http.Client) *CreateTagSetSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create tag set spaces params
func (o *CreateTagSetSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagSetResource adds the tagSetResource to the create tag set spaces params
func (o *CreateTagSetSpacesParams) WithTagSetResource(tagSetResource *models.TagSetResource) *CreateTagSetSpacesParams {
	o.SetTagSetResource(tagSetResource)
	return o
}

// SetTagSetResource adds the tagSetResource to the create tag set spaces params
func (o *CreateTagSetSpacesParams) SetTagSetResource(tagSetResource *models.TagSetResource) {
	o.TagSetResource = tagSetResource
}

// WithBaseSpaceID adds the baseSpaceID to the create tag set spaces params
func (o *CreateTagSetSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateTagSetSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create tag set spaces params
func (o *CreateTagSetSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTagSetSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
