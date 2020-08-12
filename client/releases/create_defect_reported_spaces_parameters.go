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
)

// NewCreateDefectReportedSpacesParams creates a new CreateDefectReportedSpacesParams object
// with the default values initialized.
func NewCreateDefectReportedSpacesParams() *CreateDefectReportedSpacesParams {
	var ()
	return &CreateDefectReportedSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDefectReportedSpacesParamsWithTimeout creates a new CreateDefectReportedSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDefectReportedSpacesParamsWithTimeout(timeout time.Duration) *CreateDefectReportedSpacesParams {
	var ()
	return &CreateDefectReportedSpacesParams{

		timeout: timeout,
	}
}

// NewCreateDefectReportedSpacesParamsWithContext creates a new CreateDefectReportedSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDefectReportedSpacesParamsWithContext(ctx context.Context) *CreateDefectReportedSpacesParams {
	var ()
	return &CreateDefectReportedSpacesParams{

		Context: ctx,
	}
}

// NewCreateDefectReportedSpacesParamsWithHTTPClient creates a new CreateDefectReportedSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDefectReportedSpacesParamsWithHTTPClient(client *http.Client) *CreateDefectReportedSpacesParams {
	var ()
	return &CreateDefectReportedSpacesParams{
		HTTPClient: client,
	}
}

/*CreateDefectReportedSpacesParams contains all the parameters to send to the API endpoint
for the create defect reported spaces operation typically these are written to a http.Request
*/
type CreateDefectReportedSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create defect reported spaces params
func (o *CreateDefectReportedSpacesParams) WithTimeout(timeout time.Duration) *CreateDefectReportedSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create defect reported spaces params
func (o *CreateDefectReportedSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create defect reported spaces params
func (o *CreateDefectReportedSpacesParams) WithContext(ctx context.Context) *CreateDefectReportedSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create defect reported spaces params
func (o *CreateDefectReportedSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create defect reported spaces params
func (o *CreateDefectReportedSpacesParams) WithHTTPClient(client *http.Client) *CreateDefectReportedSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create defect reported spaces params
func (o *CreateDefectReportedSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the create defect reported spaces params
func (o *CreateDefectReportedSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateDefectReportedSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create defect reported spaces params
func (o *CreateDefectReportedSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the create defect reported spaces params
func (o *CreateDefectReportedSpacesParams) WithID(id string) *CreateDefectReportedSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create defect reported spaces params
func (o *CreateDefectReportedSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDefectReportedSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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