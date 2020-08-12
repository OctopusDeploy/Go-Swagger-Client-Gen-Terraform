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
)

// NewListAllFeedsSpacesParams creates a new ListAllFeedsSpacesParams object
// with the default values initialized.
func NewListAllFeedsSpacesParams() *ListAllFeedsSpacesParams {
	var ()
	return &ListAllFeedsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllFeedsSpacesParamsWithTimeout creates a new ListAllFeedsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllFeedsSpacesParamsWithTimeout(timeout time.Duration) *ListAllFeedsSpacesParams {
	var ()
	return &ListAllFeedsSpacesParams{

		timeout: timeout,
	}
}

// NewListAllFeedsSpacesParamsWithContext creates a new ListAllFeedsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllFeedsSpacesParamsWithContext(ctx context.Context) *ListAllFeedsSpacesParams {
	var ()
	return &ListAllFeedsSpacesParams{

		Context: ctx,
	}
}

// NewListAllFeedsSpacesParamsWithHTTPClient creates a new ListAllFeedsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllFeedsSpacesParamsWithHTTPClient(client *http.Client) *ListAllFeedsSpacesParams {
	var ()
	return &ListAllFeedsSpacesParams{
		HTTPClient: client,
	}
}

/*ListAllFeedsSpacesParams contains all the parameters to send to the API endpoint
for the list all feeds spaces operation typically these are written to a http.Request
*/
type ListAllFeedsSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all feeds spaces params
func (o *ListAllFeedsSpacesParams) WithTimeout(timeout time.Duration) *ListAllFeedsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all feeds spaces params
func (o *ListAllFeedsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all feeds spaces params
func (o *ListAllFeedsSpacesParams) WithContext(ctx context.Context) *ListAllFeedsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all feeds spaces params
func (o *ListAllFeedsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all feeds spaces params
func (o *ListAllFeedsSpacesParams) WithHTTPClient(client *http.Client) *ListAllFeedsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all feeds spaces params
func (o *ListAllFeedsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the list all feeds spaces params
func (o *ListAllFeedsSpacesParams) WithBaseSpaceID(baseSpaceID string) *ListAllFeedsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the list all feeds spaces params
func (o *ListAllFeedsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllFeedsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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