// Code generated by go-swagger; DO NOT EDIT.

package worker_pools

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

// NewListAllWorkerPoolsSpacesParams creates a new ListAllWorkerPoolsSpacesParams object
// with the default values initialized.
func NewListAllWorkerPoolsSpacesParams() *ListAllWorkerPoolsSpacesParams {
	var ()
	return &ListAllWorkerPoolsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllWorkerPoolsSpacesParamsWithTimeout creates a new ListAllWorkerPoolsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllWorkerPoolsSpacesParamsWithTimeout(timeout time.Duration) *ListAllWorkerPoolsSpacesParams {
	var ()
	return &ListAllWorkerPoolsSpacesParams{

		timeout: timeout,
	}
}

// NewListAllWorkerPoolsSpacesParamsWithContext creates a new ListAllWorkerPoolsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllWorkerPoolsSpacesParamsWithContext(ctx context.Context) *ListAllWorkerPoolsSpacesParams {
	var ()
	return &ListAllWorkerPoolsSpacesParams{

		Context: ctx,
	}
}

// NewListAllWorkerPoolsSpacesParamsWithHTTPClient creates a new ListAllWorkerPoolsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllWorkerPoolsSpacesParamsWithHTTPClient(client *http.Client) *ListAllWorkerPoolsSpacesParams {
	var ()
	return &ListAllWorkerPoolsSpacesParams{
		HTTPClient: client,
	}
}

/*ListAllWorkerPoolsSpacesParams contains all the parameters to send to the API endpoint
for the list all worker pools spaces operation typically these are written to a http.Request
*/
type ListAllWorkerPoolsSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all worker pools spaces params
func (o *ListAllWorkerPoolsSpacesParams) WithTimeout(timeout time.Duration) *ListAllWorkerPoolsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all worker pools spaces params
func (o *ListAllWorkerPoolsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all worker pools spaces params
func (o *ListAllWorkerPoolsSpacesParams) WithContext(ctx context.Context) *ListAllWorkerPoolsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all worker pools spaces params
func (o *ListAllWorkerPoolsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all worker pools spaces params
func (o *ListAllWorkerPoolsSpacesParams) WithHTTPClient(client *http.Client) *ListAllWorkerPoolsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all worker pools spaces params
func (o *ListAllWorkerPoolsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the list all worker pools spaces params
func (o *ListAllWorkerPoolsSpacesParams) WithBaseSpaceID(baseSpaceID string) *ListAllWorkerPoolsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the list all worker pools spaces params
func (o *ListAllWorkerPoolsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllWorkerPoolsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
