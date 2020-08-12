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

// NewCreateTeamMembershipPreviewParams creates a new CreateTeamMembershipPreviewParams object
// with the default values initialized.
func NewCreateTeamMembershipPreviewParams() *CreateTeamMembershipPreviewParams {

	return &CreateTeamMembershipPreviewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTeamMembershipPreviewParamsWithTimeout creates a new CreateTeamMembershipPreviewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTeamMembershipPreviewParamsWithTimeout(timeout time.Duration) *CreateTeamMembershipPreviewParams {

	return &CreateTeamMembershipPreviewParams{

		timeout: timeout,
	}
}

// NewCreateTeamMembershipPreviewParamsWithContext creates a new CreateTeamMembershipPreviewParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTeamMembershipPreviewParamsWithContext(ctx context.Context) *CreateTeamMembershipPreviewParams {

	return &CreateTeamMembershipPreviewParams{

		Context: ctx,
	}
}

// NewCreateTeamMembershipPreviewParamsWithHTTPClient creates a new CreateTeamMembershipPreviewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTeamMembershipPreviewParamsWithHTTPClient(client *http.Client) *CreateTeamMembershipPreviewParams {

	return &CreateTeamMembershipPreviewParams{
		HTTPClient: client,
	}
}

/*CreateTeamMembershipPreviewParams contains all the parameters to send to the API endpoint
for the create team membership preview operation typically these are written to a http.Request
*/
type CreateTeamMembershipPreviewParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create team membership preview params
func (o *CreateTeamMembershipPreviewParams) WithTimeout(timeout time.Duration) *CreateTeamMembershipPreviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create team membership preview params
func (o *CreateTeamMembershipPreviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create team membership preview params
func (o *CreateTeamMembershipPreviewParams) WithContext(ctx context.Context) *CreateTeamMembershipPreviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create team membership preview params
func (o *CreateTeamMembershipPreviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create team membership preview params
func (o *CreateTeamMembershipPreviewParams) WithHTTPClient(client *http.Client) *CreateTeamMembershipPreviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create team membership preview params
func (o *CreateTeamMembershipPreviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTeamMembershipPreviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
