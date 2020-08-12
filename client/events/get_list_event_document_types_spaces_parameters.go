// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewGetListEventDocumentTypesSpacesParams creates a new GetListEventDocumentTypesSpacesParams object
// with the default values initialized.
func NewGetListEventDocumentTypesSpacesParams() *GetListEventDocumentTypesSpacesParams {
	var ()
	return &GetListEventDocumentTypesSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetListEventDocumentTypesSpacesParamsWithTimeout creates a new GetListEventDocumentTypesSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetListEventDocumentTypesSpacesParamsWithTimeout(timeout time.Duration) *GetListEventDocumentTypesSpacesParams {
	var ()
	return &GetListEventDocumentTypesSpacesParams{

		timeout: timeout,
	}
}

// NewGetListEventDocumentTypesSpacesParamsWithContext creates a new GetListEventDocumentTypesSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetListEventDocumentTypesSpacesParamsWithContext(ctx context.Context) *GetListEventDocumentTypesSpacesParams {
	var ()
	return &GetListEventDocumentTypesSpacesParams{

		Context: ctx,
	}
}

// NewGetListEventDocumentTypesSpacesParamsWithHTTPClient creates a new GetListEventDocumentTypesSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetListEventDocumentTypesSpacesParamsWithHTTPClient(client *http.Client) *GetListEventDocumentTypesSpacesParams {
	var ()
	return &GetListEventDocumentTypesSpacesParams{
		HTTPClient: client,
	}
}

/*GetListEventDocumentTypesSpacesParams contains all the parameters to send to the API endpoint
for the get list event document types spaces operation typically these are written to a http.Request
*/
type GetListEventDocumentTypesSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get list event document types spaces params
func (o *GetListEventDocumentTypesSpacesParams) WithTimeout(timeout time.Duration) *GetListEventDocumentTypesSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get list event document types spaces params
func (o *GetListEventDocumentTypesSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get list event document types spaces params
func (o *GetListEventDocumentTypesSpacesParams) WithContext(ctx context.Context) *GetListEventDocumentTypesSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get list event document types spaces params
func (o *GetListEventDocumentTypesSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get list event document types spaces params
func (o *GetListEventDocumentTypesSpacesParams) WithHTTPClient(client *http.Client) *GetListEventDocumentTypesSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get list event document types spaces params
func (o *GetListEventDocumentTypesSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get list event document types spaces params
func (o *GetListEventDocumentTypesSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetListEventDocumentTypesSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get list event document types spaces params
func (o *GetListEventDocumentTypesSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetListEventDocumentTypesSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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