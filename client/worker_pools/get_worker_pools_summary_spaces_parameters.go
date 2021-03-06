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

// NewGetWorkerPoolsSummarySpacesParams creates a new GetWorkerPoolsSummarySpacesParams object
// with the default values initialized.
func NewGetWorkerPoolsSummarySpacesParams() *GetWorkerPoolsSummarySpacesParams {
	var ()
	return &GetWorkerPoolsSummarySpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkerPoolsSummarySpacesParamsWithTimeout creates a new GetWorkerPoolsSummarySpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkerPoolsSummarySpacesParamsWithTimeout(timeout time.Duration) *GetWorkerPoolsSummarySpacesParams {
	var ()
	return &GetWorkerPoolsSummarySpacesParams{

		timeout: timeout,
	}
}

// NewGetWorkerPoolsSummarySpacesParamsWithContext creates a new GetWorkerPoolsSummarySpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkerPoolsSummarySpacesParamsWithContext(ctx context.Context) *GetWorkerPoolsSummarySpacesParams {
	var ()
	return &GetWorkerPoolsSummarySpacesParams{

		Context: ctx,
	}
}

// NewGetWorkerPoolsSummarySpacesParamsWithHTTPClient creates a new GetWorkerPoolsSummarySpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkerPoolsSummarySpacesParamsWithHTTPClient(client *http.Client) *GetWorkerPoolsSummarySpacesParams {
	var ()
	return &GetWorkerPoolsSummarySpacesParams{
		HTTPClient: client,
	}
}

/*GetWorkerPoolsSummarySpacesParams contains all the parameters to send to the API endpoint
for the get worker pools summary spaces operation typically these are written to a http.Request
*/
type GetWorkerPoolsSummarySpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get worker pools summary spaces params
func (o *GetWorkerPoolsSummarySpacesParams) WithTimeout(timeout time.Duration) *GetWorkerPoolsSummarySpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get worker pools summary spaces params
func (o *GetWorkerPoolsSummarySpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get worker pools summary spaces params
func (o *GetWorkerPoolsSummarySpacesParams) WithContext(ctx context.Context) *GetWorkerPoolsSummarySpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get worker pools summary spaces params
func (o *GetWorkerPoolsSummarySpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get worker pools summary spaces params
func (o *GetWorkerPoolsSummarySpacesParams) WithHTTPClient(client *http.Client) *GetWorkerPoolsSummarySpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get worker pools summary spaces params
func (o *GetWorkerPoolsSummarySpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get worker pools summary spaces params
func (o *GetWorkerPoolsSummarySpacesParams) WithBaseSpaceID(baseSpaceID string) *GetWorkerPoolsSummarySpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get worker pools summary spaces params
func (o *GetWorkerPoolsSummarySpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkerPoolsSummarySpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
