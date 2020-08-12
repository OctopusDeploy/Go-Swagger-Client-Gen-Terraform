// Code generated by go-swagger; DO NOT EDIT.

package team_membership

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

// NewGetTeamMembershipSpacesParams creates a new GetTeamMembershipSpacesParams object
// with the default values initialized.
func NewGetTeamMembershipSpacesParams() *GetTeamMembershipSpacesParams {
	var ()
	return &GetTeamMembershipSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamMembershipSpacesParamsWithTimeout creates a new GetTeamMembershipSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTeamMembershipSpacesParamsWithTimeout(timeout time.Duration) *GetTeamMembershipSpacesParams {
	var ()
	return &GetTeamMembershipSpacesParams{

		timeout: timeout,
	}
}

// NewGetTeamMembershipSpacesParamsWithContext creates a new GetTeamMembershipSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTeamMembershipSpacesParamsWithContext(ctx context.Context) *GetTeamMembershipSpacesParams {
	var ()
	return &GetTeamMembershipSpacesParams{

		Context: ctx,
	}
}

// NewGetTeamMembershipSpacesParamsWithHTTPClient creates a new GetTeamMembershipSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTeamMembershipSpacesParamsWithHTTPClient(client *http.Client) *GetTeamMembershipSpacesParams {
	var ()
	return &GetTeamMembershipSpacesParams{
		HTTPClient: client,
	}
}

/*GetTeamMembershipSpacesParams contains all the parameters to send to the API endpoint
for the get team membership spaces operation typically these are written to a http.Request
*/
type GetTeamMembershipSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*UserID
	  ID of the user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get team membership spaces params
func (o *GetTeamMembershipSpacesParams) WithTimeout(timeout time.Duration) *GetTeamMembershipSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get team membership spaces params
func (o *GetTeamMembershipSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get team membership spaces params
func (o *GetTeamMembershipSpacesParams) WithContext(ctx context.Context) *GetTeamMembershipSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get team membership spaces params
func (o *GetTeamMembershipSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get team membership spaces params
func (o *GetTeamMembershipSpacesParams) WithHTTPClient(client *http.Client) *GetTeamMembershipSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get team membership spaces params
func (o *GetTeamMembershipSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get team membership spaces params
func (o *GetTeamMembershipSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetTeamMembershipSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get team membership spaces params
func (o *GetTeamMembershipSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithUserID adds the userID to the get team membership spaces params
func (o *GetTeamMembershipSpacesParams) WithUserID(userID string) *GetTeamMembershipSpacesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get team membership spaces params
func (o *GetTeamMembershipSpacesParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamMembershipSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	// query param userId
	qrUserID := o.UserID
	qUserID := qrUserID
	if qUserID != "" {
		if err := r.SetQueryParam("userId", qUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
