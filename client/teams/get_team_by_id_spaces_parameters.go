// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewGetTeamByIDSpacesParams creates a new GetTeamByIDSpacesParams object
// with the default values initialized.
func NewGetTeamByIDSpacesParams() *GetTeamByIDSpacesParams {
	var ()
	return &GetTeamByIDSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamByIDSpacesParamsWithTimeout creates a new GetTeamByIDSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTeamByIDSpacesParamsWithTimeout(timeout time.Duration) *GetTeamByIDSpacesParams {
	var ()
	return &GetTeamByIDSpacesParams{

		timeout: timeout,
	}
}

// NewGetTeamByIDSpacesParamsWithContext creates a new GetTeamByIDSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTeamByIDSpacesParamsWithContext(ctx context.Context) *GetTeamByIDSpacesParams {
	var ()
	return &GetTeamByIDSpacesParams{

		Context: ctx,
	}
}

// NewGetTeamByIDSpacesParamsWithHTTPClient creates a new GetTeamByIDSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTeamByIDSpacesParamsWithHTTPClient(client *http.Client) *GetTeamByIDSpacesParams {
	var ()
	return &GetTeamByIDSpacesParams{
		HTTPClient: client,
	}
}

/*GetTeamByIDSpacesParams contains all the parameters to send to the API endpoint
for the get team by Id spaces operation typically these are written to a http.Request
*/
type GetTeamByIDSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the TeamResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get team by Id spaces params
func (o *GetTeamByIDSpacesParams) WithTimeout(timeout time.Duration) *GetTeamByIDSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get team by Id spaces params
func (o *GetTeamByIDSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get team by Id spaces params
func (o *GetTeamByIDSpacesParams) WithContext(ctx context.Context) *GetTeamByIDSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get team by Id spaces params
func (o *GetTeamByIDSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get team by Id spaces params
func (o *GetTeamByIDSpacesParams) WithHTTPClient(client *http.Client) *GetTeamByIDSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get team by Id spaces params
func (o *GetTeamByIDSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get team by Id spaces params
func (o *GetTeamByIDSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetTeamByIDSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get team by Id spaces params
func (o *GetTeamByIDSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get team by Id spaces params
func (o *GetTeamByIDSpacesParams) WithID(id string) *GetTeamByIDSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get team by Id spaces params
func (o *GetTeamByIDSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamByIDSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
