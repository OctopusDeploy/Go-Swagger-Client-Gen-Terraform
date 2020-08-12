// Code generated by go-swagger; DO NOT EDIT.

package feeds

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

// NewCreateFeedSpacesParams creates a new CreateFeedSpacesParams object
// with the default values initialized.
func NewCreateFeedSpacesParams() *CreateFeedSpacesParams {
	var ()
	return &CreateFeedSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFeedSpacesParamsWithTimeout creates a new CreateFeedSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateFeedSpacesParamsWithTimeout(timeout time.Duration) *CreateFeedSpacesParams {
	var ()
	return &CreateFeedSpacesParams{

		timeout: timeout,
	}
}

// NewCreateFeedSpacesParamsWithContext creates a new CreateFeedSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateFeedSpacesParamsWithContext(ctx context.Context) *CreateFeedSpacesParams {
	var ()
	return &CreateFeedSpacesParams{

		Context: ctx,
	}
}

// NewCreateFeedSpacesParamsWithHTTPClient creates a new CreateFeedSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateFeedSpacesParamsWithHTTPClient(client *http.Client) *CreateFeedSpacesParams {
	var ()
	return &CreateFeedSpacesParams{
		HTTPClient: client,
	}
}

/*CreateFeedSpacesParams contains all the parameters to send to the API endpoint
for the create feed spaces operation typically these are written to a http.Request
*/
type CreateFeedSpacesParams struct {

	/*FeedResource
	  The FeedResource resource to create

	*/
	FeedResource *models.FeedResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create feed spaces params
func (o *CreateFeedSpacesParams) WithTimeout(timeout time.Duration) *CreateFeedSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create feed spaces params
func (o *CreateFeedSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create feed spaces params
func (o *CreateFeedSpacesParams) WithContext(ctx context.Context) *CreateFeedSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create feed spaces params
func (o *CreateFeedSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create feed spaces params
func (o *CreateFeedSpacesParams) WithHTTPClient(client *http.Client) *CreateFeedSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create feed spaces params
func (o *CreateFeedSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeedResource adds the feedResource to the create feed spaces params
func (o *CreateFeedSpacesParams) WithFeedResource(feedResource *models.FeedResource) *CreateFeedSpacesParams {
	o.SetFeedResource(feedResource)
	return o
}

// SetFeedResource adds the feedResource to the create feed spaces params
func (o *CreateFeedSpacesParams) SetFeedResource(feedResource *models.FeedResource) {
	o.FeedResource = feedResource
}

// WithBaseSpaceID adds the baseSpaceID to the create feed spaces params
func (o *CreateFeedSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateFeedSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create feed spaces params
func (o *CreateFeedSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFeedSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FeedResource != nil {
		if err := r.SetBodyParam(o.FeedResource); err != nil {
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