// Code generated by go-swagger; DO NOT EDIT.

package releases

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

// NewIndexReleaseDeploymentsSpacesParams creates a new IndexReleaseDeploymentsSpacesParams object
// with the default values initialized.
func NewIndexReleaseDeploymentsSpacesParams() *IndexReleaseDeploymentsSpacesParams {
	var ()
	return &IndexReleaseDeploymentsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexReleaseDeploymentsSpacesParamsWithTimeout creates a new IndexReleaseDeploymentsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexReleaseDeploymentsSpacesParamsWithTimeout(timeout time.Duration) *IndexReleaseDeploymentsSpacesParams {
	var ()
	return &IndexReleaseDeploymentsSpacesParams{

		timeout: timeout,
	}
}

// NewIndexReleaseDeploymentsSpacesParamsWithContext creates a new IndexReleaseDeploymentsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexReleaseDeploymentsSpacesParamsWithContext(ctx context.Context) *IndexReleaseDeploymentsSpacesParams {
	var ()
	return &IndexReleaseDeploymentsSpacesParams{

		Context: ctx,
	}
}

// NewIndexReleaseDeploymentsSpacesParamsWithHTTPClient creates a new IndexReleaseDeploymentsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexReleaseDeploymentsSpacesParamsWithHTTPClient(client *http.Client) *IndexReleaseDeploymentsSpacesParams {
	var ()
	return &IndexReleaseDeploymentsSpacesParams{
		HTTPClient: client,
	}
}

/*IndexReleaseDeploymentsSpacesParams contains all the parameters to send to the API endpoint
for the index release deployments spaces operation typically these are written to a http.Request
*/
type IndexReleaseDeploymentsSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the Release

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

// WithTimeout adds the timeout to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) WithTimeout(timeout time.Duration) *IndexReleaseDeploymentsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) WithContext(ctx context.Context) *IndexReleaseDeploymentsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) WithHTTPClient(client *http.Client) *IndexReleaseDeploymentsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) WithBaseSpaceID(baseSpaceID string) *IndexReleaseDeploymentsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) WithID(id string) *IndexReleaseDeploymentsSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) SetID(id string) {
	o.ID = id
}

// WithSkip adds the skip to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) WithSkip(skip *int32) *IndexReleaseDeploymentsSpacesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) WithTake(take *int32) *IndexReleaseDeploymentsSpacesParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index release deployments spaces params
func (o *IndexReleaseDeploymentsSpacesParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexReleaseDeploymentsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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