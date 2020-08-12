// Code generated by go-swagger; DO NOT EDIT.

package channels

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

// NewCreateChannelSpacesParams creates a new CreateChannelSpacesParams object
// with the default values initialized.
func NewCreateChannelSpacesParams() *CreateChannelSpacesParams {
	var ()
	return &CreateChannelSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateChannelSpacesParamsWithTimeout creates a new CreateChannelSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateChannelSpacesParamsWithTimeout(timeout time.Duration) *CreateChannelSpacesParams {
	var ()
	return &CreateChannelSpacesParams{

		timeout: timeout,
	}
}

// NewCreateChannelSpacesParamsWithContext creates a new CreateChannelSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateChannelSpacesParamsWithContext(ctx context.Context) *CreateChannelSpacesParams {
	var ()
	return &CreateChannelSpacesParams{

		Context: ctx,
	}
}

// NewCreateChannelSpacesParamsWithHTTPClient creates a new CreateChannelSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateChannelSpacesParamsWithHTTPClient(client *http.Client) *CreateChannelSpacesParams {
	var ()
	return &CreateChannelSpacesParams{
		HTTPClient: client,
	}
}

/*CreateChannelSpacesParams contains all the parameters to send to the API endpoint
for the create channel spaces operation typically these are written to a http.Request
*/
type CreateChannelSpacesParams struct {

	/*ChannelResource
	  The ChannelResource resource to create

	*/
	ChannelResource *models.ChannelResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create channel spaces params
func (o *CreateChannelSpacesParams) WithTimeout(timeout time.Duration) *CreateChannelSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create channel spaces params
func (o *CreateChannelSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create channel spaces params
func (o *CreateChannelSpacesParams) WithContext(ctx context.Context) *CreateChannelSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create channel spaces params
func (o *CreateChannelSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create channel spaces params
func (o *CreateChannelSpacesParams) WithHTTPClient(client *http.Client) *CreateChannelSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create channel spaces params
func (o *CreateChannelSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelResource adds the channelResource to the create channel spaces params
func (o *CreateChannelSpacesParams) WithChannelResource(channelResource *models.ChannelResource) *CreateChannelSpacesParams {
	o.SetChannelResource(channelResource)
	return o
}

// SetChannelResource adds the channelResource to the create channel spaces params
func (o *CreateChannelSpacesParams) SetChannelResource(channelResource *models.ChannelResource) {
	o.ChannelResource = channelResource
}

// WithBaseSpaceID adds the baseSpaceID to the create channel spaces params
func (o *CreateChannelSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateChannelSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create channel spaces params
func (o *CreateChannelSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateChannelSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChannelResource != nil {
		if err := r.SetBodyParam(o.ChannelResource); err != nil {
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
