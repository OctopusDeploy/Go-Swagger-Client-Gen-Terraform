// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// NewDeleteSubscriptionSpacesParams creates a new DeleteSubscriptionSpacesParams object
// with the default values initialized.
func NewDeleteSubscriptionSpacesParams() *DeleteSubscriptionSpacesParams {
	var ()
	return &DeleteSubscriptionSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSubscriptionSpacesParamsWithTimeout creates a new DeleteSubscriptionSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSubscriptionSpacesParamsWithTimeout(timeout time.Duration) *DeleteSubscriptionSpacesParams {
	var ()
	return &DeleteSubscriptionSpacesParams{

		timeout: timeout,
	}
}

// NewDeleteSubscriptionSpacesParamsWithContext creates a new DeleteSubscriptionSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSubscriptionSpacesParamsWithContext(ctx context.Context) *DeleteSubscriptionSpacesParams {
	var ()
	return &DeleteSubscriptionSpacesParams{

		Context: ctx,
	}
}

// NewDeleteSubscriptionSpacesParamsWithHTTPClient creates a new DeleteSubscriptionSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSubscriptionSpacesParamsWithHTTPClient(client *http.Client) *DeleteSubscriptionSpacesParams {
	var ()
	return &DeleteSubscriptionSpacesParams{
		HTTPClient: client,
	}
}

/*DeleteSubscriptionSpacesParams contains all the parameters to send to the API endpoint
for the delete subscription spaces operation typically these are written to a http.Request
*/
type DeleteSubscriptionSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the SubscriptionResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete subscription spaces params
func (o *DeleteSubscriptionSpacesParams) WithTimeout(timeout time.Duration) *DeleteSubscriptionSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete subscription spaces params
func (o *DeleteSubscriptionSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete subscription spaces params
func (o *DeleteSubscriptionSpacesParams) WithContext(ctx context.Context) *DeleteSubscriptionSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete subscription spaces params
func (o *DeleteSubscriptionSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete subscription spaces params
func (o *DeleteSubscriptionSpacesParams) WithHTTPClient(client *http.Client) *DeleteSubscriptionSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete subscription spaces params
func (o *DeleteSubscriptionSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the delete subscription spaces params
func (o *DeleteSubscriptionSpacesParams) WithBaseSpaceID(baseSpaceID string) *DeleteSubscriptionSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the delete subscription spaces params
func (o *DeleteSubscriptionSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the delete subscription spaces params
func (o *DeleteSubscriptionSpacesParams) WithID(id string) *DeleteSubscriptionSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete subscription spaces params
func (o *DeleteSubscriptionSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSubscriptionSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
