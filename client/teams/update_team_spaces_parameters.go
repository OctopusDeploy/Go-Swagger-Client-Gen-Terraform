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

	"models"
)

// NewUpdateTeamSpacesParams creates a new UpdateTeamSpacesParams object
// with the default values initialized.
func NewUpdateTeamSpacesParams() *UpdateTeamSpacesParams {
	var ()
	return &UpdateTeamSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTeamSpacesParamsWithTimeout creates a new UpdateTeamSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTeamSpacesParamsWithTimeout(timeout time.Duration) *UpdateTeamSpacesParams {
	var ()
	return &UpdateTeamSpacesParams{

		timeout: timeout,
	}
}

// NewUpdateTeamSpacesParamsWithContext creates a new UpdateTeamSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTeamSpacesParamsWithContext(ctx context.Context) *UpdateTeamSpacesParams {
	var ()
	return &UpdateTeamSpacesParams{

		Context: ctx,
	}
}

// NewUpdateTeamSpacesParamsWithHTTPClient creates a new UpdateTeamSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTeamSpacesParamsWithHTTPClient(client *http.Client) *UpdateTeamSpacesParams {
	var ()
	return &UpdateTeamSpacesParams{
		HTTPClient: client,
	}
}

/*UpdateTeamSpacesParams contains all the parameters to send to the API endpoint
for the update team spaces operation typically these are written to a http.Request
*/
type UpdateTeamSpacesParams struct {

	/*TeamResource
	  The TeamResource resource to create

	*/
	TeamResource *models.TeamResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the TeamResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update team spaces params
func (o *UpdateTeamSpacesParams) WithTimeout(timeout time.Duration) *UpdateTeamSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update team spaces params
func (o *UpdateTeamSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update team spaces params
func (o *UpdateTeamSpacesParams) WithContext(ctx context.Context) *UpdateTeamSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update team spaces params
func (o *UpdateTeamSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update team spaces params
func (o *UpdateTeamSpacesParams) WithHTTPClient(client *http.Client) *UpdateTeamSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update team spaces params
func (o *UpdateTeamSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamResource adds the teamResource to the update team spaces params
func (o *UpdateTeamSpacesParams) WithTeamResource(teamResource *models.TeamResource) *UpdateTeamSpacesParams {
	o.SetTeamResource(teamResource)
	return o
}

// SetTeamResource adds the teamResource to the update team spaces params
func (o *UpdateTeamSpacesParams) SetTeamResource(teamResource *models.TeamResource) {
	o.TeamResource = teamResource
}

// WithBaseSpaceID adds the baseSpaceID to the update team spaces params
func (o *UpdateTeamSpacesParams) WithBaseSpaceID(baseSpaceID string) *UpdateTeamSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the update team spaces params
func (o *UpdateTeamSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the update team spaces params
func (o *UpdateTeamSpacesParams) WithID(id string) *UpdateTeamSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update team spaces params
func (o *UpdateTeamSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTeamSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TeamResource != nil {
		if err := r.SetBodyParam(o.TeamResource); err != nil {
			return err
		}
	}

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
