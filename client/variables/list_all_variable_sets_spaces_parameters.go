// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewListAllVariableSetsSpacesParams creates a new ListAllVariableSetsSpacesParams object
// with the default values initialized.
func NewListAllVariableSetsSpacesParams() *ListAllVariableSetsSpacesParams {
	var ()
	return &ListAllVariableSetsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllVariableSetsSpacesParamsWithTimeout creates a new ListAllVariableSetsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllVariableSetsSpacesParamsWithTimeout(timeout time.Duration) *ListAllVariableSetsSpacesParams {
	var ()
	return &ListAllVariableSetsSpacesParams{

		timeout: timeout,
	}
}

// NewListAllVariableSetsSpacesParamsWithContext creates a new ListAllVariableSetsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllVariableSetsSpacesParamsWithContext(ctx context.Context) *ListAllVariableSetsSpacesParams {
	var ()
	return &ListAllVariableSetsSpacesParams{

		Context: ctx,
	}
}

// NewListAllVariableSetsSpacesParamsWithHTTPClient creates a new ListAllVariableSetsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllVariableSetsSpacesParamsWithHTTPClient(client *http.Client) *ListAllVariableSetsSpacesParams {
	var ()
	return &ListAllVariableSetsSpacesParams{
		HTTPClient: client,
	}
}

/*ListAllVariableSetsSpacesParams contains all the parameters to send to the API endpoint
for the list all variable sets spaces operation typically these are written to a http.Request
*/
type ListAllVariableSetsSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all variable sets spaces params
func (o *ListAllVariableSetsSpacesParams) WithTimeout(timeout time.Duration) *ListAllVariableSetsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all variable sets spaces params
func (o *ListAllVariableSetsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all variable sets spaces params
func (o *ListAllVariableSetsSpacesParams) WithContext(ctx context.Context) *ListAllVariableSetsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all variable sets spaces params
func (o *ListAllVariableSetsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all variable sets spaces params
func (o *ListAllVariableSetsSpacesParams) WithHTTPClient(client *http.Client) *ListAllVariableSetsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all variable sets spaces params
func (o *ListAllVariableSetsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the list all variable sets spaces params
func (o *ListAllVariableSetsSpacesParams) WithBaseSpaceID(baseSpaceID string) *ListAllVariableSetsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the list all variable sets spaces params
func (o *ListAllVariableSetsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllVariableSetsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
