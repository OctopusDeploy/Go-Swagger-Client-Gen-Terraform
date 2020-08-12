// Code generated by go-swagger; DO NOT EDIT.

package project_groups

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
	"github.com/go-openapi/swag"
)

// NewIndexProjectGroupProjectsSpacesParams creates a new IndexProjectGroupProjectsSpacesParams object
// with the default values initialized.
func NewIndexProjectGroupProjectsSpacesParams() *IndexProjectGroupProjectsSpacesParams {
	var ()
	return &IndexProjectGroupProjectsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexProjectGroupProjectsSpacesParamsWithTimeout creates a new IndexProjectGroupProjectsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexProjectGroupProjectsSpacesParamsWithTimeout(timeout time.Duration) *IndexProjectGroupProjectsSpacesParams {
	var ()
	return &IndexProjectGroupProjectsSpacesParams{

		timeout: timeout,
	}
}

// NewIndexProjectGroupProjectsSpacesParamsWithContext creates a new IndexProjectGroupProjectsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexProjectGroupProjectsSpacesParamsWithContext(ctx context.Context) *IndexProjectGroupProjectsSpacesParams {
	var ()
	return &IndexProjectGroupProjectsSpacesParams{

		Context: ctx,
	}
}

// NewIndexProjectGroupProjectsSpacesParamsWithHTTPClient creates a new IndexProjectGroupProjectsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexProjectGroupProjectsSpacesParamsWithHTTPClient(client *http.Client) *IndexProjectGroupProjectsSpacesParams {
	var ()
	return &IndexProjectGroupProjectsSpacesParams{
		HTTPClient: client,
	}
}

/*IndexProjectGroupProjectsSpacesParams contains all the parameters to send to the API endpoint
for the index project group projects spaces operation typically these are written to a http.Request
*/
type IndexProjectGroupProjectsSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the ProjectGroup

	*/
	ID string
	/*Skip
	  Number of items to skip

	*/
	Skip *int32
	/*Take
	  Number of items per page

	*/
	Take *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) WithTimeout(timeout time.Duration) *IndexProjectGroupProjectsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) WithContext(ctx context.Context) *IndexProjectGroupProjectsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) WithHTTPClient(client *http.Client) *IndexProjectGroupProjectsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) WithBaseSpaceID(baseSpaceID string) *IndexProjectGroupProjectsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) WithID(id string) *IndexProjectGroupProjectsSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) SetID(id string) {
	o.ID = id
}

// WithSkip adds the skip to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) WithSkip(skip *int32) *IndexProjectGroupProjectsSpacesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) WithTake(take *int32) *IndexProjectGroupProjectsSpacesParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index project group projects spaces params
func (o *IndexProjectGroupProjectsSpacesParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexProjectGroupProjectsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Skip != nil {

		// query param skip
		var qrSkip int32
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt32(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	if o.Take != nil {

		// query param take
		var qrTake int32
		if o.Take != nil {
			qrTake = *o.Take
		}
		qTake := swag.FormatInt32(qrTake)
		if qTake != "" {
			if err := r.SetQueryParam("take", qTake); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
